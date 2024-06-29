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

func NewMiniAppDocRepository(
	mongodb mongodb.MongodbInterface,
	logger log.LoggerInterface,
	config *configs.Config,
) rep_interface.MiniAppDocRepositoryInterface {
	collection := mongodb.GetClient().Database(config.MongodbConfig.DatabaseName).Collection("mini_apps")
	return &MiniAppDocRepository{mongodb: mongodb, collection: collection, logger: logger}
}

type MiniAppDocRepository struct {
	mongodb    mongodb.MongodbInterface
	collection *mongo.Collection
	logger     log.LoggerInterface
}

func (madr *MiniAppDocRepository) GetMiniAppByIdentifier(identifier string, ctx context.Context) (*entity.MiniApp, error) {
	filter := bson.D{primitive.E{Key: "identifier", Value: identifier}}

	var appMiniApp *entity.MiniApp

	err := madr.collection.FindOne(ctx, filter).Decode(&appMiniApp)
	if err != nil && err == mongo.ErrNoDocuments {
		return nil, err
	}

	return appMiniApp, nil
}
