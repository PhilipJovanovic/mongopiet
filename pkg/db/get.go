package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOne[T any](collection string, filter bson.M) (*T, error) {
	var t T

	if err := DB.Collection(collection).FindOne(ctx, filter).Decode(&t); err != nil {
		return nil, err
	}

	return &t, nil
}

func FindOneOpts[T any](collection string, filter bson.M, opts *options.FindOneOptions) (*T, error) {
	var t T

	if err := DB.Collection(collection).FindOne(ctx, filter, opts).Decode(&t); err != nil {
		return nil, err
	}

	return &t, nil
}

func Find[T any](collection string, filter bson.M) (*[]T, error) {
	arr := []T{}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := DB.Collection(collection).Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var a T

		if err := cur.Decode(&a); err != nil {
			return nil, err
		}

		arr = append(arr, a)
	}

	return &arr, nil
}
