package router

import (
	"best-practics/common/config"
	"best-practics/common/initialize/log"
	"best-practics/common/middleware"
	"best-practics/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由
// Server holds the dependencies for a HTTP server.
type Server struct {
	UserController  *interfaces.UserController
}

func Init(server *Server) *gin.Engine {
	var r = gin.New()
	r.Use(middleware.SetLoggerMiddleware(),middleware.GinRecovery(true))
	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开

	// 注册系统路由
	InitSysRouter(r,server)

	log.Info("router register success")
	return r
}

func InitSysRouter(r *gin.Engine,server *Server) *gin.RouterGroup {
	g := r.Group("")
	baseRouter(g)
	// 静态文件
	staticFileRouter(g)
	// swagger；注意：生产环境可以注释掉
	swaggerRouter(g)
	// 无需认证
	noCheckRoleRouter(g,server)
	// 需要认证
	//checkRoleRouterInit(g, authMiddleware)
	return g
}


func baseRouter(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "hello!")
	})
	// 健康监测
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
}

func staticFileRouter(r *gin.RouterGroup) {
	r.StaticFS(config.ConfigCenter.Local.Path, http.Dir(config.ConfigCenter.Local.Path)) // 为文件提供本地静态地址
}

func swaggerRouter(r *gin.RouterGroup) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func noCheckRoleRouter(r *gin.RouterGroup,server *Server) {
	v1 := r.Group("/api/v1")
	//registerUserRouter
	{
		v1.GET("/user/:id", server.UserController.GetUser)
		v1.POST("/user", server.UserController.CreateUser)
	}
}

