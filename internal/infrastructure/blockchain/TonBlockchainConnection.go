package blockchain

import (
	"context"

	"github.com/eggsbit/metadata-server/configs"
	log "github.com/eggsbit/metadata-server/internal/infrastructure/logger"

	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
)

func NewTonBlockchainConnection(config *configs.Config, logger log.LoggerInterface) (TonBlockchainConnectionInterface, error) {
	client, err := getTonBlockchainClient(config, logger)
	if err != nil {
		return nil, err
	}

	return &TonBlockchainConnection{client: client}, nil
}

func getTonBlockchainClient(config *configs.Config, logger log.LoggerInterface) (*ton.APIClient, error) {
	client := liteclient.NewConnectionPool()

	tonBlockchainConnectionErr := client.AddConnectionsFromConfigUrl(context.Background(), config.ApplicationConfig.TonBlockchainConfigUrl)
	if tonBlockchainConnectionErr != nil {
		logger.Error(log.LogCategorySystem, tonBlockchainConnectionErr.Error())
		return nil, tonBlockchainConnectionErr
	}

	return ton.NewAPIClient(client), nil
}

type TonBlockchainConnectionInterface interface {
	GetClient() *ton.APIClient
}

type TonBlockchainConnection struct {
	client *ton.APIClient
}

func (tbc *TonBlockchainConnection) GetClient() *ton.APIClient {
	return tbc.client
}
