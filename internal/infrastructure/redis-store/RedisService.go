package redisstore

import (
	"context"
	"errors"

	log "github.com/eggsbit/metadata-server/internal/infrastructure/logger"
	"github.com/redis/go-redis/v9"
)

func NewRedisService(redis RedisInterface, logger log.LoggerInterface) RedisServiceInterface {
	return &RedisService{
		redis:  redis,
		logger: logger,
	}
}

type RedisServiceInterface interface {
	SetValue(key string, value string, ctx context.Context) error

	GetValue(key string, ctx context.Context) (*string, error)

	DeleteValue(key string, ctx context.Context) error
}

type RedisService struct {
	redis  RedisInterface
	logger log.LoggerInterface
}

func (rs RedisService) SetValue(key string, value string, ctx context.Context) error {
	err := rs.redis.GetClient().Set(ctx, key, value, 0).Err()
	if err != nil {
		rs.logger.Error(log.LogCategorySystem, err.Error())
		return err
	}

	return nil
}

func (rs RedisService) GetValue(key string, ctx context.Context) (*string, error) {
	value, err := rs.redis.GetClient().Get(ctx, key).Result()
	switch {
	case err == redis.Nil:
		rs.logger.Error(log.LogCategorySystem, "key does not exist")
		return nil, err
	case err != nil:
		rs.logger.Error(log.LogCategorySystem, "Get failed")
		return nil, err
	case value == "":
		emptyValueErr := errors.New("value is empty")
		rs.logger.Error(log.LogCategorySystem, emptyValueErr.Error())
		return nil, emptyValueErr
	}

	return &value, nil
}

func (rs RedisService) DeleteValue(key string, ctx context.Context) error {
	err := rs.redis.GetClient().Del(ctx, key).Err()
	if err != nil {
		rs.logger.Error(log.LogCategorySystem, err.Error())
		return err
	}

	return nil
}
