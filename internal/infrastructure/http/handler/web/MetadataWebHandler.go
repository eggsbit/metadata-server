package web

import (
	"net/http"

	"github.com/eggsbit/metadata-server/internal/application/builder"
	"github.com/eggsbit/metadata-server/internal/domain/service/nftmetadata"
	"github.com/gin-gonic/gin"
)

func NewMetadataWebHandler() *MetadataWebHandler {
	return &MetadataWebHandler{}
}

type MetadataWebHandler struct {
	nftMetadataService        nftmetadata.NftMetadataServiceInterface
	collectionResponseBuilder builder.EggsbitNftCollectionMetadataResponseBuilder
	itemResponseBuilder       builder.EggsbitNftItemnMetadataResponseBuilder
}

func (mwh MetadataWebHandler) HandleItemData(ctx *gin.Context) {
	// redis check
	// get response from service
	ctx.JSONP(http.StatusOK, mwh.itemResponseBuilder.BuildErrorResponse(err.Error()))
}

func (mwh MetadataWebHandler) HandleCollectionData(ctx *gin.Context) {
	// redis check
	// get response from service
	ctx.JSONP(http.StatusOK, mwh.collectionResponseBuilder.BuildErrorResponse(err.Error()))
}
