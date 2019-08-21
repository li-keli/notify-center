package main

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"notify-center/pkg/dto"
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
	Key  int
	Conn *websocket.Conn
}

func main() {

	// redis 订阅发布
	redis.NewRedisConn()
	go redis.Subscribe(func(msg *dto.RedisStreamMessage) {
		if v, b := connList.Get(msg.JsjUniqueId); b {
			conn := v.(*ConnStruct).Conn
			_ = conn.WriteMessage(1, []byte(msg.Body))
		}
	})

	engine := gin.Default()
	engine.GET("/health/:jsjUniqueId", func(c *gin.Context) {
		logrus.Info("会话列表容量", connList.Size())
		var jsjUniqueId, _ = strconv.Atoi(c.Param("jsjUniqueId"))
		_, obj := connList.Find(func(index int, value interface{}) bool {
			stru, b := value.(*ConnStruct)
			if !b {
				return false
			}
			if stru.Key == jsjUniqueId {
				return true
			}
			return false
		})
		connStruct := obj.(*ConnStruct)
		_ = connStruct.Conn.WriteMessage(1, []byte("hello"))

		c.String(http.StatusOK, string(connList.Size()))
	})
	engine.GET("/v1/ws/:targetType/:jsjUniqueId", func(c *gin.Context) {
		var (
			//targetType  = c.Param("targetType")
			jsjUniqueId, _ = strconv.Atoi(c.Param("jsjUniqueId"))
		)

		ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}

		connModule := &ConnStruct{
			Key:  jsjUniqueId,
			Conn: ws,
		}
		connList.Add(connModule)
		defer func() {
			connList.Remove(connList.IndexOf(connModule))
			ws.Close()
		}()

		for {
			//读取ws中的数据
			mt, message, err := ws.ReadMessage()
			if err != nil {
				break
			}

			if string(message) == "+" {
				message = []byte("-")
			}
			//写入ws数据
			err = ws.WriteMessage(mt, message)
			if err != nil {
				break
			}
		}
	})

	_ = engine.Run(":8081")
}
