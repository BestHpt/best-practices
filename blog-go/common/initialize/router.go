package initialize

import (
	"best-practics/common/global"
	"best-practics/interfaces/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()

	Router.StaticFS(global.GlobalConfig.Local.Path, http.Dir(global.GlobalConfig.Local.Path)) // 为文件提供本地静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.Logger.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.Logger.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	//获取路由组实例
	blogRouter := router.RouterGroupApp.Blog
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}

	//PrivateGroup := Router.Group("")
	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		blogRouter.InitBlogRouter(PublicGroup)                      // 注册BLOG相关路由
	}

	global.Logger.Info("router register success")
	return Router
}
