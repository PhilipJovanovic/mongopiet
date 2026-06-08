package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DB      *mongo.Database
	ctx     = context.TODO()
	ErrNoDB = errors.New("no database connection")
)

type CollectionIndex struct {
	Collection string
	Index      []mongo.IndexModel
}
