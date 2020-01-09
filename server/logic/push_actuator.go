package logic

import (
	"github.com/gin-gonic/gin"
	"notify-center/pkg/constant"
	"notify-center/pkg/db"
	"notify-center/server/api/v1/vo"
)

// 推送
type PushActuator interface {
	// 获取推送模式
	Mode() string
	// 执行推送
	PushMessage(param ...string) error
}

// 构造推送器
func BuildPushActuator(ctx *gin.Context, notifyVo vo.NotifyVo, app db.NotifyRegister, config db.NotifyConfig) PushActuator {
	switch app.PlatformTypeId {
	case constant.IOS:
		return PushApns{ctx, notifyVo, config}
	case constant.Android:
		return PushJPush{ctx, notifyVo, config}
	}
	return nil
}
