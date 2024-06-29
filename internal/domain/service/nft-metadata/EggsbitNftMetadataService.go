package nftmetadata

import (
	"context"
	"errors"
	"math/big"

	eggsbitnftdata "github.com/eggsbit/metadata-server/internal/domain/builder/eggsbit-nft-data"
	"github.com/eggsbit/metadata-server/internal/domain/constant"
	"github.com/eggsbit/metadata-server/internal/domain/entity"
	"github.com/eggsbit/metadata-server/internal/domain/repository"
	"github.com/eggsbit/metadata-server/internal/domain/service/blockchain"
	log "github.com/eggsbit/metadata-server/internal/infrastructure/logger"
)

func NewEggsbitNftMetadataService(
	nftCollectionRepository repository.NftCollectionDocRepositoryInterface,
	nftItemRepository repository.NftItemDocRepositoryInterface,
	nftItemBuilder eggsbitnftdata.NftItemBuilderInterface,
	imageFileBuilder eggsbitnftdata.ImageFileBuilderInterface,
	tonBlockchainService blockchain.TonBlockchainServiceInterface,
	logger log.LoggerInterface,
) EggsbitNftMetadataServiceInterface {
	return &EggsbitNftMetadataService{
		nftCollectionRepository: nftCollectionRepository,
		nftItemRepository:       nftItemRepository,
		nftItemBuilder:          nftItemBuilder,
		imageFileBuilder:        imageFileBuilder,
		tonBlockchainService:    tonBlockchainService,
		logger:                  logger,
	}
}

type EggsbitNftMetadataServiceInterface interface {
	GetCollectionByIdentifier(indentifier string, ctx context.Context) (*entity.NftCollection, error)

	GetNftItemByIndex(index *big.Int, collectionIndentifier string, ctx context.Context) (*entity.NftItem, error)
}

type EggsbitNftMetadataService struct {
	nftCollectionRepository repository.NftCollectionDocRepositoryInterface
	nftItemRepository       repository.NftItemDocRepositoryInterface
	nftItemBuilder          eggsbitnftdata.NftItemBuilderInterface
	imageFileBuilder        eggsbitnftdata.ImageFileBuilderInterface
	tonBlockchainService    blockchain.TonBlockchainServiceInterface
	logger                  log.LoggerInterface
}

func (enms *EggsbitNftMetadataService) GetCollectionByIdentifier(indentifier string, ctx context.Context) (*entity.NftCollection, error) {
	entity, err := enms.nftCollectionRepository.GetCollectionByIdentifier(indentifier, ctx)
	return entity, err
}

func (enms *EggsbitNftMetadataService) GetNftItemByIndex(index *big.Int, collectionIndentifier string, ctx context.Context) (*entity.NftItem, error) {
	itemEntity, err := enms.nftItemRepository.GetItemByIndex(index.String(), collectionIndentifier, ctx)
	if err == nil {
		return itemEntity, err
	}

	collectionNextItemIndex, nextItemIndexErr := enms.tonBlockchainService.GetCollectionNextItemIndex()
	if nextItemIndexErr != nil {
		enms.logger.Error(log.LogCategorySystem, nextItemIndexErr.Error())
		return nil, nextItemIndexErr
	}

	if collectionNextItemIndex.Cmp(index) == -1 || collectionNextItemIndex.Cmp(index) == 0 {
		indexString := index.String()
		return nil, errors.New("index (#" + indexString + ") collection item doesn't exist")
	}

	eggsbitNftItem, imageUuid := enms.nftItemBuilder.BuildStartEggByIndex(index, ctx)
	createImageErr := enms.createStartingEggImage(imageUuid, eggsbitNftItem, ctx)
	if createImageErr != nil {
		enms.logger.Error(log.LogCategoryLogic, createImageErr.Error())
	}

	enms.nftItemRepository.Add(eggsbitNftItem, ctx)
	return &eggsbitNftItem, nil
}

func (enms *EggsbitNftMetadataService) createStartingEggImage(imageUuid string, eggsbitNftItem entity.NftItem, ctx context.Context) error {
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

	return enms.imageFileBuilder.CreateStartingEggImage(imageUuid, eggPattern, eggColorScheme, ctx)
}
