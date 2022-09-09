package impl

import "context"

type pingServiceImpl struct {
}

func (ps pingServiceImpl) Ping(ctx context.Context) string {
	return "pong"
}

func ProvideServicePing() *pingServiceImpl {
	return &pingServiceImpl{}
}
