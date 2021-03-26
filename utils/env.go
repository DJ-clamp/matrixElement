package utils

import (
	"os"
)

//aime全局变量
type Am struct {
	Env //配置信息
}

//环境变量
type Env map[string]string

//环境设置接口
type IEnv interface {
	GetEnv(string, string) string
	SetEnv(string, string)
}

//Main parameter in global aimenet`s framework
var Global = Am{
	Env: Env{},
}

func AppGlobeInit() {
	//初始化全局变量Am
}

//封装环境变量获取功能,可解耦带接口实现
//获取环境变量
func (App *Am) GetEnv(key, defaultValue string) string {

	if tmp := App.Env[key]; tmp != "" {
		return tmp
	}
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	App.Env[key] = value
	return value
}

//设置零时变量
func (App *Am) SetEnv(key string, st string) {
	App.Env[key] = st
}
