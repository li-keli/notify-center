package vo

import "notify-center/pkg/constant"

type RegisterTerminalInputVo struct {
	JsjUniqueId  int                   `json:"JsjUniqueId" binding:"min=1"`
	PushToken    string                `json:"PushToken" binding:"required"`
	PlatformType constant.PlatformType `json:"PlatformType"`
	TargetType   constant.TargetType   `json:"TargetType"`
}

type UnRegisterTerminalInputVo struct {
	JsjUniqueId  int                   `json:"JsjUniqueId" binding:"min=1"`
	PushToken    string                `json:"PushToken" binding:"required"`
	PlatformType constant.PlatformType `json:"PlatformType"`
	TargetType   constant.TargetType   `json:"TargetType"`
}
