package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertOne inserts a single document into the collection
func InsertOne(collection string, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	if DB == nil {
		return nil, ErrNoDB
	}

	return DB.Collection(collection).InsertOne(ctx, document, opts...)
}

// InsertMany inserts multiple documents into the collection
func InsertMany(collection string, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	if DB == nil {
		return nil, ErrNoDB
	}

	return DB.Collection(collection).InsertMany(ctx, documents, opts...)
}
