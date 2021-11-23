package blog

import (
	"best-practics/common"
	"best-practics/common/c_error"
	"best-practics/utils/log"
	"best-practics/utils/response"
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
			log.WarnF(ctx,c_error.BookNotFoundError.String(),zap.String("test","666"))
			response.SuccessWithMessage(c,"Hello blog!")
		})
	}
}
