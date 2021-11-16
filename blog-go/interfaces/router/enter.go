package router

import (
	"best-practics/interfaces/router/blog"
)

type RouterGroup struct {
	Blog   blog.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
