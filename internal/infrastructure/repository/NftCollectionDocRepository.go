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
	collection := mongodb.GetClient().Database(config.MongodbConfig.DatabaseName).Collection("nft_collection")
	return &NftCollectionDocRepository{mongodb: mongodb, collection: collection, logger: logger}
}

type NftCollectionDocRepository struct {
	mongodb    mongodb.MongodbInterface
	collection *mongo.Collection
	logger     log.LoggerInterface
}

func (ncdr *NftCollectionDocRepository) GetCollectionByIdentifier(identifier string, ctx context.Context) (*entity.NftCollection, error) {
	filter := bson.D{primitive.E{Key: "identifier", Value: identifier}}

	var appNftCollection *entity.NftCollection

	err := ncdr.collection.FindOne(ctx, filter).Decode(&appNftCollection)
	if err != nil && err == mongo.ErrNoDocuments {
		return nil, err
	}

	return appNftCollection, nil
}
