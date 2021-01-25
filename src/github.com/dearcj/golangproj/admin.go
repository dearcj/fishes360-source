package main

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/vincent-petithory/dataurl"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var JWT_KEY = "lNMW2ChNRb"

func jwtCookieMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtstr := r.Header.Get("jwt")
		token, err := checkToken(jwtstr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			if token != nil && token.Claims != nil {
				st := token.Claims.(*Claims)
				if st.Email == nodeConfig.Admin.Login {
					next.ServeHTTP(w, r)
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}
		}

	})
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func createToken(log string, pass string) (err error, tokenString string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		log, jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * 10 * time.Hour).Unix(),
			Issuer:    "test",
		}})

	tokenString, err = token.SignedString([]byte(JWT_KEY))
	return
}

func checkToken(strToken string) (token *jwt.Token, err error) {
	token, err = jwt.ParseWithClaims(strToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {

		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return []byte(JWT_KEY), nil
	})

	if err != nil {
		return nil, errors.New("Wrong login")
	}

	if token != nil && strings.ToLower(token.Claims.(*Claims).Email) == strings.ToLower(nodeConfig.Admin.Login) {
		return
	} else {
		return nil, errors.New("Wrong login")
	}
}

func getRunSettings(w http.ResponseWriter, r *http.Request) {
	if len(server.runs) > 0 {
		server.RunMutex.RLock()
		run := server.runs[len(server.runs) - 1]

		var arr []byte
		arr, err := json.Marshal(run.timeline.Fishes)
		server.RunMutex.RUnlock()
		if err == nil {
			w.WriteHeader(http.StatusOK)
			w.Write(arr)
			return
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getCurrentSettings(w http.ResponseWriter, r *http.Request) {
	timelineFile, err := ioutil.ReadFile(nodeConfig.Server.TimelineFile)
	if err == nil {
		timeline := &FishTimeline{}
		err = TimelineFromFile(timelineFile, timeline)
		if err == nil {
			var arr []byte
			arr, err = json.Marshal(timeline)
			if err == nil {
				w.WriteHeader(http.StatusOK)
				w.Write(arr)
				return
			}
		}
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func changeFishImage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			imgid := r.Header.Get("imgid")
			i, err := strconv.ParseInt(imgid, 10, 8)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {

				data, err := dataurl.DecodeString(string(body))

				println("get fish image", string(data.Data)[:200])
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					f, err := os.OpenFile("./front/data/fishes/fish"+strconv.Itoa(int(i))+".jpg", os.O_CREATE|os.O_WRONLY, 0777)
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
					} else {
						_, err = f.Write(data.Data)
						defer f.Close()

						if err != nil {
							w.WriteHeader(http.StatusInternalServerError)
						} else {
							println("Write file succesfull")
							w.WriteHeader(http.StatusOK)
						}

					}

				}
			}
		}

	}

}

func updateSettings(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Header().Add("Content-Type", "application/json")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {

			ft := &FishTimeline{}
			err := json.Unmarshal(body, ft)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				err := ioutil.WriteFile(nodeConfig.Server.TimelineFile, body, 0644)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					w.WriteHeader(http.StatusOK)
				}
			}
		}

	}
}

func adminLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		email := r.URL.Query().Get("email")
		pass := r.URL.Query().Get("pass")

		if strings.ToLower(email) == strings.ToLower(nodeConfig.Admin.Login) &&
			strings.ToLower(pass) == strings.ToLower(nodeConfig.Admin.Pass) {
			//success
			err, token := createToken(email, pass)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {

				http.SetCookie(w, &http.Cookie{
					Name:    "jwt",
					Value:   token,
					Expires: time.Now().Add(time.Hour * 24 * 10),
				})
				//w.WriteHeader(http.StatusOK)
			}

		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
