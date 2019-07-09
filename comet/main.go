package main

import (
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	conns   = hashmap.New()
	upgrade = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {

	engine := gin.Default()

	engine.GET("/health/:jsjUniqueId", func(c *gin.Context) {
		c.String(http.StatusOK, string(conns.Size()))
		if v, isfound := conns.Get(c.Param("jsjUniqueId")); isfound {
			conn := v.(*websocket.Conn)
			_ = conn.WriteMessage(1, []byte("hello"))
		}
	})

	engine.GET("/v1/ws/:targetType/:jsjUniqueId", func(c *gin.Context) {
		var (
			jsjUniqueId = c.Param("jsjUniqueId")
		)

		ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}

		conns.Put(jsjUniqueId, ws)
		defer func() {
			conns.Remove(jsjUniqueId)
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
