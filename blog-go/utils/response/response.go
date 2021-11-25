package response

import (
	"best-practics/common/consts"
	"best-practics/common/initialize/log"
	trace2 "best-practics/common/trace"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)
const (
	SUCCESS = 0
)

func Result(c *gin.Context, errCode int, msg string, data interface{}) {
	ctx := trace2.GetTraceCtx(c)
	trace := ctx.Value(consts.TraceKey).(*trace2.Trace)
	log.Info("_com_request_out",
		zap.String("trace_id", trace.TraceId),
		zap.Int("user_id", trace.UserId),
		zap.Int("errCode", errCode),
		zap.String("msg", msg),
		zap.Any("data", data),
	)
	//c.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	c.JSON(http.StatusOK, gin.H{
		"code": errCode,
		"msg":  msg,
		"data": data,
	})
}

// 将json字符窜以标准json格式返回（例如，从redis读取json、格式的字符串，返回给浏览器json格式）
func ReturnJsonFromString(Context *gin.Context, httpCode int, jsonStr string) {
	Context.Header("Content-Type", "application/json; charset=utf-8")
	Context.String(httpCode, jsonStr)
}

// 语法糖函数封装
func Success(c *gin.Context) {
	Result(c, SUCCESS, "success", "")
}

func SuccessWithData(c *gin.Context, data interface{}) {
	Result(c, SUCCESS, "success", data)
}

func SuccessWithMessage(c *gin.Context,message string) {
	Result(c, SUCCESS, message, map[string]interface{}{})
}

// 失败的业务逻辑
func Fail(c *gin.Context, dataCode int, msg string) {
	Result(c, dataCode, msg, map[string]interface{}{})
	c.Abort()
}
