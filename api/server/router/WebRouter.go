package router

import (
	"github.com/eggsbit/metadata-server/internal/infrastructure/http/handler/web"
	"github.com/gin-gonic/gin"
)

func NewWebRouter(
	metadataWebHandler *web.MetadataWebHandler,
) *WebRouter {
	return &WebRouter{
		metadataWebHandler: metadataWebHandler,
	}
}

type WebRouter struct {
	metadataWebHandler *web.MetadataWebHandler
}

func (wr *WebRouter) Create(router *gin.Engine) {
	webRouterGroup := router.Group("/meta/eggsbit/")
	webRouterGroup.GET("/collection.json", wr.metadataWebHandler.HandleCollectionData)
	webRouterGroup.GET("/:item_slug", wr.metadataWebHandler.HandleItemData)
	webRouterGroup.GET("/tonconnect-manifest.json", wr.metadataWebHandler.HandleMiniAppData)
}
