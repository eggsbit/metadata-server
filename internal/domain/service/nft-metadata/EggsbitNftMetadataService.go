package nftmetadata

import (
	"context"

	eggsbitnftdata "github.com/eggsbit/metadata-server/internal/domain/builder/eggsbit-nft-data"
	"github.com/eggsbit/metadata-server/internal/domain/entity"
	"github.com/eggsbit/metadata-server/internal/domain/repository"
)

func NewEggsbitNftMetadataService(
	eggsbitNftCollectionRepository repository.EggsbitNftCollectionDocRepositoryInterface,
	eggsbitNftItemRepository repository.EggsbitNftItemDocRepositoryInterface,
	nftItemBuilder eggsbitnftdata.NftItemBuilderInterface,
	imageFileBuilder eggsbitnftdata.ImageFileBuilderInterface,
) EggsbitNftMetadataServiceInterface {
	return &EggsbitNftMetadataService{
		eggsbitNftCollectionRepository: eggsbitNftCollectionRepository,
		eggsbitNftItemRepository:       eggsbitNftItemRepository,
		nftItemBuilder:                 nftItemBuilder,
		imageFileBuilder:               imageFileBuilder,
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
	imagePath, imagePathErr := enms.imageFileBuilder.CreateRandomStartingEggImage()

	var eggsbitNftItem entity.EggsbitNftItem
	if imagePathErr != nil {
		eggsbitNftItem = enms.nftItemBuilder.BuildByIndex(index)
	} else {
		eggsbitNftItem = enms.nftItemBuilder.BuildByIndexAndImage(index, *imagePath)
	}

	enms.eggsbitNftItemRepository.Add(eggsbitNftItem, ctx)

	return &eggsbitNftItem, nil
}
