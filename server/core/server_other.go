//go:build !windows
// +build !windows

package core

import (
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

// endless用于实现 无缝重启 的 HTTP 服务器。它的主要功能是在不中断现有连接的情况下，重新启动或升级服务器。
// 这对于生产环境中的服务更新非常有用，可以避免停机时间。
func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Minute
	s.WriteTimeout = 10 * time.Minute
	s.MaxHeaderBytes = 1 << 20
	return s
}
