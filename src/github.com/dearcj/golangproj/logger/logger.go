package logger

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

const (
	LAYER_DEBUG     = "**action** "
	LAYER_ACTION    = "*"
	LAYER_IMPORTANT = "**important** "
	LAYER_ERROR     = "**error** "
)

func SetupLogger() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logger.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})
}

func Debug(o ...interface{}) {
	log.Println(LAYER_DEBUG, o)
}

func Action(o ...interface{}) {
	print(LAYER_ACTION)
	fmt.Printf("%+v\n", o)
	print(LAYER_ACTION)
}

func Important(msg string) {
	log.Println(LAYER_IMPORTANT, msg)
}
