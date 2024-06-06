package builder

import (
	"github.com/eggsbit/metadata-server/internal/domain/entity"
	"github.com/eggsbit/metadata-server/internal/infrastructure/http/response"
)

func NewEggsbitNftCollectionMetadataResponseBuilder() EggsbitNftCollectionMetadataResponseBuilderInterface {
	return &EggsbitNftCollectionMetadataResponseBuilder{}
}

type EggsbitNftCollectionMetadataResponseBuilderInterface interface {
	BuildResponse(entity entity.EggsbitNftCollection) response.MetadataWebResponse
}

type EggsbitNftCollectionMetadataResponseBuilder struct {
}

func (encmrb *EggsbitNftCollectionMetadataResponseBuilder) BuildResponse(entity entity.EggsbitNftCollection) response.MetadataWebResponse {
	var response = make(map[string]string)

	response["name"] = "gogo"
	response["description"] = "gogo-description"
	response["time"] = "gogo-time"

	return response
}
