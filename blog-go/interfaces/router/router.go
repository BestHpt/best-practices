package router

import (
	"best-practics/common/config"
	"best-practics/common/initialize/log"
	"best-practics/interfaces"
	"best-practics/interfaces/middleware"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

type Server struct {
	UserController *interfaces.UserController
}

func Init(server *Server) *gin.Engine {
	var r = gin.New()
	r.Use(middleware.SetLoggerMiddleware(), middleware.GinRecovery(true))
	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开

	// 注册系统路由
	InitRouter(r, server)

	log.Info("router register success")
	return r
}

func InitRouter(r *gin.Engine, server *Server) *gin.RouterGroup {
	g := r.Group("")
	//1、健康监测
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, "ok")
	})

	//2、为文件提供本地静态地址
	r.StaticFS(config.ConfigCenter.Local.Path, http.Dir(config.ConfigCenter.Local.Path))

	//3、swagger；注意：生产环境可以注释掉
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//4、User相关
	v1 := r.Group("/api/v1")
	{
		v1.GET("/user/:id", server.UserController.GetUser)
		v1.POST("/user", server.UserController.CreateUser)
	}
	return g
}
