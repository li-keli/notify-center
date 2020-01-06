// 配置
package db

import (
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
}

type IosConfig struct {
	Production bool
	P12Path    string
	BundleId   string
	Password   string

	P12DevPath  string
	P12ProdPath string
}

type AndroidConfig struct {
	AppKey string
	Secret string
}

func (*NotifyConfig) FindConfig() {

}
