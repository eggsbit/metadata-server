package metadataserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/eggsbit/metadata-server/api/server"
	router "github.com/eggsbit/metadata-server/api/server/router"
	"github.com/eggsbit/metadata-server/configs"
	"github.com/eggsbit/metadata-server/internal/infrastructure/di/common"
	apiHandler "github.com/eggsbit/metadata-server/internal/infrastructure/http/handler/api"
	webHandler "github.com/eggsbit/metadata-server/internal/infrastructure/http/handler/web"
	log "github.com/eggsbit/metadata-server/internal/infrastructure/logger"
	"go.uber.org/fx"
)

type MetadataServerModule struct{}

func (msm MetadataServerModule) BuildOptions(config *configs.Config) fx.Option {
	options := fx.Options(
		common.CommonModule{}.BuildOptions(config),
		fx.Provide(
			server.NewRouter,
			server.NewServer,
			router.NewWebRouter,
			router.NewApiRouter,
			apiHandler.NewNftCollectionApiHandler,
			apiHandler.NewNftItemApiHandler,
			webHandler.NewMetadataWebHandler,
		),
		fx.Invoke(
			server.RegisterRoutes,
			msm.startApplicationServer,
		),
	)

	return options
}

func (msm MetadataServerModule) startApplicationServer(lc fx.Lifecycle, server *http.Server, logger log.LoggerInterface) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					logger.Info(log.LogCategorySystem, "Starting HTTP Server at "+server.Addr)
					err := server.ListenAndServe()
					if err != nil {
						message := fmt.Sprintf("Failed to start HTTP Server at %s! Error: %s", server.Addr, err)
						logger.Error(log.LogCategorySystem, message)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				server.Shutdown(ctx)
				logger.Info(log.LogCategorySystem, "HTTP Server is stopped")
				return nil
			},
		},
	)
}
