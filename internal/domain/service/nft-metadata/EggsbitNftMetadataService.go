package nftmetadata

import (
	"context"

	"github.com/eggsbit/metadata-server/internal/domain/entity"
	"github.com/eggsbit/metadata-server/internal/domain/repository"
)

func NewEggsbitNftMetadataService(
	eggsbitNftCollectionRepository repository.EggsbitNftCollectionDocRepositoryInterface,
	eggsbitNftItemRepository repository.EggsbitNftItemDocRepositoryInterface,
) EggsbitNftMetadataServiceInterface {
	return &EggsbitNftMetadataService{
		eggsbitNftCollectionRepository: eggsbitNftCollectionRepository,
		eggsbitNftItemRepository:       eggsbitNftItemRepository,
	}
}

type EggsbitNftMetadataServiceInterface interface {
	GetCollectionByIdentifier(indentifier string, ctx context.Context) (*entity.EggsbitNftCollection, error)

	GetNftItemByIndex(index string, ctx context.Context) (*entity.EggsbitNftItem, error)
}

type EggsbitNftMetadataService struct {
	eggsbitNftCollectionRepository repository.EggsbitNftCollectionDocRepositoryInterface
	eggsbitNftItemRepository       repository.EggsbitNftItemDocRepositoryInterface
}

func (enms *EggsbitNftMetadataService) GetCollectionByIdentifier(indentifier string, ctx context.Context) (*entity.EggsbitNftCollection, error) {
	// check dbs
	// return
	entity, _ := enms.eggsbitNftCollectionRepository.GetCollectionByIdentifier(indentifier, ctx)
	return entity, nil
}

func (enms *EggsbitNftMetadataService) GetNftItemByIndex(index string, ctx context.Context) (*entity.EggsbitNftItem, error) {
	// check db
	// check ton chain collection index
	// create a new one
	// a call to generate image
	// return
	entity := entity.EggsbitNftItem{}
	return &entity, nil
}
