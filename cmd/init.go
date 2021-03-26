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

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}
	HTTP_PORT = os.Getenv("HTTP_PORT")

}

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
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
