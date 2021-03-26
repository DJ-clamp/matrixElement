package routers

import (
	"net/http"
	"strings"

	"github.com/DJ-clamp/matrixElement/models"
	"github.com/DJ-clamp/matrixElement/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  " it`s works!",
	})
}
func GetUserById(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "GetUserById",
		"name": name,
	})
}

var db *gorm.DB

func AddUser(c *gin.Context) {
	name := c.PostForm("name")
	user := models.User{Name: name}
	db.Create(user)
	if err := user.Create(); err != nil {
		utils.Logger.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"type":   "AddUser",
		"user":   name,
	})
}

func AddUsers(c *gin.Context) {
	names := c.PostForm("name")
	nameArr := strings.Split(names, "\n")
	for _, name := range nameArr {
		utils.Logger.Println(name)
		user := models.User{Name: name}
		db.Create(user)
		if err := user.Create(); err != nil {
			utils.Logger.Println(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"type":   "AddUsers",
		"length": len(nameArr),
	})
}

func UpdateUser(c *gin.Context) {
	user := c.PostForm("user")
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"type":   "UpdateUser",
		"user":   user,
	})
}

func DeleteUser(c *gin.Context) {
	user := c.PostForm("user")
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"type":   "DeleteUser",
		"user":   user,
	})
}
