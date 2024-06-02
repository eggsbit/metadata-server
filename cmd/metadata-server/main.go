package main

import (
	"github.com/eggsbit/metadata-server/configs"
	"github.com/eggsbit/metadata-server/internal/infrastructure/di/metadata-server"
	"go.uber.org/fx"
)

func main() {
	config, err := configs.NewConfig()
	if err != nil {
		panic("Can't read configuration file: " + err.Error())
	}

	app := NewMetadataServer(config)
	app.Run()
}

func NewMetadataServer(config *configs.Config) *fx.App {
	app := fx.New(
		metadataserver.MetadataServerModule{}.BuildOptions(config),
	)

	return app
}
