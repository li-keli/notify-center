package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notify-center/pkg/db"
	"notify-center/pkg/redis"
	v1 "notify-center/server/api/v1"
)

func main() {
	db.NewDB()
	redis.NewRedisConn()

	engine := gin.Default()
	engine.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "up")
	})
	engine.POST("/v1/terminal/register", v1.RegisterTerminal)
	engine.POST("/v1/terminal/unRegister", v1.UnRegisterTerminal)
	engine.POST("/v1/notification/send", v1.Notify)

	_ = engine.Run()
}
