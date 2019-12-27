package vo

import (
	"encoding/json"
	"notify-center/pkg/constant"
)

type NotifyVo struct {
	JsjUniqueId  int                    `json:"JsjUniqueId"`
	TargetType   constant.TargetType    `json:"TargetType"`
	Title        string                 `json:"Title"`
	Message      string                 `json:"Message"`
	Route        string                 `json:"Route"`
	ReadSendTime string                 `json:"ReadSendTime"`
	Data         map[string]interface{} `json:"Data"`
}

func (n *NotifyVo) DataToBytes() []byte {
	marshal, _ := json.Marshal(n.Data)
	return marshal
}
