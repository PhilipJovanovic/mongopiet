package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.philip.id/mongopiet/pkg/bulk"
)

// writes multiple actions to the database
// example
//
//	models := []bulk.Write{
//		&bulk.Insert{Document: bson.D{{"x", int32(1)}}},
//		&bulk.Update{
//			Filter: bson.D{{"x", int32(1)}},
//			Update: bson.D{{"$set", bson.D{{"x", int32(2)}}}},
//		&bulk.Delete{Filter: bson.D{{"x", int32(2)}}},
//	}
//	BulkWrite("test", models)
func BulkWrite(collection string, models []bulk.Write, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	if DB == nil {
		return nil, ErrNoDB
	}

	return DB.Collection(collection).BulkWrite(ctx, models, opts...)
}
