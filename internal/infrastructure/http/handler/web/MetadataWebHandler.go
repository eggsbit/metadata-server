package web

import (
	"net/http"

	"github.com/eggsbit/metadata-server/internal/application/builder"
	nftmetadata "github.com/eggsbit/metadata-server/internal/domain/service/nft-metadata"
	"github.com/gin-gonic/gin"
)

func NewMetadataWebHandler(
	eggsbitNftMetadataService nftmetadata.EggsbitNftMetadataServiceInterface,
	collectionResponseBuilder builder.EggsbitNftCollectionMetadataResponseBuilderInterface,
	itemResponseBuilder builder.EggsbitNftItemMetadataResponseBuilderInterface,
) *MetadataWebHandler {
	return &MetadataWebHandler{
		eggsbitNftMetadataService: eggsbitNftMetadataService,
		collectionResponseBuilder: collectionResponseBuilder,
		itemResponseBuilder:       itemResponseBuilder,
	}
}

type MetadataWebHandler struct {
	eggsbitNftMetadataService nftmetadata.EggsbitNftMetadataServiceInterface
	collectionResponseBuilder builder.EggsbitNftCollectionMetadataResponseBuilderInterface
	itemResponseBuilder       builder.EggsbitNftItemMetadataResponseBuilderInterface
}

func (mwh *MetadataWebHandler) HandleItemData(ctx *gin.Context) {
	// redis check
	// get response from service
	itemEntity, _ := mwh.eggsbitNftMetadataService.GetNftItemByIndex("1", ctx)
	ctx.JSONP(http.StatusOK, mwh.itemResponseBuilder.BuildResponse(*itemEntity))
}

func (mwh *MetadataWebHandler) HandleCollectionData(ctx *gin.Context) {
	// redis check
	// get response from service
	collectionEntity, _ := mwh.eggsbitNftMetadataService.GetCollectionByIdentifier("eggsbit_collection", ctx)
	ctx.JSONP(http.StatusOK, mwh.collectionResponseBuilder.BuildResponse(*collectionEntity))
}
