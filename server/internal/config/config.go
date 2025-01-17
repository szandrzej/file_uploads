package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Minio struct {
	Endpoint string
	AccessKey string
	SecretKey string
}

type Storage struct {
	Bucket string
}

type Config struct {
	AuthToken string
	Minio Minio
	Storage Storage
}

func ReadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		// TODO: handle error better here
		return Config{}, err
	}

	return Config{
		AuthToken: os.Getenv("AUTHORIZATION_TOKEN"),
		Minio: Minio{
			Endpoint: os.Getenv("MINIO_ENDPOINT"),
			AccessKey: os.Getenv("MINIO_ACCESS_KEY_ID"),
			SecretKey: os.Getenv("MINIO_SECRET_ACCESS_KEY"),
		},
		Storage: Storage{
			Bucket: os.Getenv("STORAGE_BUCKET"),
		},
	}, nil
}
