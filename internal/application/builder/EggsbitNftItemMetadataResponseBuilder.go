package builder

import (
	"github.com/eggsbit/metadata-server/internal/domain/entity"
	"github.com/eggsbit/metadata-server/internal/infrastructure/http/response"
)

func NewEggsbitNftItemMetadataResponseBuilder() EggsbitNftItemMetadataResponseBuilderInterface {
	return &EggsbitNftItemMetadataResponseBuilder{}
}

type EggsbitNftItemMetadataResponseBuilderInterface interface {
	BuildResponse(entity entity.EggsbitNftItem) response.MetadataWebResponse
}

type EggsbitNftItemMetadataResponseBuilder struct {
}

func (enimrb *EggsbitNftItemMetadataResponseBuilder) BuildResponse(entity entity.EggsbitNftItem) response.MetadataWebResponse {
	var response = make(map[string]string)

	response["name"] = "olol"
	response["description"] = "keke"
	response["time"] = "gogo-timerr"

	return response
}
