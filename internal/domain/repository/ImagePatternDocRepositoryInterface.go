package repository

import (
	"context"

	"github.com/eggsbit/metadata-server/internal/domain/entity"
)

type ImagePatternDocRepositoryInterface interface {
	GetImagePatternByIdentifier(identifier string, ctx context.Context) (*entity.ImagePattern, error)
}
