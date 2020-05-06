package logic

import (
	"github.com/gin-gonic/gin"
	"notify-center/pkg/redis"
	"notify-center/server/api/v1/vo"
)

// WebSocket推送（独立推送）
type PushWSocket struct {
	ctx      *gin.Context
	NotifyVo vo.NotifyVo
}

func (p PushWSocket) PushMessage() (err error) {
	// 发送WS广播
	redis.Publish(&redis.StreamMessage{
		UniqueId: p.NotifyVo.JsjUniqueId,
		Body: redis.StreamMessageBody{
			MAction: p.NotifyVo.Route,
			MBody:   p.NotifyVo.DataToStr(),
		},
	})
	return
}
