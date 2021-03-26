package main

import (
	"log"
	"os"

	"github.com/DJ-clamp/matrixElement/routers"
	"github.com/gin-gonic/gin"
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

}

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

//主HTTP服务
func initMainHTTP() *gin.Engine {
	hp := gin.Default()
	hp.GET("/", routers.Index)
	return hp
}
