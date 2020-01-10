package vo

import "notify-center/pkg/constant"

type RegisterTerminalInputVo struct {
	JsjUniqueId  int                   `json:"JsjUniqueId" binding:"requer"`
	PushToken    string                `json:"PushToken" binding:"min=1"`
	PlatformType constant.PlatformType `json:"PlatformType"`
	TargetType   constant.TargetType   `json:"TargetType"`
}

type UnRegisterTerminalInputVo struct {
	JsjUniqueId  int                   `json:"JsjUniqueId"`
	PushToken    string                `json:"PushToken"`
	PlatformType constant.PlatformType `json:"PlatformType"`
	TargetType   constant.TargetType   `json:"TargetType"`
}
