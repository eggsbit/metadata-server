package repository

import (
	"context"

	"github.com/eggsbit/metadata-server/internal/domain/entity"
)

type NftCollectionDocRepositoryInterface interface {
	GetCollectionByIdentifier(identifier string, ctx context.Context) (*entity.NftCollection, error)
}
