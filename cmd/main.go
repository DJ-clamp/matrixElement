package main

import (
	"fmt"

	. "github.com/DJ-clamp/matrixElement/common"
	"github.com/DJ-clamp/matrixElement/routers"
	"github.com/DJ-clamp/matrixElement/utils"
)

var (
	commit   = ""
	compiled = ""
)

func main() {
	utils.Logger.Println("----matrixElement----")
	utils.Logger.Println("build:", compiled, "commit:", commit)
	initDB()
	NewServer(fmt.Sprintf(":%s", HTTP_PORT),
		BootingErrorLog("Severice is failed.: %v"),
		BootingLog("HTTP is works"),
		DebugMode(true),
		TimeOut(10),
		AddRouters(routers.StartPage),
	)
	select {}
}
