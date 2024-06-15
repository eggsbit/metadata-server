package mongodb

import (
	"context"
	"fmt"

	"github.com/eggsbit/metadata-server/configs"
	log "github.com/eggsbit/metadata-server/internal/infrastructure/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongodbConnection(config *configs.Config, logger log.LoggerInterface) (MongodbInterface, error) {
	client, err := getMongodbClient(config, logger)
	if err != nil {
		return nil, err
	}

	return &MongodbConnection{client: client}, nil
}

func getMongodbClient(config *configs.Config, logger log.LoggerInterface) (*mongo.Client, error) {
	var ctx = context.TODO()
	clientOptions := options.Client().ApplyURI(getConnectionUri(config))
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		logger.Error(log.LogCategorySystem, err.Error())
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Error(log.LogCategorySystem, err.Error())
		return nil, err
	}

	return client, nil
}

func getConnectionUri(config *configs.Config) string {
	var connectionUri string

	if config.MongodbConfig.User != "" {
		connectionUri = fmt.Sprintf(
			"mongodb://%s:%s@%s:%s/admin",
			config.MongodbConfig.User,
			config.MongodbConfig.Password,
			config.MongodbConfig.Host,
			config.MongodbConfig.Port,
		)
	} else {
		connectionUri = fmt.Sprintf(
			"mongodb://%s:%s/",
			config.MongodbConfig.Host,
			config.MongodbConfig.Port,
		)
	}

	return connectionUri
}

type MongodbInterface interface {
	GetClient() *mongo.Client
}

type MongodbConnection struct {
	client *mongo.Client
}

func (mc *MongodbConnection) GetClient() *mongo.Client {
	return mc.client
}
