package blog

import (
	"github.com/gin-gonic/gin"
)

type BlogRouter struct {
}

func (e *BlogRouter) InitBlogRouter(Router *gin.RouterGroup) {
	routerGroup := Router.Group("blog")
	{
		routerGroup.GET("", func(c *gin.Context) {
			// c.JSON：返回JSON格式的数据
			c.JSON(200, gin.H{
				"message": "Hello blog!",
			}) // 导入Excel
		})
	}
}
