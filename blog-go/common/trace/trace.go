package trace

import (
	"best-practics/common/consts"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

type Trace struct {
	TraceId   string  `json:"trace_id"`
	SpanId    string  `json:"span_id"`
	Caller    string  `json:"caller"`
	SrcMethod *string `json:"srcMethod,omitempty"`
	UserId    int     `json:"user_id"`
}

func GetTraceCtx(c *gin.Context) context.Context {
	if value, ok := c.Get(consts.TraceCtx); ok {
		return value.(context.Context)
	} else {
		uuidStr := strings.ReplaceAll(uuid.New().String(), "-", "")
		return context.WithValue(context.Background(), consts.TraceKey, &Trace{TraceId: uuidStr})
	}
}
