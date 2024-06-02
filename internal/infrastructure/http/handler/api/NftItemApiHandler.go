package api

import (
	"github.com/gin-gonic/gin"
)

func NewNftItemApiHandler() *NftItemApiHandler {
	return &NftItemApiHandler{}
}

type NftItemApiHandler struct {
}

func (niah NftItemApiHandler) HandleActionBorn(ctx *gin.Context) {
}
