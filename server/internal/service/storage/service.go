package storage

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"mime/multipart"
	"storage_api/internal/config"
	"storage_api/internal/types"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type StorageService struct {
	config config.Config
}

func NewStorageService(cfg config.Config) *StorageService {
	return &StorageService{
		config: cfg,
	}
}

func (s *StorageService) ListFiles() ([]types.File, error) {
	// Initialize MinIO client
	minioClient, err := minio.New(s.config.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(s.config.Minio.AccessKey, s.config.Minio.SecretKey, ""),
		Secure: false,
	})

	if err != nil {
		log.Fatalln(err)
		return nil, fmt.Errorf("cannot connect to minio: %w", err)
	}

	files := make([]types.File, 0)

	for object := range minioClient.ListObjects(context.Background(), s.config.Storage.Bucket, minio.ListObjectsOptions{}) {
		files = append(files, types.File{
			Name: object.Key,
			URL: s.buildUrl(object.Key),
		})
	}
	
	return files, nil
}

func (s *StorageService) SaveFile(file *multipart.FileHeader) (string, error) {
	// Initialize MinIO client
	minioClient, err := minio.New(s.config.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(s.config.Minio.AccessKey, s.config.Minio.SecretKey, ""),
		Secure: false,
	})

	if err != nil {
		log.Fatalln(err)
		return "", fmt.Errorf("cannot connect to minio: %w", err)
	}

	// Bucket name
	bucketName := s.config.Storage.Bucket

	// Create a new bucket if it doesn't exist
	ctx := context.Background()
	// TODO: Provider region from env
	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "eu-west-1"})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			slog.Info("Bucket already exists: %s", bucketName)
		} else {
			fmt.Println("doopa 2")
			log.Fatalln(err)
			return "", err
		}
	} else {
		fmt.Println("doopa 3")
		slog.Info("Successfully created bucket: %s", bucketName)
	}

	reader, err := file.Open()
	if err != nil {
		return "", err
	}

	info, err := minioClient.PutObject(ctx, bucketName, file.Filename, reader, file.Size, minio.PutObjectOptions{
		ContentType: "text/plain",
	})
	if err != nil {
		return "", err
	}

	return s.buildUrl(info.Key), nil
}

func (s *StorageService) buildUrl(key string) string {
	return fmt.Sprintf("http://%s/%s/%s", s.config.Minio.Endpoint, s.config.Storage.Bucket, key)
}
