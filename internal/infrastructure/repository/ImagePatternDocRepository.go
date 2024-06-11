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

func NewImagePatternDocRepository(
	mongodb mongodb.MongodbInterface,
	logger log.LoggerInterface,
	config *configs.Config,
) rep_interface.ImagePatternDocRepositoryInterface {
	collection := mongodb.GetClient().Database(config.MongodbConfig.DatabaseName).Collection("image_pattern")
	return &ImagePatternDocRepository{mongodb: mongodb, collection: collection, logger: logger}
}

type ImagePatternDocRepository struct {
	mongodb    mongodb.MongodbInterface
	collection *mongo.Collection
	logger     log.LoggerInterface
}

func (eipdr *ImagePatternDocRepository) GetImagePatternByIdentifier(identifier string, ctx context.Context) (*entity.ImagePattern, error) {
	filter := bson.D{primitive.E{Key: "identifier", Value: identifier}}

	var appEggImagePattern *entity.ImagePattern

	err := eipdr.collection.FindOne(ctx, filter).Decode(&appEggImagePattern)
	if err != nil && err == mongo.ErrNoDocuments {
		return nil, err
	}

	return appEggImagePattern, nil
}
