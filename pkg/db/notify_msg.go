package db

import (
	"notify-center/pkg/constant"
)

type NotifyMsg struct {
	Id               int                   `json:"Id" gorm:"column:id;AUTO_INCREMENT"`
	JsjUniqueId      int                   `json:"JsjUniqueId" gorm:"column:jsj_unique_id"`
	PushToken        string                `json:"PushToken" gorm:"column:push_token"`
	Title            string                `json:"Title" gorm:"column:title"`
	Message          string                `json:"Message" gorm:"column:message"`
	Router           string                `json:"Router" gorm:"column:router"`
	DataContent      string                `json:"DataContent" gorm:"column:data_content"`
	PlatformTypeId   constant.PlatformType `json:"PlatformTypeId" gorm:"column:platform_type_id"`
	PlatformTypeName string                `json:"PlatformTypeName" gorm:"column:platform_type_name"`
	TargetTypeId     constant.TargetType   `json:"TargetTypeId" gorm:"column:target_type_id"`
	TargetTypeName   string                `json:"TargetTypeName" gorm:"column:target_type_name"`
	GroupName        string                `json:"GroupName" gorm:"column:group_name"`
	CreateTime       constant.JsonTime     `json:"CreateTime" gorm:"column:create_time"`
}

func (NotifyMsg) Insert(m NotifyMsg) (err error) {
	err = conn.Create(&m).Error
	return
}

func (NotifyMsg) FindAll(jsjUniqueId int, token string, offset, limit int) (n []NotifyMsg, err error) {
	err = conn.Where(`jsj_unique_id = ? or push_token = ?`, jsjUniqueId, token).Offset(offset - 1).Limit(limit).Order(`create_time desc`).Find(&n).Error
	return
}
