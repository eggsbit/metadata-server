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

func NewColorSchemeDocRepository(
	mongodb mongodb.MongodbInterface,
	logger log.LoggerInterface,
	config *configs.Config,
) rep_interface.ColorSchemeDocRepositoryInterface {
	collection := mongodb.GetClient().Database(config.MongodbConfig.DatabaseName).Collection("color_schemes")
	return &ColorSchemeDocRepository{mongodb: mongodb, collection: collection, logger: logger}
}

type ColorSchemeDocRepository struct {
	mongodb    mongodb.MongodbInterface
	collection *mongo.Collection
	logger     log.LoggerInterface
}

func (csdr *ColorSchemeDocRepository) GetColorSchemeByIdentifier(identifier string, ctx context.Context) (*entity.ColorScheme, error) {
	filter := bson.D{primitive.E{Key: "identifier", Value: identifier}}

	var appColorScheme *entity.ColorScheme

	err := csdr.collection.FindOne(ctx, filter).Decode(&appColorScheme)
	if err != nil && err == mongo.ErrNoDocuments {
		return nil, err
	}

	return appColorScheme, nil
}
