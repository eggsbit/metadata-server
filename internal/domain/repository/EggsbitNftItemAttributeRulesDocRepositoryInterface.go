package repository

import (
	"context"

	"github.com/eggsbit/metadata-server/internal/domain/entity"
)

type EggsbitNftItemAttributeRulesDocRepositoryInterface interface {
	GetRulesByParentIdentifier(identifier string, rule_type string, ctx context.Context) ([]*entity.EggsbitNftItemAttributeRule, error)
}
