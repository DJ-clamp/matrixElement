package routers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/DJ-clamp/matrixElement/models"
	"github.com/DJ-clamp/matrixElement/utils"
	"github.com/gin-gonic/gin"
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

func GetUsers(c *gin.Context) {
	count := c.Query("count")
	i, _ := strconv.Atoi(count)
	var db = models.User{}
	if users, err := db.GetUsersWithoutUsed(i); err != nil {
		utils.Logger.Println(err)
	} else {
		c.JSON(200, gin.H{
			"code":   200,
			"type":   "GetUsers",
			"length": len(users),
			"data":   users,
		})
	}

}

func AddUser(c *gin.Context) {
	name := c.PostForm("name")
	user := models.User{Name: name, Status: 0, ActivatedAt: time.Now()}
	if data, err := user.GetUserDataByName(name); err != nil {
		utils.Logger.Fatalln(err)
		return
	} else {
		if data != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusCreated,
				"type":   "AddUser",
				"msg":    "the use has existed",
				"user":   data,
			})
		} else {

			if err := user.Create(); err != nil {
				utils.Logger.Println(err)
			}
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"type":   "AddUser",
				"user":   user,
			})
		}
	}
}

type User models.User

// func saveUser(c *gin.Context, data interface{}, saveUserFunc func(models.User)) {

// }

func AddUsers(c *gin.Context) {
	// var Contains = func(l *list.List, value string) (bool, *list.Element) {
	// 	for e := l.Front(); e != nil; e = e.Next() {
	// 		if e.Value == value {
	// 			return true, e
	// 		}
	// 	}
	// 	return false, nil
	// }
	names := c.PostForm("name")
	nameArr := strings.Split(names, "\n")
	// nameList := list.New()
	// for _, name := range nameArr {
	// 	nameList.PushBack(name)
	// }

	var (
		users []models.User
		err   error
	)
	var db = models.User{}
	//处理数据
	if users, err = db.GetUsers(); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusNotFound,
			"type":   "AddUsers",
			"length": 0,
			"msg":    "error",
			// "user": nameList.
		})

		return
	}

	var tmp []string
	if len(users) > 0 {
		// for _, u := range users {
		// 	if contain, e := Contains(nameList, u.Name); contain {
		// 		nameList.Remove(e)
		// 	}
		// }

		for _, u := range nameArr {
			var isContained bool = false
			for _, user := range users {
				if u == user.Name {
					isContained = true
				}
			}
			if !isContained {
				tmp = append(tmp, u)
			}
		}
	} else {
		tmp = append(tmp, nameArr...)
	}
	if len(tmp) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusNotFound,
			"type":   "AddUsers",
			"length": len(tmp),
			"msg":    "error",
			// "user": nameList.
		})
		return
	}
	var newUsers []models.User
	for _, n := range tmp {
		newUsers = append(newUsers, models.User{Name: n, Status: 0, ActivatedAt: time.Now()})
	}
	if err := db.CreateAll(newUsers); err != nil {
		utils.Logger.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"type":   "AddUsers",
		"length": len(nameArr),
		"msg":    "success",
	})
}

//更新状态
func UpdateUser(c *gin.Context) {
	name := c.PostForm("name")
	var db = models.User{}
	user, err := db.GetUserDataByName(name)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"type":   "UpdateUser",
			"msg":    "error",
			"user":   user,
		})
		return
	}
	user.Status = 1
	db.Update(user)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"type":   "UpdateUser",
		"msg":    "success",
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
