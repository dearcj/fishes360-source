package main

import (
	"compress/gzip"
	"errors"
	"flag"
	"fmt"
	ws "github.com/gorilla/websocket"
	"github.com/jinzhu/configor"
	errors2 "github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"log"
	//"math"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	//"path"
	p "path/filepath"
	"runtime"
	"runtime/pprof"
)

const PROFILER = false
const OVERLAP_SQUARED_DIST =  14.* 14.
const OVERLAP_SQUARED_DIST_TEST =  14.* 14.

var (
	_, currentFilePath, _, _ = runtime.Caller(0)
	dirPath                  = p.Dir(currentFilePath)
	database                 Database
	upgrader                 = ws.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	server     *Server
	Stats      *ServerStats
	nodeConfig *YAMLConfig
	//timelineFile []byte
)
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func httpKeyMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		httpKey := q.Get("http_key")
		if httpKey == "" {
			httpKey = r.Header.Get("http_key")
		}

		if httpKey != nodeConfig.Server.Http_Key {
			server.logger.Error("Wrong http key", zap.Error(errors.New("Client sent wrong http key")))

			return
		}

		next.ServeHTTP(w, r)
	})
}

func Gzip(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			handler.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzw := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		handler.ServeHTTP(gzw, r)
	})
}

func ServeWS(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		//todo: check origin
		return true
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		server.logger.Error("Websocket upgrader fails", zap.Error(err))

		return
	}

	q := r.URL.Query()
	lg := q.Get("log")
	pass := q.Get("pass")
	token := q.Get("token")
	puuid := q.Get("puuid")

	info, err := server.GetUserInfo(token, puuid)

	if err != nil {
		info = &UserInfo{
			Balance: 121212,
			Id: "3232323",
		}
		server.HandleConnection(conn, lg, pass, r.RemoteAddr, CreateAccount(info, token, puuid))
		return
	}

	if err != nil {
		server.logger.Error("No tokens / WS disconnected")
		conn.Close()
		return
	} else {
		server.HandleConnection(conn, lg, pass, r.RemoteAddr, CreateAccount(info, token, puuid))
	}
}

func NoCacheWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
		next.ServeHTTP(w, r)
	})
}

func CORSWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
func CheckFishOverlap() {
	//THRESHOLD := 2
	Overlap := func(a, b *Object) bool {
		afish := a.FindByComponent(confComponents.Fish).(*Fish)
		bfish := b.FindByComponent(confComponents.Fish).(*Fish)
	/*	if afish.Curve.Inx == bfish.Curve.Inx {
			if math.Abs(float64(afish.PosOnCurve())-float64(bfish.PosOnCurve())) < float64(THRESHOLD) {
				return true
			}
		}*/
		if afish.GetCurveProp() > 1 || bfish.GetCurveProp() > 1 {
			return false
		}
		//afishPos := afish.Curve.GetCachedPoint(afish.GetCurveProp())
		//bfishPos := bfish.Curve.GetCachedPoint(bfish.GetCurveProp())

		if afish.CachedPos.DistSq(bfish.CachedPos) < OVERLAP_SQUARED_DIST_TEST {
			server.logger.Info("",
				zap.Uint32("Curve", afish.Curve.Inx),
				zap.Float64("%", afish.GetCurveProp()),
				zap.Uint32("Curve", bfish.Curve.Inx),
				zap.Float64("%", bfish.GetCurveProp()))

			return true
		}

		return false
	}

	CheckOverlap := func(allfish []*Object) int {
		count := 0

		for inx, x := range allfish {
			for jinx := inx + 1; jinx < len(allfish); jinx++ {
				if Overlap(x, allfish[jinx]) {
					count++
				}
			}

		}
		return count
	}

	run := server.CreateRun(1)
	run.timeline.CurrentScene = run.timeline.Scenes[0]
	overlaps := 0
	for x := 0; x < 5000; x++ {
		step := time.Millisecond * 200
		run.timeline.TimeDelta += step
		run.timeline.CurrentScene.FishUpdate(step, run, true, time.Millisecond*200, true, false)
		run.factory.Update(step)
		overlaps += CheckOverlap(run.factory.FilterObjects(confComponents.Fish))
	}


	fmt.Print("OVERLAPS ", overlaps)
}

func main() {
	logger := InitConfigLogger()

	defer logger.Sync()

	var err error
	err, server = CreateServer(logger)
	if err != nil {
		log.Fatal("Can't start server", err)
		panic(err)
	}

	if PROFILER {
		f, _ := os.Create("CPUProfiler")
		logger.Info("Launched CPU profiler")
		err := pprof.StartCPUProfile(f)
		if err != nil {
			logger.Error("Can't start profiler", zap.Error(err))
		}
	}

	err = saveConfig()
	if err != nil {
		panic(err)
	}

	//testing timeline load
	tl := &FishTimeline{}
	err = TimelineFromFile(GetTimelineFile(), tl)
	if err != nil {
		panic(err)
	}


	server.Loop()
	fs := http.FileServer(http.Dir(dirPath + "/front/game/dist"))
	http.HandleFunc("/login", adminLogin)
	http.Handle("/", CORSWrapper(NoCacheWrapper(Gzip(fs))))
	http.Handle("/ws/", httpKeyMiddleware(ServeWS))
	http.Handle("/changefishimage", jwtCookieMiddleware(changeFishImage))
	http.Handle("/updatesettings", jwtCookieMiddleware(updateSettings))
	http.Handle("/getsettings", jwtCookieMiddleware(getCurrentSettings))
	http.HandleFunc("/get-run-settings", getRunSettings)

	http.Handle("/metrics_custom", server.stats.Exporter)
	http.Handle("/metrics_prom", promhttp.Handler())

	if nodeConfig.Server.SSL {
		logger.Info("Listen HTTPS", zap.String("port", nodeConfig.Server.SSLPort))

		err := http.ListenAndServeTLS(nodeConfig.Server.SSLPort, "wb.crt", "wb.key", nil)
		if err != nil {
			server.logger.Error("Can't start HTTPS server", zap.Error(err))
		}
	} else {
		logger.Info("Listen HTTP", zap.String("port", nodeConfig.Server.Port))

		err := http.ListenAndServe(nodeConfig.Server.Port, nil)
		if err != nil {
			server.logger.Error("Can't start HTTP server", zap.Error(err))
		}
	}
}

func GetTimelineFile() (file []byte) {
	file, err := ioutil.ReadFile(nodeConfig.Server.TimelineFile)
	if err != nil {
		panic(errors2.WithMessage(err, "Can't load timeline file"))
	}
	return
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func InitConfigLogger() (logger *zap.Logger) {
	argsWithProg := os.Args
	fmt.Println("CommandLine args:", argsWithProg)

	var configFilename string
	flag.StringVar(&configFilename, "config", `./configs/default.yaml`, "File path to config.json file")
	flag.Parse()

	nodeConfig = DefaultYamlConfig()
	if configFilename != "" {
		dir, _ := os.Getwd()
		println("DIR: ", dir)
		confPath := path.Join(dir, configFilename)

		err := configor.Load(nodeConfig, confPath, configFilename)
		println("Set http_key:", nodeConfig.Server.Http_Key)

		if err != nil {
			panic(err)
		}
	} else {
		println("config file not found")
	}

	println("LOGFILE:" + nodeConfig.Debug.Log_File)

	logger = initLogger()

	if configFilename != "" {
		logger.Info("Loaded config", zap.String("filename", configFilename))
	}
	return
}
