package repository

import (
	"context"

	"github.com/eggsbit/metadata-server/internal/domain/entity"
)

type EggsbitNftItemDocRepositoryInterface interface {
	GetItemByIndex(index string, ctx context.Context) (*entity.EggsbitNftItem, error)

	Add(eggsbitNftItem entity.EggsbitNftItem, ctx context.Context) error
}
