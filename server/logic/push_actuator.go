package logic

// 推送逻辑
type PushActuator interface {
	PushMessage(pushKey string)
}

func BuildPushActuator(t int) PushActuator {
	if t == 400 {
		return &PushWSocket{}
	}
	return nil
}
