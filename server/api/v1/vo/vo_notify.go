package vo

import (
	"encoding/json"
	"fmt"
	"notify-center/pkg/constant"
	"strconv"
)

type NotifyVo struct {
	JsjUniqueId int                    `json:"JsjUniqueId"`
	TargetType  constant.TargetType    `json:"TargetType"`
	Title       string                 `json:"Title"`
	Message     string                 `json:"Message"`
	GroupName   string                 `json:"GroupName"`
	Route       string                 `json:"Route"`
	Data        map[string]interface{} `json:"Data"`

	// 以下字段废弃
	// ReadSendTime string                 `json:"ReadSendTime"`
}

// 将Data转换为Byte并将interface转换为对应类型
func (n *NotifyVo) DataToBytes() []byte {
	m := make(map[string]string, len(n.Data))
	for s, i := range n.Data {
		switch i.(type) {
		case int, float64:
			if v, b := i.(int); b {
				// 整形
				m[s] = strconv.Itoa(v)
			} else {
				// 浮点数
				m[s] = fmt.Sprintf("%.0f", i)
			}
		default:
			m[s] = fmt.Sprintf("%v", i)
		}
	}

	marshal, _ := json.Marshal(m)
	return marshal
}

func (n *NotifyVo) DataToStr() string {
	m := make(map[string]string, len(n.Data))
	for s, i := range n.Data {
		switch i.(type) {
		case int, float64:
			if v, b := i.(int); b {
				// 整形
				m[s] = strconv.Itoa(v)
			} else {
				// 浮点数
				m[s] = fmt.Sprintf("%.0f", i)
			}
		default:
			m[s] = fmt.Sprintf("%v", i)
		}
	}

	marshal, _ := json.Marshal(m)
	return string(marshal)
}
