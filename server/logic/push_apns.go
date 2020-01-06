package logic

import (
	"fmt"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
	"github.com/sirupsen/logrus"
	"notify-center/pkg/db"
	"notify-center/server/api/v1/vo"
)

// 苹果APNS推送
type PushApns struct {
	notifyVo vo.NotifyVo
	config   db.NotifyConfig
}

func (p *PushApns) PushMessage(pushToken string) error {
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
		logrus.Error("Error:", err)
		return err
	}

	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
	return err
}
