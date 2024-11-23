package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpdateOne updates a single document in the collection
func UpdateOne(coll string, filter, set, unset interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	if DB == nil {
		return nil, ErrNoDB
	}

	if unset != nil {
		return DB.Collection(coll).UpdateOne(ctx, filter, bson.M{"$set": set, "$unset": unset}, opts...)
	}

	return DB.Collection(coll).UpdateOne(ctx, filter, bson.M{"$set": set}, opts...)
}

// UpdateMany updates multiple documents in the collection
func UpdateMany(coll string, filter, set, unset interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	if DB == nil {
		return nil, ErrNoDB
	}

	if unset != nil {
		return DB.Collection(coll).UpdateMany(ctx, filter, bson.M{"$set": set, "$unset": unset}, opts...)
	}

	return DB.Collection(coll).UpdateMany(ctx, filter, bson.M{"$set": set}, opts...)
}
