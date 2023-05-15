package main

import (
	"github.com/RRRRomeo/QLog/api"
)

func main() {
	logger := api.LoggerInit(0, 1, "./test.log")
	logger.Debugf("test!\n")
	api.Debugf("test in logger\n")

}
