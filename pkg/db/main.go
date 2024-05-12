package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DB  *mongo.Database
	ctx = context.TODO()
)

type CollectionIndex struct {
	Collection string
	Index      []mongo.IndexModel
}
