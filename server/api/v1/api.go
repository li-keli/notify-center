package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"notify-center/pkg/constant"
	"notify-center/pkg/db"
	"notify-center/pkg/redis"
	"notify-center/server/api/v1/vo"
	"notify-center/server/logic"
	"strconv"
	"time"
)

// 注册终端
func RegisterTerminal(ctx *gin.Context) {
	var input vo.RegisterTerminalVo

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error(err.Error()))
		logrus.Panic("参数异常")
	}
	logrus.Info(input)

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
		logrus.Info("更新")
		if err := register.UpdateSpecify(register); err != nil {
			ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error("更新异常"))
			logrus.Panic("数段注册更新异常", err)
		}
	} else {
		logrus.Info("新增")
		if err := register.Insert(register); err != nil {
			ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error("更新异常"))
			logrus.Panic("终端注册入库异常", err)
		}
	}

	ctx.JSON(http.StatusOK, vo.BaseOutput{}.Success(""))
}

// 注销终端
func UnRegisterTerminal(ctx *gin.Context) {
	var input vo.UnRegisterTerminalVo

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error(err.Error()))
		logrus.Panic("参数异常")
	}
	logrus.Println(input)

	err := db.NotifyRegister{}.DelOne(input.JsjUniqueId, input.PlatformType, input.TargetType)
	if err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error(err.Error()))
		logrus.Panic("终端注销异常")
	}

	ctx.JSON(http.StatusOK, vo.BaseOutput{}.Success(""))
}

// 发送推送消息
func Notify(ctx *gin.Context) {
	var (
		input     vo.NotifyVo
		nConfig   db.NotifyConfig
		nRegister db.NotifyRegister
		nMessage  db.NotifyMsg
	)

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error(err.Error()))
		return
	}
	logrus.Info("请求数据：", input)

	// 尝试进行WebSocket推送
	if all, _ := redis.GetHashAll(strconv.Itoa(input.JsjUniqueId)); len(all) > 0 {
		err := logic.PushWSocket{NotifyVo: input}.PushMessage()
		if err == nil {
			ctx.JSON(http.StatusOK, vo.BaseOutput{}.Success("Socket推送成功"))
			return
		}
	}

	// 获取推送目标的数据
	one, err := nRegister.FindOne(input.JsjUniqueId)
	if err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error("未发现终端注册信息"))
		return
	}
	logrus.Infof("获取推送目标数据：%#v", one)

	// 获取推送目标平台配置
	config, err := nConfig.FindOne(one.PlatformTypeId, input.TargetType)
	if err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error("未发现接收平台配置信息"))
		return
	}
	logrus.Infof("获取推送目标平台配置：%#v", config)

	// 记录推送数据
	nMessage.Insert(db.NotifyMsg{
		PushToken:        one.PushToken,
		PlatformTypeId:   one.PlatformTypeId,
		PlatformTypeName: one.PlatformTypeName,
		TargetTypeId:     input.TargetType,
		TargetTypeName:   constant.TargetTypeValueOf(input.TargetType),
		DataContent:      input.DataToStr(),
		CreateTime:       time.Now(),
	})

	// 发起推送
	if err := logic.BuildPushActuator(input, one, config).PushMessage(one.PushToken); err != nil {
		ctx.JSON(http.StatusOK, vo.BaseOutput{}.Error(err.Error()))
		return
	}

	logrus.Info("离线推送成功")
	ctx.JSON(http.StatusOK, vo.BaseOutput{}.Success("离线推送成功"))
}
