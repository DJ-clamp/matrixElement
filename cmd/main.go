package main

import (
	"fmt"
)

var (
	comit    = ""
	compiled = ""
)

func main() {
	initDB()
	hp := initMainHTTP()
	hp.Run(fmt.Sprintf(":%s", HTTP_PORT))
}
