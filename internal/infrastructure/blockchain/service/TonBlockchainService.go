package service

import (
	"context"
	"math/big"

	"github.com/eggsbit/metadata-server/configs"
	domainBlockchain "github.com/eggsbit/metadata-server/internal/domain/service/blockchain"
	"github.com/eggsbit/metadata-server/internal/infrastructure/blockchain"
	log "github.com/eggsbit/metadata-server/internal/infrastructure/logger"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/ton/nft"
)

func NewTonBlockchainService(
	config *configs.Config,
	logger log.LoggerInterface,
	blockchainConnection blockchain.TonBlockchainConnectionInterface,
) domainBlockchain.TonBlockchainServiceInterface {
	return &TonBlockchainService{
		config:               config,
		logger:               logger,
		blockchainConnection: blockchainConnection,
	}
}

type TonBlockchainService struct {
	config               *configs.Config
	logger               log.LoggerInterface
	blockchainConnection blockchain.TonBlockchainConnectionInterface
}

func (tbs TonBlockchainService) GetCollectionData() (*nft.CollectionData, error) {
	collectionAddress := address.MustParseAddr(tbs.config.ApplicationConfig.NftCollectionAddress)
	collection := nft.NewCollectionClient(tbs.blockchainConnection.GetClient(), collectionAddress)
	return collection.GetCollectionData(context.Background())
}

func (tbs TonBlockchainService) GetCollectionNextItemIndex() (*big.Int, error) {
	collectionData, collectionDataErr := tbs.GetCollectionData()

	if collectionDataErr != nil {
		return nil, collectionDataErr
	}

	return collectionData.NextItemIndex, nil
}
