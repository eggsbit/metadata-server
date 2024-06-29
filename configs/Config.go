package configs

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	envFile = ".env"
)

type Config struct {
	MongodbConfig struct {
		Host         string
		Port         string
		DatabaseName string
		User         string
		Password     string
	}
	MongodbCollection struct {
		NftCollectionCollection string
		NftItemCollection       string
	}
	RedisConfig struct {
		Host     string
		Port     string
		Password string
	}
	MetadataServerConfig struct {
		Port string
	}
	ApplicationConfig struct {
		NftItemImageBaseUrl    string
		ExportPngSettingsDpi   string
		ExportFolderPath       string
		DeployWalletAddress    string
		NftCollectionAddress   string
		TonBlockchainConfigUrl string
	}
}

func NewConfig() (*Config, error) {
	config := new(Config)
	err := config.loadConfiguration()
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) loadConfiguration() error {
	err := c.loadEnvFileIfExist()
	if err != nil {
		return err
	}

	c.loadMongodbConfiguration()
	c.loadMongodbCollectionConfiguration()
	c.loadRedisConfiguration()
	c.loadMetadataServerConfiguration()
	c.loadApplicationConfiguration()

	return nil
}

func (c *Config) loadEnvFileIfExist() error {
	_, err := os.Stat(envFile)
	if err == nil {
		err := godotenv.Load(envFile)
		if err == nil {
			return err
		}
	}

	return nil
}

func (c *Config) loadMongodbConfiguration() {
	c.MongodbConfig.Host = os.Getenv("MONGODB_HOST")
	c.MongodbConfig.Port = os.Getenv("MONGODB_PORT")
	c.MongodbConfig.DatabaseName = os.Getenv("MONGODB_DATABASE_NAME")
	c.MongodbConfig.User = os.Getenv("MONGODB_USER")
	c.MongodbConfig.Password = os.Getenv("MONGODB_PASSWORD")
}

func (c *Config) loadMongodbCollectionConfiguration() {
	c.MongodbCollection.NftCollectionCollection = os.Getenv("MONGODB_NFT_COLLECTION_COLLECTION")
	c.MongodbCollection.NftItemCollection = os.Getenv("MONGODB_NFT_ITEM_COLLECTION")
}

func (c *Config) loadRedisConfiguration() {
	c.RedisConfig.Host = os.Getenv("REDIS_HOST")
	c.RedisConfig.Port = os.Getenv("REDIS_PORT")
	c.RedisConfig.Password = os.Getenv("REDIS_PASSWORD")
}

func (c *Config) loadMetadataServerConfiguration() {
	c.MetadataServerConfig.Port = os.Getenv("WEB_APPLICATION_PORT")
}

func (c *Config) loadApplicationConfiguration() {
	c.ApplicationConfig.NftItemImageBaseUrl = os.Getenv("NFT_ITEM_IMAGE_BASE_URL")
	c.ApplicationConfig.ExportPngSettingsDpi = os.Getenv("EXPORT_PNG_SETTINGS_DPI")
	c.ApplicationConfig.ExportFolderPath = os.Getenv("EXPORT_FOLDER_PATH")
	c.ApplicationConfig.DeployWalletAddress = os.Getenv("DEPLOY_WALLET_ADDRESS")
	c.ApplicationConfig.NftCollectionAddress = os.Getenv("NFT_COLLECTION_ADDRESS")
	c.ApplicationConfig.TonBlockchainConfigUrl = os.Getenv("TON_BLOCKCHAIN_CONFIG_URL")
}
