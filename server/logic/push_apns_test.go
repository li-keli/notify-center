package logic

import (
	"fmt"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"log"
	"testing"
)

func TestPushApns_PushMessage(t *testing.T) {
	cert, err := certificate.FromP12File("../certificates/kt_no1_prod.p12", "1")
	if err != nil {
		log.Fatal("Cert Error:", err)
	}

	notification := &apns2.Notification{}
	notification.DeviceToken = "abc3e4c585483a65e1fe3a7b1403524de97b804139029d8d03bbc52e2c3c10e8"
	notification.Topic = "cn.com.jsj.share"
	notification.Payload = []byte(`{"aps":{"alert":"Hello!"}}`)

	client := apns2.NewClient(cert).Production()
	res, err := client.Push(notification)

	if err != nil {
		log.Fatal("Error:", err)
	}
	if res.StatusCode != 200 {
		log.Fatal("Apns推送失败，", res.Reason)
	}
	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}
