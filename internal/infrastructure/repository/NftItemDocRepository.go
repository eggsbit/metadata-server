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
	collection := mongodb.GetClient().Database(config.MongodbConfig.DatabaseName).Collection("nft_items")
	return &NftItemDocRepository{mongodb: mongodb, collection: collection, logger: logger}
}

type NftItemDocRepository struct {
	mongodb    mongodb.MongodbInterface
	collection *mongo.Collection
	logger     log.LoggerInterface
}

func (nidr *NftItemDocRepository) GetItemByIndex(index string, collectionIdentifier string, ctx context.Context) (*entity.NftItem, error) {
	filter := bson.D{primitive.E{Key: "index", Value: index}, primitive.E{Key: "collection_identifier", Value: collectionIdentifier}}

	var appNftItem *entity.NftItem

	err := nidr.collection.FindOne(ctx, filter).Decode(&appNftItem)
	if err != nil && err == mongo.ErrNoDocuments {
		return nil, err
	}

	return appNftItem, nil
}

func (nidr *NftItemDocRepository) Add(nftItem entity.NftItem, ctx context.Context) error {
	insertResult, err := nidr.collection.InsertOne(ctx, nftItem)

	if err != mongo.ErrNilCursor {
		return err
	}

	if _, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		return nil
	} else {
		return err
	}
}
