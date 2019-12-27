package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"notify-center/pkg/redis"
	v1 "notify-center/server/api/v1"
)

func main() {
	logrus.Println("run...")
	redis.NewRedisConn()

	engine := gin.Default()
	engine.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "up")
	})
	engine.POST("/v1/terminal/registerKt", v1.RegisterKt)
	engine.GET("/notify", v1.Notify)

	_ = engine.Run("localhost:8080")
}
