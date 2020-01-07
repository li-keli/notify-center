// 配置
package db

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"notify-center/pkg/constant"
	"time"
)

type NotifyConfig struct {
	Id               int                   `gorm:"column:id;AUTO_INCREMENT"`
	PlatformTypeId   constant.PlatformType `gorm:"column:platform_type_id"`
	PlatformTypeName string                `gorm:"column:platform_type_name"`
	TargetTypeId     constant.TargetType   `gorm:"column:target_type_id"`
	TargetTypeName   string                `gorm:"column:target_type_name"`
	ConfigData       string                `gorm:"column:config_data"`
	CreateTime       time.Time             `gorm:"column:create_time"`
	ConfigDataModel  interface{}           `gorm:"-"` // 数据库映射忽略字段，此字段来源于ConfigData
}

type IosConfig struct {
	P12Path  string
	BundleId string
	Password string

	P12DevPath  string
	P12ProdPath string
}

type AndroidConfig struct {
	AppKey string
	Secret string
}

func (n NotifyConfig) IosConfig() (config IosConfig, err error) {
	if err = json.Unmarshal([]byte(n.ConfigData), &config); err != nil {
		logrus.Error("序列化IosConfig错误，原文：", n.ConfigData)
		return
	}
	if gin.Mode() == "debug" {
		config.P12Path = config.P12DevPath
	} else {
		config.P12Path = config.P12ProdPath
	}
	return
}

func (n NotifyConfig) AndroidConfig() (config AndroidConfig, err error) {
	if err := json.Unmarshal([]byte(n.ConfigData), &config); err != nil {
		logrus.Error("序列化AndroidConfig错误，原文：", n.ConfigData)
	}
	return
}

func (n NotifyConfig) FindOne(platformType constant.PlatformType, targetType constant.TargetType) (r NotifyConfig, err error) {
	err = conn.Model(&n).Where(`platform_type_id = ? and target_type_id = ?`, platformType, targetType).First(&r).Error
	return
}
