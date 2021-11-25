package blog

import (
	"best-practics/common"
	"best-practics/common/c_error"
	"best-practics/utils/log"
	"best-practics/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BlogRouter struct {
}

func (e *BlogRouter) InitBlogRouter(Router *gin.RouterGroup) {
	routerGroup := Router.Group("blog")
	{
		routerGroup.GET("", HandleHello)
	}
}

// @Tags 文章操作
// @Summary 获取文章信息
// @Accept json
// @Produce  application/json
// @Param who query string true "人名"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Failure 400 {string} string "{"msg": "who are you"}"
// @Router /blog [get]
func HandleHello(c *gin.Context) {
	ctx := common.GetTraceCtx(c)
	// c.JSON：返回JSON格式的数据
	log.WarnF(ctx,c_error.BookNotFoundError.String(),zap.String("test","666"))
	who := c.Query("who")
	response.SuccessWithMessage(c,fmt.Sprintf("Hello %s!",who))
}
