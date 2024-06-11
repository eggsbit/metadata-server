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

func NewEggImagePatternDocRepository(
	mongodb mongodb.MongodbInterface,
	logger log.LoggerInterface,
	config *configs.Config,
) rep_interface.EggImagePatternDocRepositoryInterface {
	collection := mongodb.GetClient().Database(config.MongodbConfig.DatabaseName).Collection("egg_image_pattern")
	return &EggImagePatternDocRepository{mongodb: mongodb, collection: collection, logger: logger}
}

type EggImagePatternDocRepository struct {
	mongodb    mongodb.MongodbInterface
	collection *mongo.Collection
	logger     log.LoggerInterface
}

func (eipdr *EggImagePatternDocRepository) GetImagePatternByIdentifier(identifier string, ctx context.Context) (*entity.EggImagePattern, error) {
	filter := bson.D{primitive.E{Key: "identifier", Value: identifier}}

	var appEggImagePattern *entity.EggImagePattern

	err := eipdr.collection.FindOne(ctx, filter).Decode(&appEggImagePattern)
	if err != nil && err == mongo.ErrNoDocuments {
		return nil, err
	}

	return appEggImagePattern, nil
}
