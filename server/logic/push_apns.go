package logic

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
	"notify-center/pkg/constant"
	"notify-center/pkg/db"
	"notify-center/pkg/track_log"
	"notify-center/server/api/v1/vo"
)

// APNS推送
type PushApns struct {
	ctx      *gin.Context
	notifyVo vo.NotifyVo
	config   db.NotifyConfig
}

func (p PushApns) Mode() string {
	return "IOS平台"
}

func (p PushApns) PushMessage(param ...string) error {
	var (
		pushToken = param[0]
		trackLog  = track_log.Logger(p.ctx)
	)
	config, err := p.config.IosConfig()
	if err != nil {
		trackLog.Error("序列化IosConfig错误，原文：", err.Error(), p.config.ConfigData)
		return err
	}

	cert, err := certificate.FromP12File(config.P12Path, config.Password)
	if err != nil {
		trackLog.Error("Cert Error:", err)
		return err
	}

	pload := payload.NewPayload()
	pload.Badge(1)
	pload.AlertTitle(p.notifyVo.Title)
	pload.AlertBody(p.notifyVo.Message)
	pload.Custom("MAction", p.notifyVo.Route)
	pload.Custom("MBody", p.notifyVo.DataToStr())

	notification := &apns2.Notification{}
	notification.DeviceToken = pushToken
	notification.Topic = config.BundleId
	notification.Payload = pload

	client := apns2.NewClient(cert)
	// 环境区分
	if constant.ProductionMode {
		client.Production()
	} else {
		client.Development()
	}

	res, err := client.Push(notification)
	if err != nil {
		trackLog.Errorf("Apns推送错误，%s %s %s %s %s", err.Error(), client.Host, pushToken, config.P12Path, config.BundleId)
		return err
	}
	if res.StatusCode != 200 {
		trackLog.Errorf("Apns推送失败，错误码：%d %s %s %s %s %s", res.StatusCode, client.Host, res.Reason, pushToken, config.P12Path, config.BundleId)
		return errors.New(res.Reason)
	}

	trackLog.Infof("Apns推送成功 %s %s %s %s", client.Host, pushToken, config.P12Path, config.BundleId)
	return err
}
