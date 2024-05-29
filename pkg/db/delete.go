package db

import "go.mongodb.org/mongo-driver/mongo"

func DeleteOne(coll string, filter interface{}) (*mongo.DeleteResult, error) {
	if DB == nil {
		return nil, ErrNoDB
	}

	return DB.Collection(coll).DeleteOne(ctx, filter)
}
