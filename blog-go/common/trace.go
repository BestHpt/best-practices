/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package common

import (
	"best-practics/common/consts"
	"context"
	"github.com/gin-gonic/gin"
)

type Trace struct {
	TraceId   string  `json:"trace_id"`
	SpanId    string  `json:"span_id"`
	Caller    string  `json:"caller"`
	SrcMethod *string `json:"srcMethod,omitempty"`
	UserId    int   `json:"user_id"`
}


func GetTraceCtx(c *gin.Context) context.Context {
	return c.MustGet(consts.TraceCtx).(context.Context)
}

