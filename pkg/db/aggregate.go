package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func Aggregate[T any](collection string, pipeline interface{}, opts ...*options.AggregateOptions) (*[]T, error) {
	if DB == nil {
		return nil, ErrNoDB
	}

	arr := []T{}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := DB.Collection(collection).Aggregate(ctx, pipeline, opts...)
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
