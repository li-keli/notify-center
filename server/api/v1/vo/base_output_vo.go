package vo

type BaseOutput struct {
	BaseHead BaseHead `json:"baseHead"`
}

type BaseHead struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (BaseOutput) Success() BaseOutput {
	return BaseOutput{
		BaseHead: BaseHead{
			Code:    "0000",
			Message: "",
		},
	}
}

func (BaseOutput) Error(msg string) BaseOutput {
	return BaseOutput{
		BaseHead: BaseHead{
			Code:    "0001",
			Message: msg,
		},
	}
}
