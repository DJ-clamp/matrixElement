package utils

import (
	"io"
	"log"
	"os"
	"strconv"
)

var Logger *log.Logger
var debug bool

const CustomLogOutputPath = "Log.txt"

func init() {
	f, err := os.OpenFile(CustomLogOutputPath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	//defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	Logger = log.New(wrt, "", log.LstdFlags)
}

func DebugInfo(ctx ...interface{}) {
	debug, _ = strconv.ParseBool(GetEnv("DEBUG", "false"))
	if debug {
		ctx = append([]interface{}{"[D]"}, ctx...)
		Logger.Println(ctx...)
	}
}

func Debugf(f string, ctx ...interface{}) {
	debug, _ = strconv.ParseBool(GetEnv("DEBUG", "false"))
	if debug {
		Logger.Printf(f, ctx...)
	}
}

func Info(ctx ...interface{}) {
	debug, _ = strconv.ParseBool(GetEnv("DEBUG", "false"))
	if debug {
		ctx = append([]interface{}{"[I]"}, ctx...)
		Logger.Println(ctx...)
	}
}

func Warning(ctx ...interface{}) {
	debug, _ = strconv.ParseBool(GetEnv("DEBUG", "false"))
	if debug {
		ctx = append([]interface{}{"[W]"}, ctx...)
		Logger.Println(ctx...)
	}
}

func Error(ctx ...interface{}) {
	debug, _ = strconv.ParseBool(GetEnv("DEBUG", "false"))
	if debug {
		ctx = append([]interface{}{"[E]"}, ctx...)
		Logger.Println(ctx...)
	}
}
