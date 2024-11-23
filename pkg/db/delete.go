package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// deletes a single document in the collection
func DeleteOne(coll string, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	if DB == nil {
		return nil, ErrNoDB
	}

	return DB.Collection(coll).DeleteOne(ctx, filter, opts...)
}

// deletes multiple documents in the collection
func DeleteMany(coll string, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	if DB == nil {
		return nil, ErrNoDB
	}

	return DB.Collection(coll).DeleteMany(ctx, filter, opts...)
}
