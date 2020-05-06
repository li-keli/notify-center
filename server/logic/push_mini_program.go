// 微信小程序订阅消息下发
// 相关文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html#method-http
package logic

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"notify-center/pkg/db"
	"notify-center/server/api/v1/vo"
)

var (
	// 微信服务器端点
	weChatHost           = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token="
	weChatAccessTokenUrl = "https://*******/ktgj"
)

type PushMiniProgram struct {
	ctx      *gin.Context
	notifyVo vo.NotifyVo
	config   db.NotifyConfig
}

func (p PushMiniProgram) Mode() string {
	return "微信小程序订阅消息"
}

func (p PushMiniProgram) PushMessage(param ...string) error {
	if len(param) != 4 {
		return errors.New("微信小程序订阅消息参数错误")
	}
	var (
		temp = struct {
			ToUser     string      `json:"touser"`
			TemplateId string      `json:"template_id"`
			Page       string      `json:"page"`
			Data       interface{} `json:"data"`
		}{
			param[0],
			param[1],
			param[2],
			param[3],
		}
	)
	client := http.Client{}
	marshal, _ := json.Marshal(temp)
	request, _ := http.NewRequest("POST", weChatHost+p.weChatAccessToken(), bytes.NewBuffer(marshal))
	response, err := client.Do(request)
	if err != nil {
		logrus.Error("微信小程序订阅消息推送异常", err)
	}
	defer response.Body.Close()

	result, _ := ioutil.ReadAll(response.Body)
	logrus.Info(string(result))
	return nil
}

// 微信授权token
func (PushMiniProgram) weChatAccessToken() string {
	client := http.Client{}
	request, _ := http.NewRequest("GET", weChatAccessTokenUrl, nil)
	request.Header.Add("token", "VPY***")
	response, _ := client.Do(request)
	defer response.Body.Close()

	result, _ := ioutil.ReadAll(response.Body)
	logrus.Info(string(result))
	return string(result)
}
