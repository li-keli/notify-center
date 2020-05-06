package db

import (
	"time"
)

type NotifyRegisterWeChat struct {
	Id          int       `gorm:"column:id;AUTO_INCREMENT"`
	JsjUniqueId int       `gorm:"column:jsj_unique_id"`
	PushToken   string    `gorm:"column:push_token"`
	CreateTime  time.Time `gorm:"column:create_time"`
}

func (NotifyRegisterWeChat) Insert(r NotifyRegisterWeChat) (err error) {
	err = conn.Create(&r).Error
	return
}

func (NotifyRegisterWeChat) FindOne(uniqueId int) (r NotifyRegisterWeChat, err error) {
	err = conn.Where(`jsj_unique_id = ?`, uniqueId).Order(`create_time desc`).First(&r).Error
	return
}

func (NotifyRegisterWeChat) DelOne(id int) (err error) {
	err = conn.Delete(NotifyRegisterWeChat{}, `id = ?`, id).Error
	return
}
