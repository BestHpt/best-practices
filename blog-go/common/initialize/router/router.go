/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package router

import (
	"best-practics/common/config"
	"best-practics/common/initialize/log"
	"best-practics/common/middleware"
	"best-practics/interfaces/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Init() *gin.Engine {
	var Router = gin.New()
	Router.Use(middleware.SetLoggerMiddleware(),middleware.GinRecovery(true))
	Router.StaticFS(config.ConfigCenter.Local.Path, http.Dir(config.ConfigCenter.Local.Path)) // 为文件提供本地静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	log.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Info("register swagger handler")
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

	log.Info("router register success")
	return Router
}
