package web

import (
	"net/http"

	"github.com/eggsbit/metadata-server/internal/application/builder"
	nftmetadata "github.com/eggsbit/metadata-server/internal/domain/service/nft-metadata"
	"github.com/gin-gonic/gin"
)

func NewMetadataWebHandler(
	nftMetadataService nftmetadata.NftMetadataServiceInterface,
	collectionResponseBuilder builder.EggsbitNftCollectionMetadataResponseBuilderInterface,
	itemResponseBuilder builder.EggsbitNftItemMetadataResponseBuilderInterface,
) *MetadataWebHandler {
	return &MetadataWebHandler{
		nftMetadataService:        nftMetadataService,
		collectionResponseBuilder: collectionResponseBuilder,
		itemResponseBuilder:       itemResponseBuilder,
	}
}

type MetadataWebHandler struct {
	nftMetadataService        nftmetadata.NftMetadataServiceInterface
	collectionResponseBuilder builder.EggsbitNftCollectionMetadataResponseBuilderInterface
	itemResponseBuilder       builder.EggsbitNftItemMetadataResponseBuilderInterface
}

func (mwh *MetadataWebHandler) HandleItemData(ctx *gin.Context) {
	// redis check
	// get response from service
	itemEntity, _ := mwh.nftMetadataService.GetNftItemByIndex("1", ctx)
	ctx.JSONP(http.StatusOK, mwh.itemResponseBuilder.BuildResponse(*itemEntity))
}

func (mwh *MetadataWebHandler) HandleCollectionData(ctx *gin.Context) {
	// redis check
	// get response from service
	collectionEntity, _ := mwh.nftMetadataService.GetNftCollectionByIndex("1", ctx)
	ctx.JSONP(http.StatusOK, mwh.collectionResponseBuilder.BuildResponse(*collectionEntity))
}
