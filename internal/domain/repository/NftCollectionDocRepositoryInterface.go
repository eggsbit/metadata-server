package repository

import (
	"context"

	"github.com/eggsbit/metadata-server/internal/domain/entity"
)

type NftCollectionDocRepositoryInterface interface {
	GetCollectionByIndex(index string, ctx context.Context) (*entity.EggsbitNftCollection, error)
}
