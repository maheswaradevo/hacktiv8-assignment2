package ping

import (
	"context"

	"github.com/maheswaradevo/hacktiv8-assignment2/internal/ping/impl"
)

type PingService interface {
	Ping(ctx context.Context) string
}

func NewPingService() PingService {
	return impl.ProvideServicePing()
}
