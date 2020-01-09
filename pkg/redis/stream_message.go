package redis

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

// redis订阅发布通信模型
type StreamMessage struct {
	UniqueId int               `json:"UniqueId"`
	Body     StreamMessageBody `json:"Body"`
}
type StreamMessageBody struct {
	MAction string `json:"MAction"`
	MBody   string `json:"MBody"`
}

func (rsm *StreamMessage) Marshal() string {
	bytes, e := json.Marshal(rsm)
	if e != nil {
		logrus.Error("StreamMessage序列化异常", e)
	}
	return string(bytes)
}

func (StreamMessage) UnMarshal(b []byte) (r StreamMessage, e error) {
	e = json.Unmarshal(b, &r)
	return
}

func (rsm *StreamMessageBody) Marshal() []byte {
	bytes, e := json.Marshal(rsm)
	if e != nil {
		logrus.Error("StreamMessageBody序列化异常", e)
	}
	return bytes
}
