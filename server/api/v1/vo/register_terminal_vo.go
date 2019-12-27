package vo

import "notify-center/pkg/constant"

type RegisterTerminalVo struct {
	JsjUniqueId  int                   `json:"JsjUniqueId"`
	PushToken    string                `json:"PushToken"`
	PlatformType constant.PlatformType `json:"PlatformType"`
	TargetType   constant.TargetType   `json:"TargetType"`
}

type UnRegisterTerminalVo struct {
	JsjUniqueId  int                   `json:"JsjUniqueId"`
	PushToken    string                `json:"PushToken"`
	PlatformType constant.PlatformType `json:"PlatformType"`
	TargetType   constant.TargetType   `json:"TargetType"`
}
