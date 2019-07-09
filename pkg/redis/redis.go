package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"notify-center/pkg/dto"
)

var client *redis.Client

func NewRedisConn() {
	client = redis.NewClient(&redis.Options{
		Addr:     "172.16.7.20",
		Password: "",
		DB:       3,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

// 消息发布
func Publish(msg *dto.RedisStreamMessage) {
	s, _ := msg.Marshal()
	if err := client.Publish("homework", s).Err(); err != nil {
		logrus.Error("redis消息发布异常", err)
	}
}

// 消息订阅
func Subscribe(handle func(msg *dto.RedisStreamMessage)) {
	pubSub := client.Subscribe("homework")
	if _, e := pubSub.Receive(); e != nil {
		logrus.Fatal("redis服务订阅失败")
	}
	channel := pubSub.Channel()
	for msg := range channel {
		var msgObj = dto.RedisStreamMessage{}
		logrus.Info(channel, "收到消息", msg.Payload)
		_ := msgObj.UnMarshal([]byte(msg.Payload))
		handle(&msgObj)
	}
}
