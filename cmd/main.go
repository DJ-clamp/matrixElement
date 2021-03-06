package main

import (
	"fmt"

	. "github.com/DJ-clamp/matrixElement/common"
	"github.com/DJ-clamp/matrixElement/routers"
)

var (
	commit   = ""
	compiled = ""
)

func main() {
	println("----matrixElement----")
	println("build:", compiled, "commit:", commit)
	NewServer(fmt.Sprintf(":%s", HTTP_PORT),
		BootingErrorLog("Severice is failed.: %v"),
		BootingLog("HTTP is works"),
		DebugMode(APP_DEBUG),
		TimeOut(10),
		AddRouters(routers.StartPage),
	)
	select {}
}
