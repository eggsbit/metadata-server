package builder

import (
	"github.com/eggsbit/metadata-server/internal/domain/entity"
	"github.com/eggsbit/metadata-server/internal/infrastructure/http/response"
)

func NewEggsbitNftItemMetadataResponseBuilder() EggsbitNftItemMetadataResponseBuilderInterface {
	return &EggsbitNftItemMetadataResponseBuilder{}
}

type EggsbitNftItemMetadataResponseBuilderInterface interface {
	BuildResponse(entity entity.EggsbitNftItem) response.ItemMetadataWebResponse
}

type EggsbitNftItemMetadataResponseBuilder struct {
}

func (enimrb *EggsbitNftItemMetadataResponseBuilder) BuildResponse(entity entity.EggsbitNftItem) response.ItemMetadataWebResponse {
	return response.ItemMetadataWebResponse{
		Name:        entity.Name,
		Description: entity.Description,
		Image:       entity.Image,
		Lottie:      entity.Lottie,
	}
}
