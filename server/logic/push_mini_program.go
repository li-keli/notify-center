// 微信小程序模板
package logic

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"notify-center/pkg/db"
	"notify-center/server/api/v1/vo"
)

var (
	// 微信服务器端点
	weChatHost           = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token="
	WeChatAccessTokenUrl = "https://openservice.jsjinfo.cn/tokencenter/v1/mini/ktgj"
)

type PushMiniProgram struct {
	notifyVo vo.NotifyVo
	config   db.NotifyConfig
}

func (p PushMiniProgram) Mode() string {
	return "微信小程序模板推送"
}

func (p PushMiniProgram) PushMessage(param ...string) error {
	panic("开发中...")
}

// 微信授权token
func (PushMiniProgram) weChatAccessToken() {
	client := http.Client{}
	request, _ := http.NewRequest("GET", WeChatAccessTokenUrl, nil)
	request.Header.Add("token", "VPYQnbnx2pp7rJxGI7v8HDjl2JJGyWyV")
	response, _ := client.Do(request)
	defer response.Body.Close()

	result, _ := ioutil.ReadAll(response.Body)
	logrus.Info(string(result))
}
