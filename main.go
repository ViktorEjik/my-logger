package main

import (
	"fmt"
	"strconv"

	"github.com/ViktorEjik/my-loggermy/pkg/loger"
)

func main() {
	var levels = [...]loger.LogLevel{loger.Debug, loger.Info, loger.Warning, loger.Error}
	for i, value := range levels {
		err := loger.CreateLoger("tests/PLANE/test"+strconv.Itoa(i)+".txt", value, loger.PLANE)
		if err != nil {
			fmt.Print(err)
		}
		loger.WriteLog("Test1", loger.Info)
		loger.WriteLog("Test3", loger.Debug)
		loger.WriteLog("Test4", loger.Warning)
		loger.WriteLog("Test2", loger.Error)
	}
	for i, value := range levels {
		err := loger.CreateLoger("tests/CSV/test"+strconv.Itoa(i)+".txt", value, loger.CSV)
		if err != nil {
			fmt.Print(err)
		}
		loger.WriteLog("Test1", loger.Info)
		loger.WriteLog("Test3", loger.Debug)
		loger.WriteLog("Test4", loger.Warning)
		loger.WriteLog("Test2", loger.Error)
	}
}
