package middleware

import (
	"best-practics/common/consts"
	"best-practics/common/initialize/log"
	"best-practics/common/trace"
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)


func SetLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		uuidStr := strings.ReplaceAll(uuid.New().String(), "-", "")
		path := c.Request.URL.Path
		userId := c.GetInt("user_id")
		ctx := context.WithValue(context.Background(), consts.TraceKey, &trace.Trace{TraceId: uuidStr, Caller: path, UserId: userId})
		c.Set(consts.TraceCtx,ctx)

		c.Next()
		cost := time.Since(start)
		log.Info("_com_request_info",
			zap.Int("Status", c.Writer.Status()),
			zap.String("Method", c.Request.Method),
			zap.String("IP",c.ClientIP()),
			zap.String("Path",path),
			zap.String("TraceId", uuidStr),
			zap.Int("UserId", userId),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("UserAgent",c.Request.UserAgent()),
			zap.Duration("Cost",cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					log.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					log.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					log.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
