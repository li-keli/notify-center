package constant

type PlatformType int
type TargetType int

var (
	IOS         PlatformType = 10
	Android     PlatformType = 20
	MiniProgram PlatformType = 30
	WebSocket   PlatformType = 40 // discard
	DingDing    PlatformType = 50

	JsjApp     TargetType = 100
	KtApp      TargetType = 200
	YgApp      TargetType = 300
	KtgjWxMini TargetType = 400
	Dingding   TargetType = 500
	KtNo1      TargetType = 600
)

// 转换推送的平台类型
func PlatformTypeValueOf(pType PlatformType) string {
	switch pType {
	case IOS:
		return "IOS"
	case Android:
		return "Android"
	case MiniProgram:
		return "MiniProgram"
	case WebSocket:
		return "WebSocket"
	case DingDing:
		return "DingDing"
	}
	return ""
}

// 转换推送终端类型
func TargetTypeValueOf(tType TargetType) string {
	switch tType {
	case JsjApp:
		return "金色世纪APP"
	case KtApp:
		return "空铁管家APP"
	case YgApp:
		return "员工端APP"
	case KtgjWxMini:
		return "空铁管家微信小程序"
	case Dingding:
		return "钉钉"
	case KtNo1:
		return "空铁1号"
	}
	return ""
}
