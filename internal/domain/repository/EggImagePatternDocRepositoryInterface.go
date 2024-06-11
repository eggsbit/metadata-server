package repository

import (
	"context"

	"github.com/eggsbit/metadata-server/internal/domain/entity"
)

type EggImagePatternDocRepositoryInterface interface {
	GetImagePatternByIdentifier(identifier string, ctx context.Context) (*entity.EggImagePattern, error)
}
