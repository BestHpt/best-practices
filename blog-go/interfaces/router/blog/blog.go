package blog

import (
	"best-practics/common"
	"best-practics/utils/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BlogRouter struct {
}

func (e *BlogRouter) InitBlogRouter(Router *gin.RouterGroup) {
	routerGroup := Router.Group("blog")
	{
		routerGroup.GET("", func(c *gin.Context) {
			ctx := common.GetTraceCtx(c)
			// c.JSON：返回JSON格式的数据
			log.InfoF(ctx,"TEST_TAG",zap.String("test","666"))
			log.WarnF(ctx,"TEST_TAG",zap.String("Tag","FileNotFind"))
			c.JSON(200, gin.H{
				"message": "Hello blog!",
			}) // 导入Excel
		})
	}
}
