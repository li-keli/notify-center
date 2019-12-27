package logic

import (
	"fmt"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
	"log"
	"notify-center/pkg/db"
	"notify-center/server/api/v1/vo"
)

// 苹果APNS推送
type PushApns struct {
	notifyVo vo.NotifyVo
	config   db.DicConfigEntity
}

func (p *PushApns) PushMessage(pushToken string) error {
	cert, err := certificate.FromP12File(p.config.IosConfig.P12Path, p.config.IosConfig.Password)
	if err != nil {
		log.Fatal("Cert Error:", err)
	}

	pload := payload.NewPayload()
	pload.Badge(1)
	pload.AlertTitle(p.notifyVo.Title)
	pload.AlertBody(p.notifyVo.Message)
	pload.Custom("MAction", p.notifyVo.Route)
	pload.Custom("MBody", p.notifyVo.DataToBytes())

	notification := &apns2.Notification{}
	notification.DeviceToken = pushToken
	notification.Topic = p.config.IosConfig.BundleId
	notification.Payload = pload
	client := apns2.NewClient(cert).Production()
	res, err := client.Push(notification)
	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
	return err
}
