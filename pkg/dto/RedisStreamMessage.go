package dto

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

// redis订阅发布通信模型
type RedisStreamMessage struct {
	UniqueId int                    `json:"UniqueId"`
	Body     RedisStreamMessageBody `json:"Body"`
}
type RedisStreamMessageBody struct {
	MAction string `json:"MAction"`
	MBody   string `json:"MBody"`
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

func (rsm *RedisStreamMessageBody) Marshal() []byte {
	bytes, e := json.Marshal(rsm)
	if e != nil {
		logrus.Error("RedisStreamMessageBody序列化异常", e)
	}
	return bytes
}

func (rsm *RedisStreamMessageBody) UnMarshal(b []byte) (e error) {
	e = json.Unmarshal(b, &rsm)
	return
}
