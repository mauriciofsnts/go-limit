package ctx

import (
	"context"

	"github.com/mauriciofsnts/gofast/pkg/config"
)

type Key string

const (
	ServicesKey Key = "services"
)

type Services struct {
	Config *config.Config
}

func GetServices(ctx context.Context) *Services {
	return ctx.Value(ServicesKey).(*Services)
}
