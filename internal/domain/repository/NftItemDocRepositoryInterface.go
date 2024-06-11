package repository

import (
	"context"

	"github.com/eggsbit/metadata-server/internal/domain/entity"
)

type NftItemDocRepositoryInterface interface {
	GetItemByIndex(index string, collectionIdentifier string, ctx context.Context) (*entity.NftItem, error)

	Add(nftItem entity.NftItem, ctx context.Context) error
}
