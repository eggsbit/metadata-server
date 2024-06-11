package builder

import (
	"github.com/eggsbit/metadata-server/internal/domain/entity"
	"github.com/eggsbit/metadata-server/internal/infrastructure/http/response"
)

func NewEggsbitNftCollectionMetadataResponseBuilder() EggsbitNftCollectionMetadataResponseBuilderInterface {
	return &EggsbitNftCollectionMetadataResponseBuilder{}
}

type EggsbitNftCollectionMetadataResponseBuilderInterface interface {
	BuildResponse(entity entity.NftCollection) response.CollectionMetadataWebResponse
}

type EggsbitNftCollectionMetadataResponseBuilder struct {
}

func (encmrb *EggsbitNftCollectionMetadataResponseBuilder) BuildResponse(entity entity.NftCollection) response.CollectionMetadataWebResponse {
	return response.CollectionMetadataWebResponse{
		Name:        entity.Name,
		Description: entity.Description,
		Image:       entity.Image,
		CoverImage:  entity.CoverImage,
		SocialLinks: entity.SocialLinks,
	}
}
