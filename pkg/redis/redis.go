package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"notify-center/pkg/dto"
	"time"
)

var client *redis.Client

func NewRedisConn() {
	client = redis.NewClient(&redis.Options{
		Addr:     "172.16.7.20:6379",
		Password: "",
		DB:       3,
	})
	pong, err := client.Ping().Result()
	logrus.Info(pong, err)
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
		_ = msgObj.UnMarshal([]byte(msg.Payload))
		handle(&msgObj)
	}
}

// 读取缓存
func GetHash(k, hk string) (r string, e error) {
	r, e = client.HGet(k, hk).Result()
	logrus.Info(r)
	if e == redis.Nil {
		logrus.Error("key不存在; ", k)
		return "", errors.New("key不存在")
	}
	if e != nil {
		logrus.Error("读取缓存异常; ", e)
		return "", errors.New("读取缓存异常")
	}

	return
}

// 设置缓存
func SetHash(k, hk string, v []byte, t float64) {
	s, e := client.HSet(k, hk, v).Result()
	client.Expire(k, time.Duration(t)*time.Second)
	if e != nil {
		logrus.Error("写入缓存异常", s, e)
	}
}
