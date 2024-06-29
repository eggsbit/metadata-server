package repository

import (
	"context"

	"github.com/eggsbit/metadata-server/internal/domain/entity"
)

type MiniAppDocRepositoryInterface interface {
	GetMiniAppByIdentifier(identifier string, ctx context.Context) (*entity.MiniApp, error)
}
