package nftmetadata

import (
	"context"

	eggsbitnftdata "github.com/eggsbit/metadata-server/internal/domain/builder/eggsbit-nft-data"
	"github.com/eggsbit/metadata-server/internal/domain/entity"
	"github.com/eggsbit/metadata-server/internal/domain/repository"
	log "github.com/eggsbit/metadata-server/internal/infrastructure/logger"
)

func NewEggsbitNftMetadataService(
	eggsbitNftCollectionRepository repository.EggsbitNftCollectionDocRepositoryInterface,
	eggsbitNftItemRepository repository.EggsbitNftItemDocRepositoryInterface,
	nftItemBuilder eggsbitnftdata.NftItemBuilderInterface,
	imageFileBuilder eggsbitnftdata.ImageFileBuilderInterface,
	logger log.LoggerInterface,
) EggsbitNftMetadataServiceInterface {
	return &EggsbitNftMetadataService{
		eggsbitNftCollectionRepository: eggsbitNftCollectionRepository,
		eggsbitNftItemRepository:       eggsbitNftItemRepository,
		nftItemBuilder:                 nftItemBuilder,
		imageFileBuilder:               imageFileBuilder,
		logger:                         logger,
	}
}

type EggsbitNftMetadataServiceInterface interface {
	GetCollectionByIdentifier(indentifier string, ctx context.Context) (*entity.EggsbitNftCollection, error)

	GetNftItemByIndex(index string, ctx context.Context) (*entity.EggsbitNftItem, error)
}

type EggsbitNftMetadataService struct {
	eggsbitNftCollectionRepository repository.EggsbitNftCollectionDocRepositoryInterface
	eggsbitNftItemRepository       repository.EggsbitNftItemDocRepositoryInterface
	nftItemBuilder                 eggsbitnftdata.NftItemBuilderInterface
	imageFileBuilder               eggsbitnftdata.ImageFileBuilderInterface
	logger                         log.LoggerInterface
}

func (enms *EggsbitNftMetadataService) GetCollectionByIdentifier(indentifier string, ctx context.Context) (*entity.EggsbitNftCollection, error) {
	// check dbs
	// return
	entity, err := enms.eggsbitNftCollectionRepository.GetCollectionByIdentifier(indentifier, ctx)
	return entity, err
}

func (enms *EggsbitNftMetadataService) GetNftItemByIndex(index string, ctx context.Context) (*entity.EggsbitNftItem, error) {
	itemEntity, err := enms.eggsbitNftItemRepository.GetItemByIndex(index, ctx)
	if err == nil {
		return itemEntity, err
	}

	// check ton chain collection index
	//imagePath, imagePathErr := enms.imageFileBuilder.CreateRandomStartingEggImage()

	eggsbitNftItem := enms.nftItemBuilder.BuildStartEggByIndex(index, ctx)
	enms.eggsbitNftItemRepository.Add(eggsbitNftItem, ctx)

	return &eggsbitNftItem, nil
}
