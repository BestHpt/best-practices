package gin_server

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

type Server interface {
	ListenAndServe() error
}

func Init(address string, router *gin.Engine) Server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
