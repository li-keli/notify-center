package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"notify-center/pkg/db"
	"notify-center/pkg/redis"
	"notify-center/server/api/v1/vo"
	"notify-center/server/logic"
	"strconv"
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
	logrus.Info("请求数据：%v", input)

	// 获取推送目标平台配置
	config, err := (&db.DicConfigEntity{}).FindConfig(input.TargetType)
	if err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error(err.Error()))
		return
	}
	logrus.Info("获取推送目标平台配置：%v", config)

	// 尝试进行WebSocket推送
	if all, _ := redis.GetHashAll(strconv.Itoa(input.JsjUniqueId)); len(all) > 0 {
		err = logic.PushWSocket{NotifyVo: input}.PushMessage()
		if err == nil {
			ctx.JSON(http.StatusOK, vo.BaseOutput{}.Success())
			return
		}
	}

	// 获取推送目标数据
	entity, err := (&db.AppEntity{}).FindAppEntityByJsjId(input.JsjUniqueId)
	if err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error(err.Error()))
		return
	}
	logrus.Info("获取推送目标数据：%v", entity)

	// 发起推送
	if err := logic.BuildPushActuator(input, entity, config).PushMessage(entity.PushToken); err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error(err.Error()))
		return
	}

	logrus.Info("推送成功")
	ctx.JSON(http.StatusOK, vo.BaseOutput{}.Success())
}
