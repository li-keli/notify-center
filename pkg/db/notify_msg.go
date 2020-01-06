package db

import (
	"github.com/sirupsen/logrus"
	"notify-center/pkg/constant"
	"time"
)

type NotifyMsg struct {
	Id               int                   `gorm:"column:id;AUTO_INCREMENT"`
	PushToken        string                `gorm:"column:push_token"`
	DataContent      string                `gorm:"column:data_content"`
	PlatformTypeId   constant.PlatformType `gorm:"column:platform_type_id"`
	PlatformTypeName string                `gorm:"column:platform_type_name"`
	TargetTypeId     constant.TargetType   `gorm:"column:target_type_id"`
	TargetTypeName   string                `gorm:"column:target_type_name"`
	CreateTime       time.Time             `gorm:"column:create_time"`
}

func (NotifyMsg) Insert(m NotifyMsg) {
	if err := conn.Create(&m).Error; err != nil {
		logrus.Error("历史消息存储错误，但忽略了这个错误")
	}
}
