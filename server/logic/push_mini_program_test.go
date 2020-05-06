package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPushMiniProgram_(t *testing.T) {
	var (
		accessToken = "29_CNSqVFzE7cksCXVRCTMVOfVYuD5zRch69SJQO6tjiWXK23UhbuT_UyeurmF1xRL-pCSZSvI5LbQkXOyTVtI_VrSKpPDZxxcJFLRhkVPxHi3TYI1I4XczYupJZOZmNDOAGIMtPWfoOYLHlo8GXZCiABABDG"
		temp        struct {
			ToUser     string      `json:"touser"`
			TemplateId string      `json:"template_id"`
			Page       string      `json:"page"`
			Data       interface{} `json:"data"`
		}
	)

	m := make(map[string]map[string]string)
	m["amount1"] = map[string]string{"value": "100元"}
	m["date2"] = map[string]string{"value": "2020年01月09日"}

	temp.ToUser = "oG5Z_4iYeizOQSJ8za5qM0FMseQA"
	temp.TemplateId = "nPyY9kOiXExATn7cxB7DVTVgETl1lE1QNLSayNez8vA"
	temp.Page = ""
	temp.Data = m

	client := http.Client{}
	marshal, _ := json.Marshal(temp)
	fmt.Println(string(marshal))

	request, _ := http.NewRequest("POST", weChatHost+accessToken, bytes.NewBuffer(marshal))
	response, err := client.Do(request)
	if err != nil {
		logrus.Error("微信小程序订阅消息推送异常", err)
	}
	defer response.Body.Close()

	result, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(result))
}

func TestPushMiniProgram_weChatAccessToken(t *testing.T) {
	t.Log(PushMiniProgram{}.weChatAccessToken())
}
