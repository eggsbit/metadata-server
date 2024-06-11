package repository

import (
	"context"

	"github.com/eggsbit/metadata-server/internal/domain/entity"
)

type ColorSchemeDocRepositoryInterface interface {
	GetColorSchemeByIdentifier(identifier string, ctx context.Context) (*entity.ColorScheme, error)
}
