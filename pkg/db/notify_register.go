// app 配置
package db

import (
	"notify-center/pkg/constant"
	"time"
)

type NotifyRegister struct {
	Id               int                   `gorm:"column:id;AUTO_INCREMENT"`
	JsjUniqueId      int                   `gorm:"column:jsj_unique_id"`
	PushToken        string                `gorm:"column:push_token"`
	PlatformTypeId   constant.PlatformType `gorm:"column:platform_type_id"`
	PlatformTypeName string                `gorm:"column:platform_type_name"`
	TargetTypeId     constant.TargetType   `gorm:"column:target_type_id"`
	TargetTypeName   string                `gorm:"column:target_type_name"`
	CreateTime       time.Time             `gorm:"column:create_time"`
	UpdateTime       time.Time             `gorm:"column:update_time"`
}

func (NotifyRegister) Insert(r NotifyRegister) (err error) {
	err = conn.Create(&r).Error
	return
}

func (NotifyRegister) FindOne(uniqueId int) (r NotifyRegister, err error) {
	err = conn.Where(`jsj_unique_id = ?`, uniqueId).First(&r).Error
	return
}

func (NotifyRegister) UpdateSpecify(r NotifyRegister) (err error) {
	err = conn.Model(&r).Where(`jsj_unique_id = ?`, r.JsjUniqueId).Updates(map[string]interface{}{
		"push_token":         r.PushToken,
		"platform_type_id":   r.PlatformTypeId,
		"platform_type_name": r.PlatformTypeName,
		"target_type_id":     r.TargetTypeId,
		"target_type_name":   r.TargetTypeName,
		"update_time":        r.UpdateTime,
	}).Error
	return
}

func (NotifyRegister) DelOne(uniqueId int, platformId constant.PlatformType, targetTypeId constant.TargetType) (err error) {
	err = conn.Where(`jsj_unique_id = ? and platform_type_id = ? and target_type_id = ?`, uniqueId, platformId, targetTypeId).
		Delete(&NotifyRegister{}).Error
	return
}
