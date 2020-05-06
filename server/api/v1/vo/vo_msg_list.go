package vo

import "notify-center/pkg/db"

type MsgListInputVo struct {
	JsjUniqueId int    `json:"JsjUniqueId"`
	PushToken   string `json:"PushToken"`
	Offset      int    `json:"Offset" binding:"min=1"`
	Limit       int    `json:"Limit" binding:"max=20"`
}

type MsgListOutputVo struct {
	BaseOutput
	Msg []db.NotifyMsg `json:"Msg"`
}
