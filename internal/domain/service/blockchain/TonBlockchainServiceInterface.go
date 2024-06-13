package blockchain

import (
	"math/big"

	"github.com/xssnick/tonutils-go/ton/nft"
)

type TonBlockchainServiceInterface interface {
	GetCollectionData() (*nft.CollectionData, error)

	GetCollectionNextItemIndex() (*big.Int, error)
}
