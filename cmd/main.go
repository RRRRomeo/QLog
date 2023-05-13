package main

import (
	"github.com/QLog/api"
)

func main() {
	logger := api.LoggerInit(0, 1, "./test.log")
	logger.Debugf("test!\n")

}
