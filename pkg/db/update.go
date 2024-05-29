package db

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateOne(coll string, filter, set, unset interface{}) (*mongo.UpdateResult, error) {
	if DB == nil {
		return nil, ErrNoDB
	}

	if unset != nil {
		fmt.Println("with unset")
		return DB.Collection(coll).UpdateOne(ctx, filter, bson.M{"$set": set, "$unset": unset})
	}

	return DB.Collection(coll).UpdateOne(ctx, filter, bson.M{"$set": set})
}
