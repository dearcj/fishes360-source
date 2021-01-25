package main

import (
	"encoding/json"
	"github.com/dearcj/golangproj/bitmask"
	"github.com/fatih/structs"
	"go.uber.org/zap"
	"io/ioutil"
	"time"
)

type MapID = uint32

type Components struct {
	Unit       bitmask.Bitmask
	Fish       bitmask.Bitmask
	Player     bitmask.Bitmask
	MiscObject bitmask.Bitmask
}

type ConnectionRemoveReason uint32

const (
	CON_REM_DISCONNECT ConnectionRemoveReason = iota
	CON_REM_BYE
)

type SessionStateType uint32

type TurnStateType uint32

type ConnectionMessage struct {
	RECONNECT_AVAILABLE     uint32
	ERR_NO_SUCH_CREDENTIALS uint32
	CON_SUCCESS_LOGIN       uint32
}

type Commands struct {
	CMD_FINAL_EXIT_GAME MapID
	CMD_READY           MapID
	CMD_RECONNECT_YES   MapID
	CMD_RECONNECT_NO    MapID
	CMD_LEAVE_QUEUE     MapID
	CMD_SHOOT           MapID
	CMD_ANGLE_CHANGE    MapID
	CMD_CHANGE_GUN      MapID
	CMD_NEXT_SCENE_HACK MapID
	CMD_REQUEST_ROOM    MapID
}

type PhysConfig struct {
	ServerIterations int
	ClientIterations int
	MoneySize        float64
	PlayerSlowdown   float64
	PlayerRadius     float64
}

type ClientConfig struct {
	StateElapsedPrecision int
	ScreenWidth           int
	ScreenHeight          int
	MaxScreenWidth        int
	MaxScreenHeight       int
	Host                  string
	PhysIterations        int
}

type SessionState struct {
	InQueue       SessionStateType
	InQueueReady  SessionStateType
	InGame        SessionStateType
	InShop        SessionStateType
	SelectingChar SessionStateType
}

type PNotID int32
type PlayerNotifications map[PNotID]string

const (
	___ PNotID = iota
	pnNeedPicklock
	pnOpened
)

func DefaultYamlConfig() *YAMLConfig {
	return &YAMLConfig{
		Debug: DebugConfig{},
		Admin: AdminConfig{},
		Server: ServerConfig{
			GameUID:               "e9ab1471d4dedadf2269a877936e8248",
			GameName:              "Dead Fish 360",
			TimelineFile:          "front/data/timeline.json",
			Http_Key:              "", //TODO: REMOVE THIS
			MongoServer:           "localhost",
			DBConLine:             "postgres://postgres:admin@localhost:5432/?dbname=Willbreaker&port=5432&sslmode=disable",
			MinimumToStart:        1,
			SSL:                   false,
			SSLPort:               ":443",
			Port:                  ":80",
			SocialResolutionTime:  6 * time.Second,
			DBUpdateDelay:         20 * time.Second,
			ResourcePath:          "/front/",
			StatsReportDelay:      1 * time.Second,
			MaxRoomsPlayers:       20,
			WriteWait:             10 * time.Second,
			PongWait:              60 * time.Second,
			PingPeriod:            (60 * time.Second * 7) / 10,
			MaxMessageSize:        1024,
			ServerPhysDelta:       1 / 60.,
			FrameDelay:            30 * time.Millisecond,
			RoomLogDelay:          2 * time.Second,
			QueueTime:             20000 * time.Second,
			ConnectionsBufferSize: 30,
			StatsUpdateTime:       1 * time.Second,
			AnalyticsUpdateTime:   5 * time.Second,
		},
	}
}

type ServerConfig struct {
	GameUID               string
	GameName              string
	TimelineFile          string
	Http_Key              string
	MongoServer           string
	DBConLine             string
	SSL                   bool `default: "false"`
	MinimumToStart        int
	SSLPort               string
	Port                  string
	SocialResolutionTime  time.Duration
	ResourcePath          string
	MaxRoomsPlayers       int
	FrameDelay            time.Duration
	RoomLogDelay          time.Duration
	ServerPhysDelta       float64
	WriteWait             time.Duration
	PongWait              time.Duration
	PingPeriod            time.Duration
	MaxMessageSize        int64
	DBUpdateDelay         time.Duration
	QueueTime             time.Duration
	StatsReportDelay      time.Duration
	StatsUpdateTime       time.Duration
	AnalyticsUpdateTime   time.Duration
	ConnectionsBufferSize int
}

type AdminConfig struct {
	Login string `default: "admin@gmail.com"`
	Pass  string `default: "admin123"`
}

type DebugConfig struct {
	Log_File  string `default: ""`
	Log_Level string `default: ""`
}

