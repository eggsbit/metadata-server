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

func NewEggsbitNftCollectionDocRepository(
	mongodb mongodb.MongodbInterface,
	logger log.LoggerInterface,
	config *configs.Config,
) rep_interface.EggsbitNftCollectionDocRepositoryInterface {
	collection := mongodb.GetClient().Database(config.MongodbConfig.DatabaseName).Collection("nft_collection")
	return &EggsbitNftCollectionDocRepository{mongodb: mongodb, collection: collection, logger: logger}
}

type EggsbitNftCollectionDocRepository struct {
	mongodb    mongodb.MongodbInterface
	collection *mongo.Collection
	logger     log.LoggerInterface
}

func (encdr *EggsbitNftCollectionDocRepository) GetCollectionByIdentifier(identifier string, ctx context.Context) (*entity.EggsbitNftCollection, error) {
	filter := bson.D{primitive.E{Key: "identifier", Value: identifier}}

	var appNftCollection *entity.EggsbitNftCollection

	encdr.collection.FindOne(ctx, filter).Decode(&appNftCollection)

	return appNftCollection, bson.ErrDecodeToNil
}
