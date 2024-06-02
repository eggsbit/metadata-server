package server

import (
	"fmt"
	"net/http"

	"github.com/eggsbit/metadata-server/api/server/router"
	"github.com/eggsbit/metadata-server/configs"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func NewServer(router *gin.Engine, config *configs.Config) *http.Server {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.MetadataServerConfig.Port),
		Handler: router,
	}

	return server
}

func RegisterRoutes(
	router *gin.Engine,
	apiRouter *router.ApiRouter,
	webRouter *router.WebRouter,
) {
	apiRouter.Create(router)
	webRouter.Create(router)
}
