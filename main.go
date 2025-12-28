package main

import (
	"fmt"
	"strconv"

	logger "github.com/ViktorEjik/my-loggermy/pkg/logger"
)

func main() {
	var levels = [...]logger.LogLevel{logger.Debug, logger.Info, logger.Warning, logger.Error}
	for i, value := range levels {
		err := logger.CreateLogger("tests/PLANE/test"+strconv.Itoa(i)+".txt", value, logger.PLANE)
		if err != nil {
			fmt.Print(err)
		}
		logger.WriteLog("Test1", logger.Info)
		logger.WriteLog("Test3", logger.Debug)
		logger.WriteLog("Test4", logger.Warning)
		logger.WriteLog("Test2", logger.Error)
	}
	for i, value := range levels {
		err := logger.CreateLogger("tests/CSV/test"+strconv.Itoa(i)+".txt", value, logger.CSV)
		if err != nil {
			fmt.Print(err)
		}
		logger.WriteLog("Test1", logger.Info)
		logger.WriteLog("Test3", logger.Debug)
		logger.WriteLog("Test4", logger.Warning)
		logger.WriteLog("Test2", logger.Error)
	}
}
