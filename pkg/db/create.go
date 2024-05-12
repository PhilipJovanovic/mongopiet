package db

import "go.mongodb.org/mongo-driver/mongo"

func Create(collection string, document interface{}) (*mongo.InsertOneResult, error) {
	return DB.Collection(collection).InsertOne(ctx, document)
}
