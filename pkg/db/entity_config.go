// 配置
package db

import (
	"errors"
	"notify-center/pkg/constant"
)

type DicConfigEntity struct {
	PlatformType  constant.TargetType `json:"PlatformType"`
	IosConfig     IosConfig           `json:"IosConfig"`
	AndroidConfig AndroidConfig       `json:"AndroidConfig"`
}

type IosConfig struct {
	P12Path  string `json:"P12Path"`
	BundleId string `json:"BundleId"`
	Password string `json:"Password"`
}

type AndroidConfig struct {
	AppKey string `json:"AppKey"`
	Secret string `json:"Secret"`
}

func (*DicConfigEntity) FindConfig(platformType constant.TargetType) (DicConfigEntity, error) {
	switch platformType {
	case constant.YgApp:
		return DicConfigEntity{
			PlatformType: constant.YgApp,
			IosConfig: IosConfig{
				BundleId: "com.jsj.AirRailwayMannagerForEmployee",
				P12Path:  "resources/yg_app_debug.p12",
				Password: "1",
			},
			AndroidConfig: AndroidConfig{
				AppKey: "10aa9bf027e3378b5a212350",
				Secret: "2041e32d04ec4f32a4dc7e94",
			},
		}, nil
	}
	return DicConfigEntity{}, errors.New("未发现配置")
}
