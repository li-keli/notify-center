package logic

import (
	"encoding/json"
	"notify-center/pkg/dto"
	"notify-center/pkg/redis"
	"notify-center/server/api/v1/vo"
)

// WebSocket推送
type PushWSocket struct {
	notifyVo vo.NotifyVo
}

func (p *PushWSocket) PushMessage(pushToken string) error {
	marshal, e := json.Marshal(p.notifyVo.Data)
	// 发送WS广播
	redis.Publish(&dto.RedisStreamMessage{
		UniqueId: p.notifyVo.JsjUniqueId,
		Body: dto.RedisStreamMessageBody{
			MAction: p.notifyVo.Route,
			MBody:   string(marshal),
		},
	})

	return e
}
