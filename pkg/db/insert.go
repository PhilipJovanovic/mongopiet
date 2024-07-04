package db

import "go.mongodb.org/mongo-driver/mongo"

// InsertOne inserts a single document into the collection
func InsertOne(collection string, document interface{}) (*mongo.InsertOneResult, error) {
	if DB == nil {
		return nil, ErrNoDB
	}

	return DB.Collection(collection).InsertOne(ctx, document)
}
