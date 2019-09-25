package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	v1 "notify-center/server/api/v1"
)

func main() {
	logrus.Println("run...")

	engine := gin.Default()
	engine.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "up")
	})
	engine.POST("/v1/terminal/registerKt", v1.RegisterKt)
	_ = engine.Run()
}
