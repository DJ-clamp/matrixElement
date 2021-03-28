package models

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string
	Status      uint
	ActivatedAt time.Time
}

// 设置表名为user
func (User) TableName() string {
	return "user"
}

func (user User) GetUsers() ([]User, error) {
	var users []User
	result := DB.Find(&users)
	return users, result.Error
}

func (user User) GetUsersWithoutUsed(count int) ([]User, error) {
	var users []User
	result := DB.Where("status = ?", 0).Limit(count).Find(&users)
	return users, result.Error
}

func (user User) GetUsersWithStatus(status int) ([]User, error) {
	var users []User
	result := DB.Where("status = ?", status).Find(&users)
	return users, result.Error
}

func (User) GetUserDataById(id string) (*User, error) {
	pid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	var data User
	err = DB.First(&data, "id = ?", pid).Error
	return &data, err
}
func (User) GetUserDataByName(name string) (*User, error) {
	var data User
	err := DB.First(&data, "name = ?", name).Error
	return &data, err
}

//添加数据
func (user *User) Create() error {
	return DB.Create(user).Error
}
func (user *User) CreateAll(users []User) error {
	return DB.Create(users).Error
}
func (user *User) Update(column interface{}) error {
	return DB.Save(column).Error
}

func (user *User) Delete() error {
	return DB.Delete(user).Error
}
