package builder

import (
	"github.com/eggsbit/metadata-server/internal/domain/entity"
	"github.com/eggsbit/metadata-server/internal/infrastructure/http/response"
)

func NewEggsbitMiniAppMetadataResponseBuilder() EggsbitMiniAppMetadataResponseBuilderInterface {
	return &EggsbitMiniAppMetadataResponseBuilder{}
}

type EggsbitMiniAppMetadataResponseBuilderInterface interface {
	BuildResponse(entity entity.MiniApp) response.MiniAppMetadataWebResponse
}

type EggsbitMiniAppMetadataResponseBuilder struct {
}

func (emamrb *EggsbitMiniAppMetadataResponseBuilder) BuildResponse(entity entity.MiniApp) response.MiniAppMetadataWebResponse {
	return response.MiniAppMetadataWebResponse{
		Url:              entity.Url,
		Name:             entity.Name,
		IconUrl:          entity.IconUrl,
		TermsOfUseUrl:    entity.TermsOfUseUrl,
		PrivacyPolicyUrl: entity.PrivacyPolicyUrl,
	}
}
