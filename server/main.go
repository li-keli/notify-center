// 推送处理器，集成WebSocket、Apns、极光实现多方的推送
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notify-center/pkg/db"
	"notify-center/pkg/redis"
	"notify-center/pkg/track_log"
	v1 "notify-center/server/api/v1"
)

func main() {
	db.NewDB()
	redis.NewRedisConn()

	engine := gin.Default()
	track_log.UseLogMiddle(engine)
	engine.GET("/health", func(ctx *gin.Context) { ctx.String(http.StatusOK, "up") })
	v1.RegisterNotify(engine)

	_ = engine.Run()
}
