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
func BuildPushActuator(notifyVo vo.NotifyVo, app db.NotifyRegister, config db.NotifyConfig) PushActuator {
	//switch app.PlatformTypeId {
	//case constant.IOS:
	//	return &PushApns{notifyVo, config}
	//case constant.Android:
	//	return &PushJPush{notifyVo, config}
	//case constant.DingDing:
	//	return nil
	//}

	return nil
}
