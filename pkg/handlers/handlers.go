package handlers

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/irateswami/jurassic_park/pkg/db"
)

var (
	INFOLOGGER    *log.Logger
	WARNINGLOGGER *log.Logger
	ERRORLOGGER   *log.Logger
)

func init() {
	INFOLOGGER = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WARNINGLOGGER = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ERRORLOGGER = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func GetDino(s db.Storage, ctx *gin.Context) {

	dino, err := s.GetDino(ctx.Request.Body)
	if err != nil {
		//handle the error
	}

	ctx.JSON(200, dino)
}

func PutDino(s db.Storage, ctx *gin.Context) {
	if err := s.PutDino(ctx); err != nil {
		ERRORLOGGER.Printf("put dino error: %s\n", err)
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
	if err := s.PutCage(ctx); err != nil {
		ERRORLOGGER.Printf("put cage error: %s\n", err)
		ctx.Status(500)
		return
	}

	ctx.Status(200)
}

func PostCage(s db.Storage, ctx *gin.Context) {
}

func DeleteCage(s db.Storage, ctx *gin.Context) {
}
