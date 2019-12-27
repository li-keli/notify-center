package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"notify-center/pkg/db"
	"notify-center/server/api/v1/vo"
	"notify-center/server/logic"
)

// 注册终端
func RegisterTerminal(ctx *gin.Context) {
	var input vo.RegisterTerminalVo

	_ = ctx.BindJSON(&input)
	logrus.Println(input)

	ctx.AsciiJSON(http.StatusOK, input)
}

// 注销终端
func UnRegisterTerminal(ctx *gin.Context) {
	var input vo.UnRegisterTerminalVo

	_ = ctx.BindJSON(&input)
	logrus.Println(input)

	ctx.AsciiJSON(http.StatusOK, input)
}

// 发送推送消息
func Notify(ctx *gin.Context) {
	var input vo.NotifyVo

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error(err.Error()))
		return
	}

	// 获取推送目标数据
	entity, err := (&db.AppEntity{}).FindAppEntityByJsjId(input.JsjUniqueId)
	if err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error(err.Error()))
		return
	}

	// 获取推送目标平台配置
	config, err := (&db.DicConfigEntity{}).FindConfig(entity.TargetType)
	if err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error(err.Error()))
		return
	}

	if err := logic.BuildPushActuator(input, entity, config).PushMessage(entity.PushToken); err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, vo.BaseOutput{}.Success())
}
