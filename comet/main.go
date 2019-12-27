package main

import (
	"encoding/json"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
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
	Sid  string
	Conn *websocket.Conn `json:"-"`
}

func main() {

	// 连接redis
	redis.NewRedisConn()

	// redis 订阅发布
	go redis.Subscribe(func(msg *dto.RedisStreamMessage) {
		connList.Each(func(index int, value interface{}) {
			targetConn := value.(*ConnStruct)
			if targetConn.Key == msg.UniqueId {
				logrus.Infof("处理广播消息 %s", targetConn.Key)
				if e := targetConn.Conn.WriteMessage(1, []byte(msg.Body.Marshal())); e != nil {
					redis.DelHashField(strconv.Itoa(msg.UniqueId), targetConn.Sid)
				}
			}
		})
	})

	engine := gin.Default()
	engine.GET("/health/:uniqueId", func(c *gin.Context) {
		logrus.Info("会话列表容量", connList.Size())
		var uniqueId, _ = strconv.Atoi(c.Param("uniqueId"))
		_, obj := connList.Find(func(index int, value interface{}) bool {
			stru, b := value.(*ConnStruct)
			if !b {
				return false
			}
			if stru.Key == uniqueId {
				return true
			}
			return false
		})
		connStruct := obj.(*ConnStruct)
		_ = connStruct.Conn.WriteMessage(1, []byte("hello"))

		c.String(http.StatusOK, string(connList.Size()))
	})
	engine.GET("/v1/ws/:targetType/:uniqueId", func(c *gin.Context) {
		var (
			//targetType, _  = strconv.Atoi(c.Param("targetType"))
			uniqueId, _ = strconv.Atoi(c.Param("uniqueId"))
			sId         = uuid.NewV4().String()
		)

		ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}

		connModule := &ConnStruct{
			Key:  uniqueId,
			Sid:  sId,
			Conn: ws,
		}
		connList.Add(connModule)
		connModuleBytes, _ := json.Marshal(connModule)
		redis.SetHash(strconv.Itoa(uniqueId), sId, connModuleBytes, 60)
		logrus.Infof("WS连接数：%d", connList.Size())
		defer func() {
			ws.Close()
			connList.Remove(connList.IndexOf(connModule))
			redis.DelHashField(strconv.Itoa(uniqueId), sId)
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

	_ = engine.Run("localhost:8081")
}
