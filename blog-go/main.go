package main

import (
	"best-practics/common"
	"best-practics/common/initialize"
	"best-practics/utils/log"
	"fmt"
	"go.uber.org/zap"
	"time"
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
	common.Viper = initialize.InitViper()
	//2、初始化zap日志库
	log.InitZap()
	//3、gorm连接数据库
	common.GVA_DB = initialize.InitGorm()
	if common.GVA_DB != nil {
		initialize.MysqlTables(common.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := common.GVA_DB.DB()
		defer db.Close()
	}
	//4、设置routers
	Router := initialize.Routers()

	//5、初始化gin server
	address := fmt.Sprintf(":%d", common.GlobalConfig.System.Addr)
	s := initialize.InitServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	common.Logger.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 best-practices
	当前版本:V0.0.1 Golang
    作者:微信号：bestbear666 公众号：简凡丶
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
	默认后端测试路径:http://127.0.0.1%s/blog
`, address, address)
	common.Logger.Error(s.ListenAndServe().Error())

}


