package dto

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

// redis订阅发布通信模型
type RedisStreamMessage struct {
	UniqueId int
	Body     RedisStreamMessageBody
}
type RedisStreamMessageBody struct {
	MAction string
	MBody   string
}

func (rsm *RedisStreamMessage) Marshal() string {
	bytes, e := json.Marshal(rsm)
	if e != nil {
		logrus.Error("RedisStreamMessage序列化异常", e)
	}
	return string(bytes)
}

func (rsm *RedisStreamMessage) UnMarshal(b []byte) (e error) {
	e = json.Unmarshal(b, &rsm)
	return
}

func (rsm *RedisStreamMessageBody) Marshal() string {
	bytes, e := json.Marshal(rsm)
	if e != nil {
		logrus.Error("RedisStreamMessageBody序列化异常", e)
	}
	return string(bytes)
}

func (rsm *RedisStreamMessageBody) UnMarshal(b []byte) (e error) {
	e = json.Unmarshal(b, &rsm)
	return
}
