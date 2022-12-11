package db

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/irateswami/jurassic_park/pkg/objects"
)

type Storage interface {
	GetDino(rc io.ReadCloser) (objects.Dinosaur, error)
	PutDino(*gin.Context) error
	PostDino(rc io.ReadCloser) (objects.Dinosaur, error)
	DeleteDino(rc io.ReadCloser)
	GetCage(rc io.ReadCloser) (objects.Cage, error)
	PutCage(rc io.ReadCloser)
	PostCage(rc io.ReadCloser)
	DeleteCage(rc io.ReadCloser)
}
