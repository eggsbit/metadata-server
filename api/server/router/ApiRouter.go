package router

import (
	"github.com/eggsbit/metadata-server/internal/infrastructure/http/handler/api"
	"github.com/gin-gonic/gin"
)

func NewApiRouter(
	nftCollectionApiHandler *api.NftCollectionApiHandler,
	nftItemApiHandler *api.NftItemApiHandler,
) *ApiRouter {
	return &ApiRouter{
		nftCollectionApiHandler: nftCollectionApiHandler,
		nftItemApiHandler:       nftItemApiHandler,
	}
}

type ApiRouter struct {
	nftCollectionApiHandler *api.NftCollectionApiHandler
	nftItemApiHandler       *api.NftItemApiHandler
}

func (ar *ApiRouter) Create(router *gin.Engine) {
	collectionApiRouterGroup := router.Group("/api/v1/nft-collection/")
	collectionApiRouterGroup.POST("/", ar.nftCollectionApiHandler.HandleAddCollectionData)
	collectionApiRouterGroup.PUT("/", ar.nftCollectionApiHandler.HandleUpdateCollectionData)

	itemApiRouterGroup := router.Group("/api/v1/nft-item/")
	itemApiRouterGroup.POST("/{:id}/action/born", ar.nftItemApiHandler.HandleActionBorn)
}
