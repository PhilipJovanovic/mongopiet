package db

import "go.mongodb.org/mongo-driver/mongo"

func InsertOne(collection string, document interface{}) (*mongo.InsertOneResult, error) {
	if DB == nil {
		return nil, ErrNoDB
	}

	return DB.Collection(collection).InsertOne(ctx, document)
}
