package main

import (
	"fmt"

	. "github.com/DJ-clamp/matrixElement/common"
	"github.com/DJ-clamp/matrixElement/routers"
)

var (
	comit    = ""
	compiled = ""
)

func main() {
	initDB()
	NewServer(fmt.Sprintf(":%s", HTTP_PORT),
		BootingErrorLog("Severice is failed.: %v"),
		BootingLog("HTTP is works"),
		DebugMode(true),
		TimeOut(10),
		AddRouters(routers.StartPage),
	)
}
