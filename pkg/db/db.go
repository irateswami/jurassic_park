package db

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"github.com/irateswami/jurassic_park/pkg/objects"
	_ "github.com/mattn/go-sqlite3"
)

type LiteDb struct {
	db *sql.DB
}

const (
	DBTIMEOUT       = 10 * time.Second
	DINOSAURS_TABLE = "dinosaurs"
	CAGES_TABLE     = "cages"
)

type Storage interface {
	GetDino(rc io.ReadCloser) (objects.Dinosaur, error)
	PutDino(*gin.Context) error
	PostDino(rc io.ReadCloser) (objects.Dinosaur, error)
	DeleteDino(rc io.ReadCloser)
	GetCage(rc io.ReadCloser) (objects.Cage, error)
	PutCage(*gin.Context) error
	PostCage(rc io.ReadCloser)
	DeleteCage(rc io.ReadCloser)
}

// New DB
func NewLiteDb() (LiteDb, error) {
	newDb, err := sql.Open("sqlite3", "./pkg/db/jurassic_park.db")
	return LiteDb{db: newDb}, err
}

func NewLiteDbInMem() (LiteDb, error) {
	newDb, err := sql.Open("sqlite3", ":memory:")
	return LiteDb{db: newDb}, err
}

// Dinosaurs
func (store LiteDb) GetDino(rc io.ReadCloser) (objects.Dinosaur, error) {
	return objects.Dinosaur{}, nil
}

type ErrorSlice []error

func (es ErrorSlice) Error() string {
	var bigString string
	for _, err := range es {
		bigString += fmt.Sprint(" | " + err.Error() + " | ")
	}

	return bigString
}

func (store LiteDb) PutDino(c *gin.Context) error {

	// Try decoding our input
	dinos := []objects.Dinosaur{}
	err := c.ShouldBindJSON(&dinos)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", dinos)

	// Cancel timeout
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT)
	defer cancel()

	// Start a transaction
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	var transactionErrors ErrorSlice

	// Prepare the insert
	for _, val := range dinos {
		gi := goqu.Insert(DINOSAURS_TABLE).Cols(
			"id",
			"species",
			"name",
			"herb_or_carn",
			"cage",
		).Vals(
			goqu.Vals{
				val.Id,
				val.Species,
				val.Name,
				val.Cage,
			},
		)

		insertSQL, args, err := gi.ToSQL()
		if err != nil {
			transactionErrors = append(transactionErrors, err)
		}

		fmt.Println(insertSQL, args)

		_, err = tx.ExecContext(ctx, fmt.Sprint(insertSQL, args))
		if err != nil {
			transactionErrors = append(transactionErrors, err)
		}
	}

	if len(transactionErrors) != 0 {
		return transactionErrors
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
func (store LiteDb) PostDino(rc io.ReadCloser) (objects.Dinosaur, error) {
	return objects.Dinosaur{}, nil
}
func (store LiteDb) DeleteDino(rc io.ReadCloser) {}

// Cages
func (store LiteDb) GetCage(rc io.ReadCloser) (objects.Cage, error) {
	return objects.Cage{}, nil
}
func (store LiteDb) PutCage(c *gin.Context) error {

	cages := []objects.Cage{}
	err := c.ShouldBind(&cages)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT)
	defer cancel()

	// Start a transaction
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	var transactionErrors ErrorSlice

	// Prepare the insert
	for _, val := range cages {
		gi := goqu.Insert(CAGES_TABLE).Cols(
			"id",
			"Carnivore",
			"active",
			"max_capacity",
		).Vals(
			goqu.Vals{
				val.Id,
				val.Carnivore,
				val.Active,
				val.MaxCapacity,
			},
		)

		insertSQL, args, err := gi.ToSQL()
		if err != nil {
			transactionErrors = append(transactionErrors, err)
		}

		insertSQLString := fmt.Sprint(insertSQL, args)
		insertSQLString = strings.Replace(insertSQLString, "[", "", -1)
		insertSQLString = strings.Replace(insertSQLString, "]", "", -1)

		fmt.Println(insertSQLString, args)

		_, err = tx.ExecContext(ctx, insertSQLString)
		if err != nil {
			transactionErrors = append(transactionErrors, err)
		}
	}

	if len(transactionErrors) != 0 {
		return transactionErrors
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
func (store LiteDb) PostCage(rc io.ReadCloser)   {}
func (store LiteDb) DeleteCage(rc io.ReadCloser) {}
