package vo

type RegisterKtInput struct {
	FormId      string `json:"FormId"`
	JsjUniqueId int    `json:"JsjUniqueId"`
	Prepay      bool   `json:"Prepay"`
	PushToken   string `json:"PushToken"`
}

type RegisterKtOutput struct {
	BaseHead struct {
		Code    string `json:"Code"`
		Message string `json:"Message"`
	} `json:"BaseHead"`
}
