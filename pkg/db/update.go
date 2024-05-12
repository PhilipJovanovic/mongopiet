package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateOne[T any](coll string, filter interface{}, set interface{}) (*mongo.UpdateResult, error) {
	return DB.Collection(coll).UpdateOne(ctx, filter, bson.M{"$set": set})
}
