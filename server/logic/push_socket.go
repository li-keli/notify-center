package logic

import (
	"notify-center/pkg/dto"
	"notify-center/pkg/redis"
	"strconv"
)

// WebSocket推送
type PushWSocket struct {
}

func (p *PushWSocket) PushMessage(pushKey string) {
	jsjUniqueId, _ := strconv.Atoi(pushKey)
	// 发送WS广播
	redis.Publish(&dto.RedisStreamMessage{
		UniqueId: jsjUniqueId,
		Body: dto.RedisStreamMessageBody{
			MAction: "demo",
			MBody:   "Ok i got the message",
		},
	})
}
