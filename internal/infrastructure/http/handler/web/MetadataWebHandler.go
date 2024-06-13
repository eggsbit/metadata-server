package web

import (
	"errors"
	"math/big"
	"net/http"
	"regexp"

	"github.com/eggsbit/metadata-server/internal/application/builder"
	nftmetadata "github.com/eggsbit/metadata-server/internal/domain/service/nft-metadata"
	log "github.com/eggsbit/metadata-server/internal/infrastructure/logger"
	"github.com/gin-gonic/gin"
)

func NewMetadataWebHandler(
	eggsbitNftMetadataService nftmetadata.EggsbitNftMetadataServiceInterface,
	collectionResponseBuilder builder.EggsbitNftCollectionMetadataResponseBuilderInterface,
	itemResponseBuilder builder.EggsbitNftItemMetadataResponseBuilderInterface,
	logger log.LoggerInterface,
) *MetadataWebHandler {
	return &MetadataWebHandler{
		eggsbitNftMetadataService: eggsbitNftMetadataService,
		collectionResponseBuilder: collectionResponseBuilder,
		itemResponseBuilder:       itemResponseBuilder,
		logger:                    logger,
	}
}

type MetadataWebHandler struct {
	eggsbitNftMetadataService nftmetadata.EggsbitNftMetadataServiceInterface
	collectionResponseBuilder builder.EggsbitNftCollectionMetadataResponseBuilderInterface
	itemResponseBuilder       builder.EggsbitNftItemMetadataResponseBuilderInterface
	logger                    log.LoggerInterface
}

func (mwh *MetadataWebHandler) HandleCollectionData(ctx *gin.Context) {
	eggsbitCollectionIdentifier := "eggsbit_collection"
	// redis check

	collectionEntity, err := mwh.eggsbitNftMetadataService.GetCollectionByIdentifier(eggsbitCollectionIdentifier, ctx)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.JSONP(http.StatusOK, mwh.collectionResponseBuilder.BuildResponse(*collectionEntity))
}

func (mwh *MetadataWebHandler) HandleItemData(ctx *gin.Context) {
	index, parsingErr := mwh.getNftItemIndex(ctx)

	if parsingErr != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	// redis check
	eggsbitCollectionIdentifier := "eggsbit_collection"
	itemEntity, err := mwh.eggsbitNftMetadataService.GetNftItemByIndex(index, eggsbitCollectionIdentifier, ctx)
	if err != nil {
		mwh.logger.Error(log.LogCategoryLogic, err.Error())
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.JSONP(http.StatusOK, mwh.itemResponseBuilder.BuildResponse(*itemEntity))
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
