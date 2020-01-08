package logic

import (
	"github.com/sirupsen/logrus"
	jpushclient "github.com/ylywyn/jpush-api-go-client"
	"notify-center/pkg/db"
	"notify-center/server/api/v1/vo"
)

// Android JPush推送
type PushJPush struct {
	notifyVo vo.NotifyVo
	config   db.NotifyConfig
}

func (p PushJPush) Mode() string {
	return "Android平台"
}

func (p PushJPush) PushMessage(param ...string) error {
	var pushToken = param[0]
	config, err := p.config.AndroidConfig()
	if err != nil {
		logrus.Error("构造极光推送配置错误", err)
		return err
	}

	AndroidClient := jpushclient.NewPushClient(config.Secret, config.AppKey)
	var pf jpushclient.Platform
	_ = pf.Add(jpushclient.ANDROID)

	var ad jpushclient.Audience
	ad.SetID([]string{pushToken})

	var notice jpushclient.Notice
	notice.SetAndroidNotice(&jpushclient.AndroidNotice{
		Alert: p.notifyVo.Title,
		Extras: map[string]interface{}{
			"MAction": p.notifyVo.Route,
			"MBody":   string(p.notifyVo.DataToBytes()),
		}},
	)

	var msg jpushclient.Message
	msg.Title = p.notifyVo.Title
	msg.Content = p.notifyVo.Message
	//msg.Extras = p.notifyVo.Data

	payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&ad)
	payload.SetMessage(&msg)
	payload.SetNotice(&notice)
	bytes, err := payload.ToBytes()
	logrus.Info(string(bytes))

	str, err := AndroidClient.Send(bytes)
	logrus.Info(str)

	return err
}
