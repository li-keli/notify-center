package logic

import (
	"errors"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
	"github.com/sirupsen/logrus"
	"notify-center/pkg/db"
	"notify-center/server/api/v1/vo"
)

// APNS推送
type PushApns struct {
	notifyVo vo.NotifyVo
	config   db.NotifyConfig
}

func (p PushApns) Mode() string {
	return "APNS推送"
}

func (p PushApns) PushMessage(param ...string) error {
	var pushToken = param[0]
	config, err := p.config.IosConfig()
	if err != nil {
		logrus.Error("构造APNS推送配置错误", err)
		return err
	}

	cert, err := certificate.FromP12File(config.P12Path, config.Password)
	if err != nil {
		logrus.Error("Cert Error:", err)
		return err
	}

	pload := payload.NewPayload()
	pload.Badge(1)
	pload.AlertTitle(p.notifyVo.Title)
	pload.AlertBody(p.notifyVo.Message)
	pload.Custom("MAction", p.notifyVo.Route)
	pload.Custom("MBody", p.notifyVo.DataToBytes())

	notification := &apns2.Notification{}
	notification.DeviceToken = pushToken
	notification.Topic = config.BundleId
	notification.Payload = pload
	client := apns2.NewClient(cert).Production()
	res, err := client.Push(notification)
	if err != nil {
		logrus.Errorf("Apns推送错误，%s %s %s %s", res.Reason, pushToken, config.P12Path, config.BundleId)
		return err
	}
	if res.StatusCode != 200 {
		logrus.Errorf("Apns推送失败，错误码：%d %s %s %s %s", res.StatusCode, res.Reason, pushToken, config.P12Path, config.BundleId)
		return errors.New(res.Reason)
	}

	return err
}
