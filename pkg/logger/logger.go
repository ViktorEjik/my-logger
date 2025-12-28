package loger

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
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
)

var logerCfg struct {
	logFile   *os.File
	logLevel  LogLevel
	logFormat LogFormat
}

func CreateLogger(path string, minLogLevel LogLevel, logFormat LogFormat) error {
	if path == "console" {
		logerCfg.logFile = os.Stdout
	} else {
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

		if err != nil {
			return errors.New("Can not create or open file! " + err.Error())
		}

		logerCfg.logFile = file
	}

	if logFormat == CSV {
		if path != "console" {
			if err := os.Truncate(path, 0); err != nil {
				return errors.New("Failed to truncate: " + err.Error())
			}
		}
		_, err := logerCfg.logFile.WriteString("Date,Level,File,Line,Func,MESSAGE\n")
		if err != nil {
			return errors.New("Can not write message: " + err.Error())
		}
	}

	if minLogLevel > Error {
		return errors.New("Min log level larget then Error")
	}

	logerCfg.logLevel = minLogLevel

	if PLANE > logFormat || CSV < logFormat {
		return errors.New("Undefined log format")
	}

	logerCfg.logFormat = logFormat
	return nil
}

func getCallerInfo(skip int) (function, file string, line int) {
	pc, filePath, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown", "unknown", 0
	}

	function = runtime.FuncForPC(pc).Name()
	file = filepath.Base(filePath)

	if idx := strings.LastIndex(function, "/"); idx != -1 {
		function = function[idx+1:]
	}

	return function, file, line
}

func getLevelString(level LogLevel) string {
	switch level {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

func writePLANE(message string, level LogLevel) error {
	callFunction, callFile, callLine := getCallerInfo(3)
	callStr := fmt.Sprintf("%s:%d %s", callFile, callLine, callFunction)

	line := fmt.Sprintf("Date: %v, Func: %s, Message: %s, Level: %s\n", time.Now().Format(time.Stamp), callStr, message, getLevelString(level))

	_, err := logerCfg.logFile.WriteString(line)

	if err != nil {
		return errors.New("Can not write message: " + err.Error())
	}
	return nil
}

func writeCSV(message string, level LogLevel) error {
	callFunction, callFile, callLine := getCallerInfo(3)

	line := fmt.Sprintf("%v,%s,%s,%d,%s,%s\n",
		time.Now().Format(time.Stamp),
		getLevelString(level),
		callFile,
		callLine,
		callFunction,
		message,
	)
	_, err := logerCfg.logFile.WriteString(line)
	if err != nil {
		return errors.New("Can not write message: " + err.Error())
	}
	return nil
}

func WriteLog(message string, level LogLevel) error {
	if level < logerCfg.logLevel {
		return nil
	}
	switch logerCfg.logFormat {
	case PLANE:
		err := writePLANE(message, level)
		if err != nil {
			return err
		}
	case CSV:
		err := writeCSV(message, level)
		if err != nil {
			return err
		}
	default:
		return errors.New("Unsupport")
	}
	return nil
}
