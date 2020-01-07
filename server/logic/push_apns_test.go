package logic

import (
	"fmt"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"log"
	"testing"
)

func TestPushApns_PushMessage(t *testing.T) {
	cert, err := certificate.FromP12File("../certificates/kt_no1_debug.p12", "1")
	if err != nil {
		log.Fatal("Cert Error:", err)
	}

	notification := &apns2.Notification{}
	notification.DeviceToken = "5c660ad94edf7a42668533d30b41b248609279bb84043e1eafd3657dd6c4a077"
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
