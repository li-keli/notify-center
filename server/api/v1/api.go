package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notify-center/pkg/constant"
	"notify-center/pkg/db"
	"notify-center/pkg/redis"
	"notify-center/pkg/track_log"
	"notify-center/server/api/v1/vo"
	"notify-center/server/logic"
	"strconv"
	"time"
)

func RegisterNotify(engine *gin.Engine, logMiddle func(ctx *gin.Context)) {
	v1 := engine.Group(`/v1`, logMiddle)
	{
		// 注册
		v1.POST("/terminal/register", TerminalRegister)
		v1.POST("/terminal/unRegister", TerminalUnRegister)

		// 下发通知
		v1.POST(`/wechat/send`, WeChatNotify)
		v1.POST("/notification/send", Notify)

		// 获取历史消息
		v1.POST("/msg", MessageList)
	}
}

// 注册终端
func TerminalRegister(ctx *gin.Context) {
	var (
		input    vo.RegisterTerminalInputVo
		trackLog = track_log.Logger(ctx)
	)

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error(err.Error()))
		trackLog.Panic("参数异常")
	}
	trackLog.Info(input)

	register := db.NotifyRegister{
		JsjUniqueId:      input.JsjUniqueId,
		PushToken:        input.PushToken,
		PlatformTypeId:   input.PlatformType,
		PlatformTypeName: constant.PlatformTypeValueOf(input.PlatformType),
		TargetTypeId:     input.TargetType,
		TargetTypeName:   constant.TargetTypeValueOf(input.TargetType),
		CreateTime:       time.Now(),
		UpdateTime:       time.Now(),
	}

	if one, _ := register.FindOne(register.JsjUniqueId); one.JsjUniqueId != 0 {
		trackLog.Info("更新")
		if err := register.UpdateSpecify(register); err != nil {
			ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error("更新异常"))
			trackLog.Panic("数段注册更新异常", err)
		}
	} else {
		trackLog.Info("新增")
		if err := register.Insert(register); err != nil {
			ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error("更新异常"))
			trackLog.Panic("终端注册入库异常", err)
		}
	}
}

// 注销终端
func TerminalUnRegister(ctx *gin.Context) {
	var (
		input    vo.UnRegisterTerminalInputVo
		trackLog = track_log.Logger(ctx)
	)

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error(err.Error()))
		trackLog.Panic("参数异常")
	}
	trackLog.Info(input)

	err := db.NotifyRegister{}.DelOne(input.JsjUniqueId, input.PlatformType, input.TargetType)
	if err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error(err.Error()))
		trackLog.Panic("终端注销异常")
	}
}

// 发送推送消息
func Notify(ctx *gin.Context) {
	var (
		input     vo.NotifyVo
		nConfig   db.NotifyConfig
		nRegister db.NotifyRegister
		nMessage  db.NotifyMsg

		trackLog = track_log.Logger(ctx)
	)
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error(err.Error()))
		return
	}
	trackLog.Info("请求数据：", input)

	// 尝试进行WebSocket推送
	if all, _ := redis.GetHashAll(strconv.Itoa(input.JsjUniqueId)); len(all) > 0 {
		err := logic.PushWSocket{NotifyVo: input}.PushMessage()
		if err == nil {
			ctx.JSON(http.StatusOK, vo.BaseOutput{}.Success("Socket推送成功"))
			return
		} else {
			trackLog.Warn("WebSocket实时推送失败")
		}
	} else {
		trackLog.Warn("WebSocket通道不存在")
	}

	// 获取推送目标的数据
	one, err := nRegister.FindOne(input.JsjUniqueId)
	if err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error("未发现终端注册信息"))
		return
	}
	trackLog.Info("推送设备注册数据：", one)

	// 停用微信小程序模板推送
	if one.PlatformTypeId == constant.MiniProgram {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Success("小程序模板推送暂时下线，按照微信新规停用调整"))
		return
	}

	// 获取推送目标平台配置
	config, err := nConfig.FindOne(one.PlatformTypeId, input.TargetType)
	if err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error("未发现接收平台配置信息"))
		return
	}
	trackLog.Info("推送平台配置：", config)

	// 记录推送数据
	nMessage.Insert(db.NotifyMsg{
		JsjUniqueId:      one.JsjUniqueId,
		PushToken:        one.PushToken,
		PlatformTypeId:   one.PlatformTypeId,
		PlatformTypeName: one.PlatformTypeName,
		TargetTypeId:     input.TargetType,
		TargetTypeName:   constant.TargetTypeValueOf(input.TargetType),
		DataContent:      input.DataToStr(),
		GroupName:        input.GroupName,
		CreateTime:       time.Now(),
	})

	// 发起推送
	offlinePushActuator := logic.BuildPushActuator(ctx, input, one, config)
	if err := offlinePushActuator.PushMessage(one.PushToken); err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error(offlinePushActuator.Mode()+"平台推送失败："+err.Error()))
		return
	}

	trackLog.Infof("%s平台推送成功", offlinePushActuator.Mode())
	ctx.JSON(http.StatusOK, vo.BaseOutput{}.Success(offlinePushActuator.Mode()+"平台推送成功"))
}

// 微信小程序订阅消息下发
func WeChatNotify(ctx *gin.Context) {
	//var (
	//	input     vo.NotifyVo
	//	nConfig   db.NotifyConfig
	//	nRegister db.NotifyRegister
	//	nMessage  db.NotifyMsg
	//
	//	trackLog = track_log.Logger(ctx)
	//)
}

// 消息列表
func MessageList(ctx *gin.Context) {
	var (
		input    vo.MsgListInputVo
		trackLog = track_log.Logger(ctx)
	)

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error(err.Error()))
		return
	}
	trackLog.Info("请求数据：", input)

	if input.JsjUniqueId == 0 && input.PushToken == "" {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error("JsjUniqueId和PushToken不能同时为空"))
		return
	}

	all, err := db.NotifyMsg{}.FindAll(input.JsjUniqueId, input.PushToken, input.Offset, input.Limit)
	if err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, vo.MsgListOutputVo{
		BaseOutput: vo.BaseOutput{}.Success(""),
		Msg:        all,
	})
}
