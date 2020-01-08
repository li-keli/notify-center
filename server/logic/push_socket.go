package logic

import (
	"encoding/json"
	"fmt"
	"notify-center/pkg/dto"
	"notify-center/pkg/redis"
	"notify-center/server/api/v1/vo"
	"strconv"
)

// WebSocket推送（独立推送）
type PushWSocket struct {
	NotifyVo vo.NotifyVo
}

func (p PushWSocket) PushMessage() error {
	m := make(map[string]string, len(p.NotifyVo.Data))
	for s, i := range p.NotifyVo.Data {
		// todo 转换为int/float的时候还存在些许问题
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
	marshal, e := json.Marshal(m)
	fmt.Println(string(marshal))
	// 发送WS广播
	redis.Publish(&dto.RedisStreamMessage{
		UniqueId: p.NotifyVo.JsjUniqueId,
		Body: dto.RedisStreamMessageBody{
			MAction: p.NotifyVo.Route,
			MBody:   string(marshal),
		},
	})

	return e
}
