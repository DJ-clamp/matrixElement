package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DJ-clamp/matrixElement/models"
	"github.com/DJ-clamp/matrixElement/utils"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	HTTP_PORT    string = ""
	DB_NAME      string = ""
	LoggerPrefix string = "[ME]"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}
	HTTP_PORT = os.Getenv("HTTP_PORT")
	//数据库出手啊
	initDB()
	isAutoMigrate := flag.Bool("migrate", false, "auto migrate database")
	flag.Parse()
	if *isAutoMigrate {
		defer os.Exit(0)
		RunDBCommand()
	}
	//初始化日志
	loadLogger()
	//程序退出操作
	endInit()

}

func initDB() {
	timeout := make(chan *gorm.DB)
	go func() {
		select {
		case <-time.After(time.Second * 15):
			log.Println("DB connection timeout")
			os.Exit(syscall.AF_UNSPEC)
		case <-timeout:
			return
		}
	}()
	db := models.InitDB()
	timeout <- db
}

func endInit() {
	signal1 := make(chan os.Signal, 1)
	signal.Notify(signal1, os.Interrupt)
	go func() {
		for range signal1 {
			utils.Logger.Println("Bye bye")
			os.Exit(0)
		}
	}()
}

func loadLogger() {
	utils.Logger.SetPrefix(LoggerPrefix)
}

//主HTTP服务
// func initMainHTTP() *gin.Engine {
// 	hp := gin.Default()
// 	routers.StartPage(hp)
// 	return hp
// }

//RunDBCommand 数据库迁移
func RunDBCommand() {
	models.AddTables()
	models.Migrate(models.TableList)
}
