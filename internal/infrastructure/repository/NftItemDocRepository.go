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

func NewNftItemDocRepository(
	mongodb mongodb.MongodbInterface,
	logger log.LoggerInterface,
	config *configs.Config,
) rep_interface.NftItemDocRepositoryInterface {
	collection := mongodb.GetClient().Database(config.MongodbConfig.DatabaseName).Collection(config.MongodbCollection.NftItemCollection)
	return &NftItemDocRepository{mongodb: mongodb, collection: collection, logger: logger}
}

type NftItemDocRepository struct {
	mongodb    mongodb.MongodbInterface
	collection *mongo.Collection
	logger     log.LoggerInterface
}

func (nidr *NftItemDocRepository) GetItemByIndex(index string, ctx context.Context) (*entity.EggsbitNftItem, error) {
	filter := bson.D{primitive.E{Key: "index", Value: index}}

	var appNftItem *entity.EggsbitNftItem

	nidr.collection.FindOne(ctx, filter).Decode(&appNftItem)

	return appNftItem, bson.ErrDecodeToNil
}
