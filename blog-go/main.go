package main

import (
	"best-practics/common"
	"best-practics/common/global"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.Viper = common.InitViper()      // 初始化Viper
	global.Logger = common.InitZap()       // 初始化zap日志库

}
