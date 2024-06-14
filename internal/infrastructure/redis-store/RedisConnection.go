package redisstore

import (
	"context"
	"fmt"

	"github.com/eggsbit/metadata-server/configs"
	log "github.com/eggsbit/metadata-server/internal/infrastructure/logger"

	"github.com/redis/go-redis/v9"
)

func NewRedisConnection(config *configs.Config, logger log.LoggerInterface) (RedisInterface, error) {
	client, err := getRedisClient(config)
	if err != nil {
		logger.Error(log.LogCategorySystem, "Error connecting to Redis: "+err.Error())
		return nil, err
	}

	return &RedisConnection{client: client}, nil
}

func getRedisClient(config *configs.Config) (*redis.Client, error) {
	var ctx = context.TODO()
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedisConfig.Host, config.RedisConfig.Port),
		Password: config.RedisConfig.Password,
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

type RedisInterface interface {
	GetClient() *redis.Client
}

type RedisConnection struct {
	client *redis.Client
}

func (rc *RedisConnection) GetClient() *redis.Client {
	return rc.client
}
