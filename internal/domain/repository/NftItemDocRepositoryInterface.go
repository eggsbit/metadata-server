package repository

import (
	"context"

	"github.com/eggsbit/metadata-server/internal/domain/entity"
)

type NftItemDocRepositoryInterface interface {
	GetItemByIndex(index string, ctx context.Context) (*entity.EggsbitNftItem, error)
}
