package web

import (
	"github.com/gin-gonic/gin"
)

func NewMetadataWebHandler() *MetadataWebHandler {
	return &MetadataWebHandler{}
}

type MetadataWebHandler struct {
}

func (mwh MetadataWebHandler) HandleItemData(ctx *gin.Context) {
}

func (mwh MetadataWebHandler) HandleCollectionData(ctx *gin.Context) {
}
