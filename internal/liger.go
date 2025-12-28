package loger

import (
	"os"
)

type LogLevel uint8

type LogFormat uint8

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)

const (
	PLANE LogFormat = iota
	CSV
	JSON
)

var LogerCfg struct {
	logFile   *os.File
	logLevel  LogLevel
	logFormat LogFormat
}

func CreateLoger(path string, minLogLevel LogLevel, logFormat LogFormat) interface{} {
	if path == "console" {
		LogerCfg.logFile = os.Stdout
	} else {
		file, err := os.OpenFile(path, os.O_APPEND, 0666)

		if err != nil {
			return "Can not create or open file!"
		}

		LogerCfg.logFile = file
	}

	if minLogLevel > Error {
		return "Min log level larget then Error"
	}

	LogerCfg.logLevel = min(Debug, minLogLevel)

	if PLANE > logFormat || JSON < logFormat {
		return "Undefined log format"
	}

	LogerCfg.logFormat = logFormat
	return nil
}
