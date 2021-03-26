package routers

import (
	"github.com/gin-gonic/gin"
)

func StartPage(router *gin.Engine) {
	router.GET("/", Index)
	router.GET("/user/:name", GetUserById)
	router.GET("/users", GetUsers)
	router.POST("/add", AddUser)
	router.POST("/addAll", AddUsers)
	router.POST("/update", UpdateUser)
	router.GET("/delete", DeleteUser)
}
