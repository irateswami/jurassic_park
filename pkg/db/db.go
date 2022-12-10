package db

import (
	"database/sql"
	"io"

	"github.com/irateswami/jurassic_park/pkg/objects"
	_ "github.com/mattn/go-sqlite3"
)

type LiteDb struct {
	db *sql.DB
}

func NewLiteDb() (LiteDb, error) {
	newDb, err := sql.Open("sqlite3", "./jurassic_park.db")
	return LiteDb{db: newDb}, err
}

func NewLiteDbInMem() (*sql.DB, error) {
	return sql.Open("sqlite3", ":memory:")
}

func (dbl LiteDb) GetDino(rc io.ReadCloser) (objects.Dinosaur, error) {
	return objects.Dinosaur{}, nil
}
func (dbl LiteDb) PutDino(rc io.ReadCloser) error {
	return nil
}
func (dbl LiteDb) PostDino(rc io.ReadCloser) (objects.Dinosaur, error) {
	return objects.Dinosaur{}, nil
}
func (dbl LiteDb) DeleteDino(rc io.ReadCloser) {}
func (dbl LiteDb) GetCage(rc io.ReadCloser) (objects.Cage, error) {
	return objects.Cage{}, nil
}
func (dbl LiteDb) PutCage(rc io.ReadCloser)    {}
func (dbl LiteDb) PostCage(rc io.ReadCloser)   {}
func (dbl LiteDb) DeleteCage(rc io.ReadCloser) {}
