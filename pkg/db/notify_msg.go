package db

import (
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
