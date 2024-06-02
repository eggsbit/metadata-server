package common

import (
	"github.com/eggsbit/metadata-server/configs"
	"github.com/eggsbit/metadata-server/internal/infrastructure/logger"
	"go.uber.org/fx"
)

type CommonModule struct{}

func (cm CommonModule) BuildOptions(config *configs.Config) fx.Option {
	options := fx.Options(
		fx.Provide(
			func() *configs.Config {
				return config
			},
			logger.NewLogger,
		),
	)

	return options
}
