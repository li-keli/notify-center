package vo

import (
	"encoding/json"
	"notify-center/pkg/constant"
)

type NotifyVo struct {
	JsjUniqueId int                    `json:"JsjUniqueId"`
	TargetType  constant.TargetType    `json:"TargetType"`
	Title       string                 `json:"Title"`
	Message     string                 `json:"Message"`
	Route       string                 `json:"Route"`
	Data        map[string]interface{} `json:"Data"`

	// 以下字段废弃
	// ReadSendTime string                 `json:"ReadSendTime"`
}

func (n *NotifyVo) DataToBytes() []byte {
	marshal, _ := json.Marshal(n.Data)
	return marshal
}
