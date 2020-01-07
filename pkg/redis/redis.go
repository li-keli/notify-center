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
	if _, e := client.Ping().Result(); e != nil {
		logrus.Info(e)
	}
	logrus.Info("Redis连接成功...")
}

// 消息发布
func Publish(msg *dto.RedisStreamMessage) {
	s := msg.Marshal()
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
		logrus.Info("收到消息 ", msg.Payload)
		_ = msgObj.UnMarshal([]byte(msg.Payload))
		handle(&msgObj)
	}
}

// 读取缓存(Hash)
func GetHashAll(k string) (r []string, e error) {
	result, e := client.HGetAll(k).Result()
	if e == redis.Nil {
		logrus.Error("key不存在; ", k)
		return []string{}, errors.New("key不存在")
	}
	if e != nil {
		logrus.Error("读取缓存异常; ", e)
		return []string{}, errors.New("读取缓存异常")
	}

	for s, _ := range result {
		r = append(r, s)
	}
	return
}

// 删除Hash中的Field
func DelHashField(k, f string) {
	go client.HDel(k, f)
}

// 设置缓存(Hash)
func SetHash(k, hk string, v []byte, t float64) {
	s, e := client.HSet(k, hk, v).Result()
	client.Expire(k, time.Duration(t)*time.Second)
	if e != nil {
		logrus.Error("写入缓存异常", s, e)
	}
}