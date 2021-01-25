package analytics

import (
	zap "go.uber.org/zap"
	"gopkg.in/mgo.v2"
	"time"
)

const ACTION_CHAN_BUFF = 1

const RUN_CHAN_BUFF = 1

type ActionCustom struct {
	RunId  uint32
	UserId string
	Action *string
	Param  *string
	Time   time.Time
}

type ActionRun struct {
	CharType *string
	UserId   string
	RunId    uint32
}

type TestAnalytics struct{}

type AnalyticsInterface interface {
	InsertData()
	InitRun(chartype *string, userid string, runid uint32)
	AddAction(actionName *string, param *string, userid string, runId uint32)
	StartLoop()
	GetMgoSession() *mgo.Session
}

func (t *TestAnalytics) GetMgoSession() *mgo.Session                                              { return nil }
func (t *TestAnalytics) InsertData()                                                              {}
func (t *TestAnalytics) InitRun(chartype *string, userid string, runid uint32)                    {}
func (t *TestAnalytics) AddAction(actionName *string, param *string, userid string, runId uint32) {}
func (t *TestAnalytics) StartLoop()                                                               {}

type Analytics struct {
	loopDelay  time.Duration
	logger     *zap.Logger
	Runs       []interface{}
	colRuns    *mgo.Collection
	colActions *mgo.Collection
	Actions    []interface{}
	session    *mgo.Session
	RunsChan   chan interface{}
	ActionChan chan interface{}
}

func InitTestAnalytics() (error, AnalyticsInterface) {
	return nil, &TestAnalytics{}
}

func InitAnalytics(server string, log *zap.Logger) (error, AnalyticsInterface) {
	session, err := mgo.Dial(server)
	if err != nil {
		return err, nil
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Eventual, true)
	colRuns := session.DB("willbreaker").C("runs")
	colActions := session.DB("willbreaker").C("actions")

	return nil, &Analytics{
		logger:     log,
		session:    session,
		colRuns:    colRuns,
		colActions: colActions,
		ActionChan: make(chan interface{}, ACTION_CHAN_BUFF),
		RunsChan:   make(chan interface{}, RUN_CHAN_BUFF),
	}
}

func (a *Analytics) InsertData() {
	if len(a.Runs) > 0 {
		err := a.colRuns.Insert(a.Runs...)
		if err != nil {
			a.logger.Error("Can't insert run info", zap.Error(err))
		}
		a.Runs = make([]interface{}, 0)
	}

	if len(a.Actions) > 0 {
		err := a.colActions.Insert(a.Actions...)
		if err != nil {
			a.logger.Error("Can't action info", zap.Error(err))
		}
		a.Actions = make([]interface{}, 0)
	}
}

func (a *Analytics) InitRun(chartype *string, userid string, runid uint32) {
	a.RunsChan <- &ActionRun{
		CharType: chartype,
		RunId:    runid,
		UserId:   userid,
	}
}

func (a *Analytics) StartLoop() {
	tickerAnalytics := time.NewTicker(time.Minute)

	go func(tickerAnalytics *time.Ticker) {
		for {
			select {
			case <-tickerAnalytics.C:
				a.InsertData()
			case x := <-a.ActionChan:
				a.Actions = append(a.Actions, x)
			case x := <-a.RunsChan:
				a.Runs = append(a.Runs, x)
			}
		}
	}(tickerAnalytics)
}

func (a *Analytics) GetMgoSession() *mgo.Session {
	return a.session
}

func (a *Analytics) AddAction(actionName *string, param *string, userid string, runId uint32) {
	a.ActionChan <- &ActionCustom{
		RunId:  runId,
		UserId: userid,
		Action: actionName,
		Param:  param,
		Time:   time.Now(),
	}
}
