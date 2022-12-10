package main

import (
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/irateswami/jurassic_park/pkg/db"
	"github.com/irateswami/jurassic_park/pkg/handlers"
)

func run() error {

	router := gin.Default()
	storage, err := db.NewLiteDb()
	if err != nil {
		return err
	}

	// Dinosaurs
	router.GET("/dino", func(ctx *gin.Context) {
		handlers.GetDino(storage, ctx)
	})
	router.PUT("/dino", func(ctx *gin.Context) {})
	router.POST("/dino", func(ctx *gin.Context) {})
	router.DELETE("/dino", func(ctx *gin.Context) {})

	// Cages
	router.GET("/cages", func(ctx *gin.Context) {
		handlers.GetDino(storage, ctx)
	})
	router.PUT("/cages", func(ctx *gin.Context) {})
	router.POST("/cages", func(ctx *gin.Context) {})
	router.DELETE("/cages", func(ctx *gin.Context) {})

	return router.Run()
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	if err := run(); err != nil {
		logger.Fatal("runner failed", zap.Error(err))
	}
}
