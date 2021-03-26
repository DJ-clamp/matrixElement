package routers

import "github.com/gin-gonic/gin"

func StartPage(router *gin.Engine) {
	router.GET("/", Index)
}
