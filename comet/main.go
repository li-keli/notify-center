package main

import (
	"encoding/json"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"notify-center/pkg/constant"
	"notify-center/pkg/redis"
	"strconv"
)

var (
	connList = arraylist.New()
	upgrade  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type ConnStruct struct {
	Key   int
	Sid   string
	Tid   int
	TName string
	Conn  *websocket.Conn `json:"-"`
}

func main() {

	// 连接redis
	redis.NewRedisConn()

	// redis 订阅发布
	go redis.Subscribe(func(id int, body string) {
		connList.Each(func(index int, value interface{}) {
			targetConn := value.(*ConnStruct)
			if targetConn.Key == id {
				logrus.Infof("处理广播消息 %d; %s", targetConn.Key, body)
				if e := targetConn.Conn.WriteMessage(websocket.TextMessage, []byte(body)); e != nil {
					redis.DelHashField(strconv.Itoa(id), targetConn.Sid)
				}
			}
		})
	})

	engine := gin.Default()
	engine.GET("/v1/ws/:targetType/:uniqueId", func(c *gin.Context) {
		var (
			targetType, _ = strconv.Atoi(c.Param("targetType"))
			uniqueId, _   = strconv.Atoi(c.Param("uniqueId"))
			sId           = uuid.NewV4().String()
			ws, err       = upgrade.Upgrade(c.Writer, c.Request, nil)
			connModule    = &ConnStruct{
				Key:   uniqueId,
				Sid:   sId,
				Tid:   targetType,
				TName: constant.TargetTypeValueOf(constant.TargetType(targetType)),
				Conn:  ws,
			}
		)

		if err != nil {
			return
		}

		connList.Add(connModule)
		connModuleBytes, _ := json.Marshal(connModule)
		redis.SetHash(strconv.Itoa(uniqueId), sId, connModuleBytes, 10)
		logrus.Infof("WS连接数：%d", connList.Size())
		defer func() {
			ws.Close()
			connList.Remove(connList.IndexOf(connModule))
			redis.DelHashField(strconv.Itoa(uniqueId), sId)
			logrus.Infof("WS连接数：%d", connList.Size())
		}()

		for {
			//读取ws中的数据
			mt, message, err := ws.ReadMessage()
			if err != nil {
				break
			}

			if string(message) == "+" {
				message = []byte("-")
				redis.SetHash(strconv.Itoa(uniqueId), sId, connModuleBytes, 10)
			}
			//写入ws数据
			err = ws.WriteMessage(mt, message)
			if err != nil {
				break
			}
		}
	})

	_ = engine.Run()
}
