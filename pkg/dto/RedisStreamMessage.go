package dto

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

// redis订阅发布通信模型
type RedisStreamMessage struct {
	JsjUniqueId string
	Body        string
}

func (rsm *RedisStreamMessage) Marshal() (string, error) {
	bytes, e := json.Marshal(rsm)
	if e != nil {
		logrus.Error("RedisStreamMessage序列化异常", e)
	}
	return string(bytes), e
}

func (rsm *RedisStreamMessage) UnMarshal(b []byte) (e error) {
	e = json.Unmarshal(b, &rsm)
	return
}
