package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresHost      string
	PostgresPort      string
	PostgresUser      string
	PostgresPassword  string
	PostgresDb        string
	ServerPort        string
	ClientID          string
	ClientSecret      string
	ClientCallbackUrl string
}

func NewConfig() *Config {
	return &Config{}
}

func (config *Config) LoadEnvironment() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	// Postgres configuration
	config.PostgresHost = os.Getenv("POSTGRES_HOST")
	config.PostgresPort = os.Getenv("POSTGRES_PORT")
	config.PostgresUser = os.Getenv("POSTGRES_USER")
	config.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	config.PostgresDb = os.Getenv("POSTGRES_DB")

	// Server configuration
	config.ServerPort = os.Getenv("SERVER_PORT")

	// OAuth2.0 configuration
	config.ClientID = os.Getenv("CLIENT_ID")
	config.ClientSecret = os.Getenv("CLIENT_SECRET")
	config.ClientCallbackUrl = os.Getenv("CLIENT_CALLBACK_URL")

	return nil
}
