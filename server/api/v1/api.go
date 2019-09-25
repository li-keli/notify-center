package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"notify-center/server/api/v1/vo"
)

// 空铁小程序注册
func RegisterKt(ctx *gin.Context) {
	var input vo.RegisterKtInput

	_ = ctx.BindJSON(&input)
	logrus.Println(input)

	ctx.AsciiJSON(http.StatusOK, input)
}
