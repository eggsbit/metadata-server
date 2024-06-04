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

func NewNftCollectionDocRepository(
	mongodb mongodb.MongodbInterface,
	logger log.LoggerInterface,
	config *configs.Config,
) rep_interface.NftCollectionDocRepositoryInterface {
	collection := mongodb.GetClient().Database(config.MongodbConfig.DatabaseName).Collection(config.MongodbCollection.NftCollectionCollection)
	return &NftCollectionDocRepository{mongodb: mongodb, collection: collection, logger: logger}
}

type NftCollectionDocRepository struct {
	mongodb    mongodb.MongodbInterface
	collection *mongo.Collection
	logger     log.LoggerInterface
}

func (ncdr *NftCollectionDocRepository) GetCollectionByIndex(index string, ctx context.Context) (*entity.EggsbitNftCollection, error) {
	filter := bson.D{primitive.E{Key: "index", Value: index}}

	var appNftCollection *entity.EggsbitNftCollection

	ncdr.collection.FindOne(ctx, filter).Decode(&appNftCollection)

	return appNftCollection, bson.ErrDecodeToNil
}
