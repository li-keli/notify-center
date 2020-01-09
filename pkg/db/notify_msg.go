package db

import (
	"github.com/sirupsen/logrus"
	"notify-center/pkg/constant"
	"time"
)

type NotifyMsg struct {
	Id               int                   `json:"Id" gorm:"column:id;AUTO_INCREMENT"`
	JsjUniqueId      int                   `json:"JsjUniqueId" gorm:"column:jsj_unique_id"`
	PushToken        string                `json:"PushToken" gorm:"column:push_token"`
	DataContent      string                `json:"DataContent" gorm:"column:data_content"`
	PlatformTypeId   constant.PlatformType `json:"PlatformTypeId" gorm:"column:platform_type_id"`
	PlatformTypeName string                `json:"PlatformTypeName" gorm:"column:platform_type_name"`
	TargetTypeId     constant.TargetType   `json:"TargetTypeId" gorm:"column:target_type_id"`
	TargetTypeName   string                `json:"TargetTypeName" gorm:"column:target_type_name"`
	GroupName        string                `json:"GroupName" gorm:"column:group_name"`
	CreateTime       time.Time             `json:"CreateTime" gorm:"column:create_time"`
}

func (NotifyMsg) Insert(m NotifyMsg) {
	if err := conn.Create(&m).Error; err != nil {
		logrus.Error("历史消息存储错误，但忽略了这个错误")
	}
}

func (NotifyMsg) FindAll(jsjUniqueId int, token string, offset, limit int) (n []NotifyMsg, err error) {
	err = conn.Where(`jsj_unique_id = ? or push_token = ?`, jsjUniqueId, token).Offset(offset - 1).Limit(limit).Order(`create_time desc`).Find(&n).Error
	location, _ := time.LoadLocation("")
	for i, msg := range n {
		n[i].CreateTime = msg.CreateTime.In(location)
	}
	return
}
