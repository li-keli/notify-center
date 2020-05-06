package constant

type PlatformType int
type TargetType int

var (
	IOS         PlatformType = 10
	Android     PlatformType = 20
	MiniProgram PlatformType = 30 // discard
	WebSocket   PlatformType = 40 // discard
	DingDing    PlatformType = 50

	DemoApp TargetType = 100
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
	case DemoApp:
		return "演示APP"
	}
	return ""
}
