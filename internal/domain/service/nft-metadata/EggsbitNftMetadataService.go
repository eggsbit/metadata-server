package nftmetadata

import (
	"context"

	"github.com/eggsbit/metadata-server/configs"
	eggsbitnftdata "github.com/eggsbit/metadata-server/internal/domain/builder/eggsbit-nft-data"
	"github.com/eggsbit/metadata-server/internal/domain/constant"
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
	config *configs.Config,
) EggsbitNftMetadataServiceInterface {
	return &EggsbitNftMetadataService{
		eggsbitNftCollectionRepository: eggsbitNftCollectionRepository,
		eggsbitNftItemRepository:       eggsbitNftItemRepository,
		nftItemBuilder:                 nftItemBuilder,
		imageFileBuilder:               imageFileBuilder,
		logger:                         logger,
		config:                         config,
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
	config                         *configs.Config
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
	eggsbitNftItem, imageUuid := enms.nftItemBuilder.BuildStartEggByIndex(index, ctx)
	createImageErr := enms.createStartingEggImage(imageUuid, eggsbitNftItem)
	if createImageErr != nil {
		enms.logger.Error(log.LogCategoryLogic, createImageErr.Error())
	}

	enms.eggsbitNftItemRepository.Add(eggsbitNftItem, ctx)
	return &eggsbitNftItem, nil
}

func (enms *EggsbitNftMetadataService) createStartingEggImage(imageUuid string, eggsbitNftItem entity.EggsbitNftItem) error {
	var eggPattern string
	var eggColorScheme string

	for _, nftAttribute := range eggsbitNftItem.Attributes {
		if nftAttribute.TraitType == constant.KeyAttributePattern {
			eggPattern = *nftAttribute.Value
		}

		if nftAttribute.TraitType == constant.KeyAttributeColorSchema {
			eggColorScheme = *nftAttribute.Value
		}
	}

	return enms.imageFileBuilder.CreateStartingEggImage(imageUuid, eggPattern, eggColorScheme)
}
