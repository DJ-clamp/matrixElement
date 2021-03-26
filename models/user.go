package models

import (
	"strconv"
	"time"
)

type User struct {
	BaseModel
	Name        string
	Status      int
	ActivatedAt time.Time
}

// 设置表名为user
func (User) TableName() string {
	return "user"
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

//添加数据
func (user *User) Create() error {
	return DB.Create(user).Error
}

func (user *User) Update(column map[string]interface{}) error {
	return DB.Model(user).Updates(column).Error
}

func (user *User) Delete() error {
	return DB.Delete(user).Error
}
