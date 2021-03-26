package main

import (
	"fmt"
)

var (
	comit            = ""
	compiled         = ""
	HTTP_PORT string = ""
)

func main() {
	initDB()
	hp := initMainHTTP()
	hp.Run(fmt.Sprintf(":%s", HTTP_PORT))
}
