package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort           string
	ClientOrigin         string
	AccessTokenPublicKey string
	DatabaseURL          string
	PostgresHost         string
	PostgresUser         string
	PostgresPassword     string
	PostgresDB           string
	PostgresPort         string
}

func LoadConfig(path string) (config Config, err error) {
	err = godotenv.Load(path + "/app.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config = Config{
		ServerPort:           getEnv("SERVER_PORT", "8080"),
		ClientOrigin:         getEnv("CLIENT_ORIGIN", "http://localhost:3000"),
		AccessTokenPublicKey: getEnv("ACCESS_TOKEN_PUBLIC_KEY", ""),
		DatabaseURL:          getEnv("DATABASE_URL", ""),
		PostgresHost:         getEnv("POSTGRES_HOST", "127.0.0.1"),
		PostgresUser:         getEnv("POSTGRES_USER", "postgres"),
		PostgresPassword:     getEnv("POSTGRES_PASSWORD", "password123"),
		PostgresDB:           getEnv("POSTGRES_DB", "golang-gorm"),
		PostgresPort:         getEnv("POSTGRES_PORT", "6500"),
	}

	if config.DatabaseURL == "" {
		config.DatabaseURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			config.PostgresUser, config.PostgresPassword, config.PostgresHost, config.PostgresPort, config.PostgresDB)
	}

	return
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
