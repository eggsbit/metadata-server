package api

import (
	"github.com/gin-gonic/gin"
)

func NewNftCollectionApiHandler() *NftCollectionApiHandler {
	return &NftCollectionApiHandler{}
}

type NftCollectionApiHandler struct {
}

func (ncah NftCollectionApiHandler) HandleAddCollectionData(ctx *gin.Context) {
}

func (ncah NftCollectionApiHandler) HandleUpdateCollectionData(ctx *gin.Context) {
}
