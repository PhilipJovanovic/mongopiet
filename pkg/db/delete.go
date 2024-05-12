package db

import "go.mongodb.org/mongo-driver/mongo"

func DeleteOne(coll string, filter interface{}) (*mongo.DeleteResult, error) {
	return DB.Collection(coll).DeleteOne(ctx, filter)
}
