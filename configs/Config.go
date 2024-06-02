package configs

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	envFile = ".env"
)

type Config struct {
	DatabaseConfig struct {
		Host     string
		Port     string
		Name     string
		User     string
		Password string
	}
	RedisConfig struct {
		Host     string
		Port     string
		Password string
	}
	MetadataServerConfig struct {
		Port string
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

	c.loadDatabaseConfiguration()
	c.loadRedisConfiguration()
	c.loadMetadataServerConfiguration()

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

func (c *Config) loadDatabaseConfiguration() {
	c.DatabaseConfig.Host = os.Getenv("DATABASE_HOST")
	c.DatabaseConfig.Port = os.Getenv("DATABASE_PORT")
	c.DatabaseConfig.Name = os.Getenv("DATABASE_NAME")
	c.DatabaseConfig.User = os.Getenv("DATABASE_USER")
	c.DatabaseConfig.Password = os.Getenv("DATABASE_PASSWORD")
}

func (c *Config) loadRedisConfiguration() {
	c.DatabaseConfig.Host = os.Getenv("REDIS_HOST")
	c.DatabaseConfig.Port = os.Getenv("REDIS_PORT")
	c.DatabaseConfig.Password = os.Getenv("REDIS_PASSWORD")
}

func (c *Config) loadMetadataServerConfiguration() {
	c.MetadataServerConfig.Port = os.Getenv("WEB_APPLICATION_PORT")
}
