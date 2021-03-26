package common

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

//监控进程
type Monitor struct {
	debug          bool
	welcomeText    string
	tlsCrt         string
	tlsKey         string
	staticFilePath string
	bootFailed     string
	timeOut        time.Duration
	routers        []func(*gin.Engine)
}
type ServerConfig func(*Monitor)

func BootingLog(text string) ServerConfig {
	return func(monitor *Monitor) {
		monitor.welcomeText = text
	}
}

func StaticFilePath(path string) ServerConfig {
	return func(monitor *Monitor) {
		monitor.staticFilePath = path
	}
}
func AddRouters(route ...func(*gin.Engine)) ServerConfig {
	return func(monitor *Monitor) {
		monitor.routers = route
	}
}

func BootingErrorLog(err string) ServerConfig {
	return func(monitor *Monitor) {
		monitor.bootFailed = err
	}
}

func TLSMode(tlsCrt string, tlsKey string) ServerConfig {
	return func(monitor *Monitor) {
		monitor.tlsCrt = tlsCrt
		monitor.tlsKey = tlsKey
	}
}

func TimeOut(t int) ServerConfig {
	return func(monitor *Monitor) {
		monitor.timeOut = time.Duration(t) * time.Second
	}
}

func DebugMode(isDebug bool) ServerConfig {
	return func(monitor *Monitor) {
		monitor.debug = isDebug
	}
}

var defaultServerOption = Monitor{
	debug:      false,
	bootFailed: "Aimenet HTTP server failed...: %v",
	timeOut:    10,
}

func NewServer(add string, opts ...ServerConfig) {
	options := defaultServerOption
	for _, opt := range opts {
		opt(&options)
	}

	go func() {
		if options.welcomeText != "" {
			log.Println(options.welcomeText)
		}
		if !options.debug {
			gin.SetMode(gin.ReleaseMode)
		}
		r := gin.New()
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
		// r.Use(middleware.TimeoutMiddleware(options.timeOut))
		//s := &http.Server{
		//	Addr:           ":8080",
		//	Handler:        r,
		//	ReadTimeout:    10 * time.Second,
		//	WriteTimeout:   10 * time.Second,
		//	MaxHeaderBytes: 1 << 20,
		//}
		if len(options.routers) > 0 {
			for _, route := range options.routers {
				route(r)
			}
		}
		if options.tlsKey != "" {
			if e := r.RunTLS(add, options.tlsCrt, options.tlsKey); e != nil {
				utils.Logger.Printf(options.bootFailed, e)
			}
		} else {
			if e := r.Run(add); e != nil {
				utils.Logger.Printf(options.bootFailed, e)
			}
		}
	}()
}
