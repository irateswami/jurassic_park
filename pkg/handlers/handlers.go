package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/irateswami/jurassic_park/pkg/db"
	"go.uber.org/zap"
)

var (
	LOGGER *zap.Logger
)

func init() {
	LOGGER, _ = zap.NewProduction()
}

func GetDino(s db.Storage, ctx *gin.Context) {

	dino, err := s.GetDino(ctx.Request.Body)
	if err != nil {
		//handle the error
	}

	ctx.JSON(200, dino)
}

func PutDino(s db.Storage, ctx *gin.Context) {
	defer LOGGER.Sync()
	if err := s.PutDino(ctx); err != nil {
		LOGGER.Error("put dino error", zap.Error(err))
		ctx.Status(500)
		return
	}

	ctx.Status(200)
}

func PostDino(s db.Storage, ctx *gin.Context) {
}

func DeleteDino(s db.Storage, ctx *gin.Context) {
}

func GetCage(s db.Storage, ctx *gin.Context) {
}

func PutCage(s db.Storage, ctx *gin.Context) {
}

func PostCage(s db.Storage, ctx *gin.Context) {
}

func DeleteCage(s db.Storage, ctx *gin.Context) {
}
