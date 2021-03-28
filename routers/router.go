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
		"type": "GetUserById",
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
		//changed status = 2 for users
		for _, user := range users {
			user.Status = 2
			if err = db.Update(user); err != nil {
				utils.Logger.Println(err)
			}
		}
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
	if len(name) != 11 {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusBadRequest,
			"type":   "AddUser",
			"msg":    "error",
			"length": 0,
			"data":   "",
		})
		return
	}
	user := models.User{Name: name, Status: 0, ActivatedAt: time.Now()}
	data, isok := user.GetUserDataByName(name)
	if isok != nil {
		utils.Logger.Println(isok)
		if err := user.Create(); err != nil {
			utils.Logger.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"type":   "AddUser",
			"msg":    "",
			"length": 1,
			"data":   user,
		})

	} else {
		if data != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusCreated,
				"type": "AddUser",
				"msg":  "the use has existed",
				"data": data,
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
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusNotFound,
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
			if len(u) != 11 {
				continue
			}
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
			"code":   http.StatusNotFound,
			"type":   "AddUsers",
			"length": len(tmp),
			"msg":    "error",
			// "user": nameList.
		})
		return
	}
	var newUsers []models.User
	if len(tmp) <= 10 {
		for _, n := range tmp {
			newUsers = append(newUsers, models.User{Name: n, Status: 0, ActivatedAt: time.Now()})
		}
		if err := db.CreateAll(newUsers); err != nil {
			utils.Logger.Println(err)
			return
		}
	} else {
		go func() {
			for _, n := range tmp {
				newUsers := models.User{Name: n, Status: 0, ActivatedAt: time.Now()}
				if err := newUsers.Create(); err != nil {
					utils.Logger.Println(err)
				}
			}
		}()
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"type":   "AddUsers",
			"length": len(tmp),
			"msg":    "success",
		})
	}

}

//更新状态
func UpdateUser(c *gin.Context) {
	name := c.PostForm("name")
	var db = models.User{}
	user, err := db.GetUserDataByName(name)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"type": "UpdateUser",
			"msg":  err,
			"user": user,
		})
		return
	}
	user.Status = 1
	if err := db.Update(user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"type": "UpdateUser",
			"msg":  err,
			"user": user,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"type": "UpdateUser",
			"msg":  "success",
			"user": user,
		})
	}
}

func DeleteUser(c *gin.Context) {
	user := c.PostForm("name")
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"type": "DeleteUser",
		"msg":  "success",
		"user": user,
	})
}

func ResetUsers(c *gin.Context) {
	status := c.PostForm("status")
	i, _ := strconv.Atoi(status)
	if status != "" {
		var db = models.User{}
		users, err := db.GetUsersWithStatus(i)
		if err != nil {
			utils.Logger.Println(err)
		}
		for _, user := range users {
			user.Status = 1
			if err := db.Update(user); err != nil {
				utils.Logger.Println(err)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"type":   "ResetUsers",
			"msg":    "success",
			"status": status,
		})
	}
}