type YAMLConfig struct {
	Admin  AdminConfig  `json:"-"`
	Debug  DebugConfig  `json:"-"`
	Server ServerConfig `json:"-"`
}

type AnimationsConfig struct {
	Fireball uint32
}

type Config struct {
	MaxPlayers        int          `json:"-"`
	SessionState      SessionState `json:"-"`
	ConnectionMessage ConnectionMessage
	Second            time.Duration
	Millisecond       time.Duration
	Components        Components
	Commands          Commands
	Client            ClientConfig `json:"-"`
	Player            PlayerConfig `json:"-"`
	Actions           *Actions
	ConfBuffMap       map[string]interface{} `json:"-"`
}

type PlayerConfig struct {
	MoveIterations        uint32
	MoveDuration          time.Duration
	DisconnectWait        time.Duration
	InventorySize         int
	PrepareTime           time.Duration
	TurnTime              time.Duration
	ExecTime              time.Duration
	TurnNotificationDelay time.Duration
	EnergyPerFrame        float64
	MaxEnergy             float64
	DefaultMaxHealth      float64
	FireDelay             time.Duration
	AVGPing               time.Duration
	Speed                 float64
	CellCrossTime         time.Duration
	DoorTransitionTime    time.Duration
	XPPerLevel            []uint32
}

type TurnStateTypes struct {
	BATTLE_PREPARE  TurnStateType
	BATTLE_OUR_TURN TurnStateType
	PAUSE           TurnStateType
	BATTLE_EXEC     TurnStateType
}

type AnalyticActions struct {
	ActionGameplay     string
	ActionView         string
	ActionBuyItem      string
	ActionPickChar     string
	ActionDeath        string
	ActionUseSkill     string
	ActionLogin        string
	ActionWaitOn       string
	FinishedBattleTurn string
	FinishedLootTurn   string
	FinishedSocialTurn string
	ActionSelectPerk   string
	ActionGetLevel     string
	ActionWonDesire    string
}

var confAnalytics = AnalyticActions{
	ActionLogin:        "login",
	ActionGameplay:     "gameplay",
	ActionView:         "views",
	ActionBuyItem:      "buyitem",
	ActionPickChar:     "pickchar",
	ActionDeath:        "death",
	ActionUseSkill:     "useskill",
	ActionWaitOn:       "waiton",
	FinishedBattleTurn: "fin_battle_turn",
	FinishedLootTurn:   "fin_loot_turn",
	FinishedSocialTurn: "fin_social_turn",
	ActionWonDesire:    "desire_won",
}

type ProgressConfig struct {
}

type ActionTypes struct {
	IAT_TAKE uint32
	IAT_GIVE uint32
}

var confComponents = (UNSAFE_INCREMENT_UINT32_POW_2_STRUCT(&Components{})).(*Components)

var confCommands = (UNSAFE_INCREMENT_UINT32_STRUCT(&Commands{})).(*Commands)
var confAnims = (UNSAFE_INCREMENT_UINT32_STRUCT(&AnimationsConfig{})).(*AnimationsConfig)

var config = Config{
	MaxPlayers: 8,
	ConnectionMessage: ConnectionMessage{
		//OK MSG
		RECONNECT_AVAILABLE: 1,
		CON_SUCCESS_LOGIN:   2,

		//ERRORS
		ERR_NO_SUCH_CREDENTIALS: 1024,
	},
	Player: PlayerConfig{
		DisconnectWait: 0 * time.Second,
	},
	Second:      time.Second,
	Millisecond: time.Millisecond,

	Actions: confActions,

	//PLAYER COMMANDS
	Commands: *confCommands,

	//COLLISION GROUPS
	//For now 1 ColGroupId == ClassUniqueID
	Components: *confComponents,
}

var confCommandsMap = InterfaceMapToMapID(structs.New(confCommands).Map())

func FillConfigIDS() {

	/*for key := range config.Monsters.MonstersList {
		f := config.Monsters.MonstersList[key]
		f.ID = key
	}*/
}

func saveConfig() error {
	FillConfigIDS()
	ba, err := json.Marshal(config)
	if err != nil {
		server.logger.Error("Config marshal error", zap.Error(err))
	}
	err = ioutil.WriteFile(dirPath+"/front/config.json", ba, 0644)
	if err != nil {
		server.logger.Error("Config write config", zap.Error(err))
	}

	//	body := `define("config", ['require', 'exports'], function (require, exports) {  exports.config = ` + string(ba) + `;
	//return {config: exports.config}});`
	body := `export default 
` + string(ba) + `
;`

	err = ioutil.WriteFile(dirPath+"/front/config.js", []byte(body), 0644)

	if err != nil {
		return err
	}

	//body = "module.exports.config = " + string(ba) + ""
	//err = ioutil.WriteFile(dirPath+"/front/config_require.js", []byte(body), 0644)

	return err
}
