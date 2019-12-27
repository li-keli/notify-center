package logic

import (
	"notify-center/pkg/db"
	"notify-center/server/api/v1/vo"
)

// 推送逻辑
type PushActuator interface {
	PushMessage(pushToken string) error
}

// 构造推送器
func BuildPushActuator(notifyVo vo.NotifyVo, app db.AppEntity, config db.DicConfigEntity) PushActuator {
	return &PushWSocket{notifyVo}
}
