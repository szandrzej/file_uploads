package main

import (
	"fmt"
	"log"
	"storage_api/internal/config"
	"storage_api/internal/server"
	"storage_api/internal/service/storage"
)


func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("cannot run program, cannot read config: %w", err)
	}
	fmt.Println(cfg);

	storageService := storage.NewStorageService(cfg)

	server.RunServer(cfg, storageService)
}
