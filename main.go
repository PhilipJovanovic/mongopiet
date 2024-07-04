package mongopiet

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.philip.id/mongopiet/pkg/db"
)

// Creates a new mongodb connection and pings the database
func NewClient(url, database string) error {
	clientOptions := options.Client().ApplyURI(url)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return err
	}

	db.DB = client.Database(database)

	return nil
}

// RegisterIndexes registers indexes for collections.
//
// Example:
//
//	err := mongopiet.RegisterIndexes([]db.CollectionIndex{{
//		Collection: "users",
//		Index: []mongo.IndexModel{
//			{
//				Keys:    bson.D{{Key: "username", Value: 1}},
//				Options: options.Index().SetUnique(true),
//			},
//			{
//				Keys:    bson.D{{Key: "email", Value: 1}},
//				Options: options.Index().SetUnique(true),
//			},
//		},
//	}})
func RegisterIndexes(indexes []db.CollectionIndex) error {
	for _, index := range indexes {
		_, err := db.DB.Collection(index.Collection).Indexes().CreateMany(context.TODO(), index.Index)
		if err != nil {
			return err
		}
	}

	return nil
}
