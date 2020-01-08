package db

import (
	"github.com/sirupsen/logrus"
	"notify-center/pkg/constant"
	"time"
)

type NotifyMsg struct {
	Id               int                   `json:"Id" gorm:"column:id;AUTO_INCREMENT"`
	PushToken        string                `json:"PushToken" gorm:"column:push_token"`
	DataContent      string                `json:"DataContent" gorm:"column:data_content"`
	PlatformTypeId   constant.PlatformType `json:"PlatformTypeId" gorm:"column:platform_type_id"`
	PlatformTypeName string                `json:"PlatformTypeName" gorm:"column:platform_type_name"`
	TargetTypeId     constant.TargetType   `json:"TargetTypeId" gorm:"column:target_type_id"`
	TargetTypeName   string                `json:"TargetTypeName" gorm:"column:target_type_name"`
	CreateTime       time.Time             `json:"CreateTime" gorm:"column:create_time"`
}

func (NotifyMsg) Insert(m NotifyMsg) {
	if err := conn.Create(&m).Error; err != nil {
		logrus.Error("历史消息存储错误，但忽略了这个错误")
	}
}

func (NotifyMsg) FindAll(token, start, end string, limit int) (n []NotifyMsg, err error) {
	err = conn.Where(`push_token = ? and create_time between ? and ?`, token, start, end).Limit(limit).Order(`create_time desc`).Find(&n).Error
	location, _ := time.LoadLocation("")
	for i, msg := range n {
		n[i].CreateTime = msg.CreateTime.In(location)
	}
	return
}
