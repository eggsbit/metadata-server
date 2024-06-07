package repository

import (
	"context"

	"github.com/eggsbit/metadata-server/internal/domain/entity"
)

type EggsbitNftCollectionDocRepositoryInterface interface {
	GetCollectionByIdentifier(index string, ctx context.Context) (*entity.EggsbitNftCollection, error)
}
