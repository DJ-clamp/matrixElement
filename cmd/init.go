package main

import (
	"flag"
	"log"
	"os"

	"github.com/DJ-clamp/matrixElement/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	HTTP_PORT string = ""
	DB_NAME   string = ""
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}
	HTTP_PORT = os.Getenv("HTTP_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	isAutoMigrate := flag.Bool("migrate", false, "auto migrate database")
	flag.Parse()
	initDB()
	if *isAutoMigrate {
		defer os.Exit(0)
		RunDBCommand()
	}

}

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(models.User{})
	return db
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
