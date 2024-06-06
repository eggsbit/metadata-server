package nftmetadata

import (
	"context"

	"github.com/eggsbit/metadata-server/internal/domain/entity"
	"github.com/eggsbit/metadata-server/internal/domain/repository"
)

func NewNftMetadataService(
	nftCollectionRepository repository.NftCollectionDocRepositoryInterface,
	nftItemRepository repository.NftItemDocRepositoryInterface,
) NftMetadataServiceInterface {
	return &NftMetadataService{nftCollectionRepository: nftCollectionRepository, nftItemRepository: nftItemRepository}
}

type NftMetadataServiceInterface interface {
	GetNftCollectionByIndex(index string, ctx context.Context) (*entity.EggsbitNftCollection, error)

	GetNftItemByIndex(index string, ctx context.Context) (*entity.EggsbitNftItem, error)
}

type NftMetadataService struct {
	nftCollectionRepository repository.NftCollectionDocRepositoryInterface
	nftItemRepository       repository.NftItemDocRepositoryInterface
}

func (nms *NftMetadataService) GetNftCollectionByIndex(index string, ctx context.Context) (*entity.EggsbitNftCollection, error) {
	// check dbs
	// return
	entity := entity.EggsbitNftCollection{}
	return &entity, nil
}

func (nms *NftMetadataService) GetNftItemByIndex(index string, ctx context.Context) (*entity.EggsbitNftItem, error) {
	// check db
	// check ton chain collection index
	// create a new one
	// a call to generate image
	// return
	entity := entity.EggsbitNftItem{}
	return &entity, nil
}
