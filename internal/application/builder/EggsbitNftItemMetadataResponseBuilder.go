package builder

import (
	"github.com/eggsbit/metadata-server/internal/domain/entity"
	"github.com/eggsbit/metadata-server/internal/infrastructure/http/response"

	"github.com/jinzhu/copier"
)

func NewEggsbitNftItemMetadataResponseBuilder() EggsbitNftItemMetadataResponseBuilderInterface {
	return &EggsbitNftItemMetadataResponseBuilder{}
}

type EggsbitNftItemMetadataResponseBuilderInterface interface {
	BuildResponse(entity entity.NftItem) response.ItemMetadataWebResponse
}

type EggsbitNftItemMetadataResponseBuilder struct {
}

func (enimrb *EggsbitNftItemMetadataResponseBuilder) BuildResponse(entity entity.NftItem) response.ItemMetadataWebResponse {
	attributes := []response.ItemAttributeWebResponse{}
	copier.Copy(&attributes, &entity.Attributes)

	return response.ItemMetadataWebResponse{
		Name:        entity.Name,
		Description: entity.Description,
		Image:       entity.Image,
		Lottie:      entity.Lottie,
		Attributes:  attributes,
	}
}
