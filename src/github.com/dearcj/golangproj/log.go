package main

import (
	"encoding/json"
	"go.uber.org/zap"
	"log"
)

var templateStr = `
{
  "level": "debug",
  "encoding": "json",
  "outputPaths": ["stdout"],
  "errorOutputPaths": ["stderr"],
  "encoderConfig": {
    "messageKey": "message",
    "levelKey": "l",
    "levelEncoder": "lowercase"
  }
}
`

func initLogger() *zap.Logger {
	var cfg zap.Config

	if err := json.Unmarshal([]byte(templateStr), &cfg); err != nil {
		panic(err)
	}

	if nodeConfig.Debug.Log_File != "" {
		println("Dump logfiles to: ", nodeConfig.Debug.Log_File)
		cfg.OutputPaths = append(cfg.OutputPaths, nodeConfig.Debug.Log_File)
	}

	if nodeConfig.Debug.Log_Level != "" {
		err := cfg.Level.UnmarshalText([]byte(nodeConfig.Debug.Log_Level))
		if err != nil {
			log.Fatalf("No such log leve", err)
		}

	}
	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("Can't build zap logger config: %v", err)
	}

	return logger
}
