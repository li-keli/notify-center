package logic

import (
	"fmt"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"log"
	"notify-center/pkg/db"
	"notify-center/server/api/v1/vo"
	"testing"
	"time"
)

func TestPushApns(t *testing.T) {
	cert, err := certificate.FromP12File("../certificates/certificates_debug.p12", "1")
	if err != nil {
		log.Fatal("Cert Error:", err)
	}

	notification := &apns2.Notification{}
	notification.DeviceToken = "383be6e2c0b7bacca5ee0ef5417e40aa321e2985359c8244b4b55b51affd908c"
	notification.Topic = "cn.com.jsj.share"
	notification.Payload = []byte(`{"aps":{"alert":"Hello!"}}`)

	client := apns2.NewClient(cert).Development()
	res, err := client.Push(notification)

	if err != nil {
		log.Fatal("Error:", err)
	}
	if res.StatusCode != 200 {
		log.Fatal("Apns推送失败，", res.Reason)
	}
	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}

func TestPushApns_PushMessage(t *testing.T) {
	err := PushApns{
		notifyVo: vo.NotifyVo{
			JsjUniqueId: 2059797,
			TargetType:  600,
			Title:       "分享消息",
			Message:     "扫描成功",
			Route:       "ktNo1.share.qrcode",
			Data: map[string]interface{}{
				"jsjUniqueId": 2059797,
				"orderNumber": 100002,
				"opType":      1,
				"title":       "分享消息",
				"message":     "扫描成功",
			},
		},
		config: db.NotifyConfig{
			PlatformTypeId:   10,
			PlatformTypeName: "IOS",
			TargetTypeId:     600,
			TargetTypeName:   "KtNo1",
			ConfigData:       `{"bundleId":"cn.com.jsj.share","p12DevPath":"../certificates/certificates_debug.p12","p12ProdPath":"certificates/certificates_debug.p12","password":"123"}`,
			CreateTime:       time.Now(),
		},
	}.PushMessage("383be6e2c0b7bacca5ee0ef5417e40aa321e2985359c8244b4b55b51affd908c")
	if err != nil {
		t.Error(err)
	}
}
