package logic

import (
	"encoding/json"
	"notify-center/pkg/dto"
	"notify-center/pkg/redis"
	"notify-center/server/api/v1/vo"
)

// WebSocket推送（独立推送）
type PushWSocket struct {
	NotifyVo vo.NotifyVo
}

func (p PushWSocket) PushMessage() error {
	marshal, e := json.Marshal(p.NotifyVo.Data)
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
