package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"notify-center/pkg/db"
	"notify-center/pkg/redis"
	v1 "notify-center/server/api/v1"
)

func main() {
	logrus.Println("run...")

	db.NewDB()
	redis.NewRedisConn()

	engine := gin.Default()
	engine.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "up")
	})
	engine.POST("/v2/terminal/register", v1.RegisterTerminal)
	engine.POST("/v2/terminal/unRegister", v1.UnRegisterTerminal)
	engine.POST("/v2/notification/send", v1.Notify)

	_ = engine.Run(":8081")
}
