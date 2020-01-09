package vo

// 微信小程序订阅消息注册
type RegisterWeChatInputVo struct {
	JsjUniqueId int    `json:"JsjUniqueId"` // 金色世纪用户唯一编号
	PushToken   string `json:"PushToken"`   // 指推送目标的OpenId
	FormId      string `json:"FormId"`      // 微信小程序推送用formId
}
