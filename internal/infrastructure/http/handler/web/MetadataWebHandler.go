package web

import (
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"regexp"

	"github.com/eggsbit/metadata-server/internal/application/builder"
	nftmetadata "github.com/eggsbit/metadata-server/internal/domain/service/nft-metadata"
	"github.com/eggsbit/metadata-server/internal/infrastructure/http/response"
	log "github.com/eggsbit/metadata-server/internal/infrastructure/logger"
	redisstore "github.com/eggsbit/metadata-server/internal/infrastructure/redis-store"
	"github.com/gin-gonic/gin"
)

func NewMetadataWebHandler(
	eggsbitNftMetadataService nftmetadata.EggsbitNftMetadataServiceInterface,
	collectionResponseBuilder builder.EggsbitNftCollectionMetadataResponseBuilderInterface,
	itemResponseBuilder builder.EggsbitNftItemMetadataResponseBuilderInterface,
	redisService redisstore.RedisServiceInterface,
	logger log.LoggerInterface,
) *MetadataWebHandler {
	return &MetadataWebHandler{
		eggsbitNftMetadataService: eggsbitNftMetadataService,
		collectionResponseBuilder: collectionResponseBuilder,
		itemResponseBuilder:       itemResponseBuilder,
		redisService:              redisService,
		logger:                    logger,
	}
}

type MetadataWebHandler struct {
	eggsbitNftMetadataService nftmetadata.EggsbitNftMetadataServiceInterface
	collectionResponseBuilder builder.EggsbitNftCollectionMetadataResponseBuilderInterface
	itemResponseBuilder       builder.EggsbitNftItemMetadataResponseBuilderInterface
	redisService              redisstore.RedisServiceInterface
	logger                    log.LoggerInterface
}

func (mwh *MetadataWebHandler) HandleCollectionData(ctx *gin.Context) {
	eggsbitCollectionIdentifier := "eggsbit_collection"

	// Get a value from redis if it exists
	redisValue, redisGetErr := mwh.redisService.GetValue(eggsbitCollectionIdentifier, ctx)
	if redisGetErr == nil {
		var responseFromRedis response.CollectionMetadataWebResponse
		jsonDecodeErr := json.Unmarshal([]byte(*redisValue), &responseFromRedis)
		if jsonDecodeErr == nil {
			ctx.JSON(http.StatusOK, responseFromRedis)
			return
		}
	}

	collectionEntity, collectionEntityErr := mwh.eggsbitNftMetadataService.GetCollectionByIdentifier(eggsbitCollectionIdentifier, ctx)
	if collectionEntityErr != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	// Set a value to redis
	response := mwh.collectionResponseBuilder.BuildResponse(*collectionEntity)
	mwh.saveResponseInRedis(eggsbitCollectionIdentifier, response, ctx)

	ctx.JSON(http.StatusOK, response)
}

func (mwh *MetadataWebHandler) HandleItemData(ctx *gin.Context) {
	index, parsingErr := mwh.getNftItemIndex(ctx)

	if parsingErr != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	eggsbitCollectionIdentifier := "eggsbit_collection"

	// Get a value from redis if it exists
	itemRedisKey := eggsbitCollectionIdentifier + "_item_" + index.String()
	redisValue, redisGetErr := mwh.redisService.GetValue(itemRedisKey, ctx)
	if redisGetErr == nil {
		var responseFromRedis response.ItemMetadataWebResponse
		jsonDecodeErr := json.Unmarshal([]byte(*redisValue), &responseFromRedis)
		if jsonDecodeErr == nil {
			ctx.JSON(http.StatusOK, responseFromRedis)
			return
		}
	}

	itemEntity, err := mwh.eggsbitNftMetadataService.GetNftItemByIndex(index, eggsbitCollectionIdentifier, ctx)
	if err != nil {
		mwh.logger.Error(log.LogCategoryLogic, err.Error())
		ctx.Status(http.StatusNotFound)
		return
	}

	// Set a value to redis
	response := mwh.itemResponseBuilder.BuildResponse(*itemEntity)
	mwh.saveResponseInRedis(itemRedisKey, response, ctx)

	ctx.JSONP(http.StatusOK, response)
}

func (mwh *MetadataWebHandler) getNftItemIndex(ctx *gin.Context) (*big.Int, error) {
	re := regexp.MustCompile(`^item_(0|[1-9][0-9]*)\.json$`)
	match := re.FindStringSubmatch(ctx.Param("item_slug"))
	if len(match) > 0 {
		newBigInt := new(big.Int)
		newBigInt, newBigIntStatus := newBigInt.SetString(match[1], 16)
		if !newBigIntStatus {
			newBigIntError := errors.New("nft item index from string to bigInt was failed")
			mwh.logger.Error(log.LogCategoryInputData, newBigIntError.Error())
			return nil, newBigIntError
		}

		return newBigInt, nil
	} else {
		return nil, errors.New("item slug is not correct")
	}
}

func (mwh *MetadataWebHandler) saveResponseInRedis(key string, obj any, ctx *gin.Context) {
	responseJson, jsonEncodeErr := json.Marshal(obj)
	if jsonEncodeErr == nil {
		redisSetErr := mwh.redisService.SetValue(key, string(responseJson), ctx)
		if redisSetErr != nil {
			mwh.logger.Error(log.LogCategorySystem, redisSetErr.Error())
		}
	} else {
		mwh.logger.Error(log.LogCategorySystem, jsonEncodeErr.Error())
	}
}
