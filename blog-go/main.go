package main

import (
	"best-practics/common"
	"best-practics/common/config"
	"best-practics/common/initialize/log"
	"best-practics/common/initialize/mysql"
	"best-practics/common/initialize/router"
	"best-practics/common/initialize/server"
	"best-practics/common/initialize/viper"
	"fmt"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download


// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @termsOfService https://besthpt.github.io/

// @contact.name besthpt.github.io
// @contact.url 948748073@qq.com
// @contact.email 948748073@qq.com
// @BasePath /
func main() {
	//1、初始化Viper
	common.Viper = viper.Init()
	//2、初始化zap日志库
	log.InitZap()
	//3、gorm连接数据库
	common.GVA_DB = mysql.Init()
	if common.GVA_DB != nil {
		mysql.MysqlTables(common.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := common.GVA_DB.DB()
		defer db.Close()
	}
	//4、设置routers
	Router := router.Init()
	//5、初始化gin server
	address := fmt.Sprintf(":%d", config.ConfigCenter.System.Addr)
	s := server.Init(address, Router)
	log.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 best-practices
	当前版本:V0.0.1 Golang
    作者:微信号：bestbear666 公众号：简凡丶
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
	默认后端测试路径:http://127.0.0.1%s/blog
`, address, address)
	log.Error(s.ListenAndServe().Error())

}


