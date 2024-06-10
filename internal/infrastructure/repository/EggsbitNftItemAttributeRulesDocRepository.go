package repository

import (
	"context"

	"github.com/eggsbit/metadata-server/configs"
	"github.com/eggsbit/metadata-server/internal/domain/entity"
	rep_interface "github.com/eggsbit/metadata-server/internal/domain/repository"
	log "github.com/eggsbit/metadata-server/internal/infrastructure/logger"
	"github.com/eggsbit/metadata-server/internal/infrastructure/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewEggsbitNftItemAttributeRulesDocRepository(
	mongodb mongodb.MongodbInterface,
	logger log.LoggerInterface,
	config *configs.Config,
) rep_interface.EggsbitNftItemAttributeRulesDocRepositoryInterface {
	collection := mongodb.GetClient().Database(config.MongodbConfig.DatabaseName).Collection("egssbit_nft_item_attribute_rules")
	return &EggsbitNftItemAttributeRulesDocRepository{mongodb: mongodb, collection: collection, logger: logger}
}

type EggsbitNftItemAttributeRulesDocRepository struct {
	mongodb    mongodb.MongodbInterface
	collection *mongo.Collection
	logger     log.LoggerInterface
}

func (enidr *EggsbitNftItemAttributeRulesDocRepository) GetRulesByParentIdentifier(identifier string, rule_type string, ctx context.Context) ([]*entity.EggsbitNftItemAttributeRule, error) {
	filter := bson.D{primitive.E{Key: "parent_identifier", Value: identifier}, primitive.E{Key: "type", Value: rule_type}}

	cursor, err := enidr.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var attributeRules []*entity.EggsbitNftItemAttributeRule

	for cursor.Next(ctx) {
		var attributeRule entity.EggsbitNftItemAttributeRule
		if err := cursor.Decode(&attributeRule); err != nil {
			enidr.logger.Error(log.LogCategorySystem, err.Error())
			return nil, err
		}

		attributeRules = append(attributeRules, &attributeRule)
	}

	cursor.Close(ctx)

	if err != nil && err == mongo.ErrNoDocuments {
		return nil, err
	}

	return attributeRules, nil
}
