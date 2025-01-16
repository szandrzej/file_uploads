package server

import (
	"mime/multipart"
	"net/http"
	"storage_api/internal/config"
	"storage_api/internal/types"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type StorageService interface {
	SaveFile(file *multipart.FileHeader) (string, error)
	ListFiles() ([]types.File, error)
}

func RunServer(config config.Config, storage StorageService) {
	
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	
	apiRouter := r.Group("/api")
	apiRouter.Use(AuthMiddleware(config))
	apiRouter.GET("/files/", func(c *gin.Context) {
		files, err := storage.ListFiles()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "File uploaded successfully!", "files": files})
	})
	apiRouter.POST("/files/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
			return
		}

		info, err := storage.SaveFile(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "File uploaded successfully!", "url": info})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
