package miniappmetadata

import (
	"context"

	"github.com/eggsbit/metadata-server/internal/domain/entity"
	"github.com/eggsbit/metadata-server/internal/domain/repository"
	log "github.com/eggsbit/metadata-server/internal/infrastructure/logger"
)

func NewEggsbitMiniAppMetadataService(
	miniAppRepository repository.MiniAppDocRepositoryInterface,
	logger log.LoggerInterface,
) EggsbitMiniAppMetadataServiceInterface {
	return &EggsbitMiniAppMetadataService{
		miniAppRepository: miniAppRepository,
		logger:            logger,
	}
}

type EggsbitMiniAppMetadataServiceInterface interface {
	GetMiniAppByIdentifier(indentifier string, ctx context.Context) (*entity.MiniApp, error)
}

type EggsbitMiniAppMetadataService struct {
	miniAppRepository repository.MiniAppDocRepositoryInterface
	logger            log.LoggerInterface
}

func (emams *EggsbitMiniAppMetadataService) GetMiniAppByIdentifier(indentifier string, ctx context.Context) (*entity.MiniApp, error) {
	entity, err := emams.miniAppRepository.GetMiniAppByIdentifier(indentifier, ctx)
	return entity, err
}
