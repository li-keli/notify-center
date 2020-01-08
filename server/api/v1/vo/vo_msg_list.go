package vo

import "notify-center/pkg/db"

type MsgListInputVo struct {
	PushToken string `json:"PushToken"`
	StartTime string `json:"StartTime"`
	EntTime   string `json:"EntTime"`
	Limit     int    `json:"Limit"`
	Offset    int    `json:"Offset"`
}

type MsgListOutputVo struct {
	BaseOutput
	Msg []db.NotifyMsg `json:"Msg"`
}
