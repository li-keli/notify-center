// 钉钉推送

package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"notify-center/pkg/tracklog"
	"notify-center/server/api/v1/vo"
)

var (
	dingDingAccessTokenUrl = "https://openservice.jsjinfo.cn/tokencenter/v1/mini/dingding"
)

type PushDingDing struct {
	ctx      *gin.Context
	notifyVo vo.NotifyVo
}

type DingDingPushIn struct {
	AgentId    string          `json:"agent_id"`
	UserIdList string          `json:"userid_list"`
	Msg        DingDingPushMsg `json:"msg"`
}
type DingDingPushMsg struct {
	Markdown DingDingPushMarkdown `json:"markdown"`
}
type DingDingPushMarkdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
type DingDingPushOut struct {
	ErrCode   int    `json:"errcode"`
	TaskId    int    `json:"task_id"`
	RequestId string `json:"request_id"`
}

type DingDingGetByMobileOut struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	UserId  string `json:"userid"`
}

func (p PushDingDing) Mode() string {
	return "钉钉推送"
}

func (p PushDingDing) PushMessage(param ...string) error {
	var (
		in = DingDingPushIn{
			AgentId:    "272776007",
			UserIdList: p.dingDingUserId(param[0]),
			Msg: DingDingPushMsg{Markdown: DingDingPushMarkdown{
				Title: p.notifyVo.Title,
				Text:  p.notifyVo.Message,
			}},
		}
		trackLog = tracklog.Logger(p.ctx)
		out      DingDingPushOut
	)
	marshal, _ := json.Marshal(in)
	client := http.Client{}
	request, err := http.NewRequest("POST", "https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2?access_token="+p.dingDingAccessToken(), bytes.NewBuffer(marshal))
	if err != nil {
		trackLog.Error(err)
	}
	response, err := client.Do(request)
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(result, &out)

	return err
}

func (p PushDingDing) dingDingUserId(mobile string) (userId string) {
	var out DingDingGetByMobileOut
	response, err := http.Get(fmt.Sprintf("https://oapi.dingtalk.com/user/get_by_mobile?access_token=%s&mobile=%s", p.dingDingAccessToken(), mobile))
	defer response.Body.Close()
	if err != nil {
		logrus.Error(err)
		return ""
	}
	result, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(result, &out)
	if out.ErrCode != 0 {
		logrus.Error(out.ErrMsg)
	}

	return out.UserId
}

func (PushDingDing) dingDingAccessToken() string {
	client := http.Client{}
	request, _ := http.NewRequest("GET", dingDingAccessTokenUrl, nil)
	request.Header.Add("token", "VPYQnbnx2pp7rJxGI7v8HDjl2JJGyWyV")
	response, _ := client.Do(request)
	defer response.Body.Close()

	result, _ := ioutil.ReadAll(response.Body)
	logrus.Info(string(result))
	return string(result)
}
