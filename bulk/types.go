package bulk

import "go.mongodb.org/mongo-driver/mongo"

// WriteModel is an interface implemented by models that can be used in a BulkWrite operation. Each WriteModel represents a write.
//
// Original type:
//
//	type WriteModel interface {
//		writeModel()
//	}
//
// Example:
//
//	write := []bulk.Write{
//		&bulk.Insert{Document: bson.D{{"x", int32(1)}}},
//		&bulk.Update{
//			Filter: bson.D{{"x", int32(1)}},
//			Update: bson.D{{"$set", bson.D{{"x", int32(2)}}}},
//		&bulk.Delete{Filter: bson.D{{"x", int32(2)}}},
//	}
//
// Can be used with the following types: bulk.Insert, bulk.Update, bulk.UpdateMany, bulk.Delete,  bulk.DeleteMany, bulk.Replace, bulk.Search, mongo.IndexModel
type Write = mongo.WriteModel

// InsertOneModel is used to insert a single document in a BulkWrite operation.
//
// Original type:
//
//	type InsertOneModel struct {
//		Document interface{}
//	}
type Insert = mongo.InsertOneModel

// UpdateOneModel is used to update at most one document in a BulkWrite operation.
//
// Original type:
//
//	type UpdateOneModel struct {
//		Collation *options.Collation
//		Upsert *bool
//		Filter interface{}
//		Update interface{}
//		ArrayFilters *options.ArrayFilters
//		Hint interface{}
//	}
type Update = mongo.UpdateOneModel

// UpdateManyModel is used to update multiple documents in a BulkWrite operation.
//
// Original type:
//
//	type UpdateManyModel struct {
//		Collation *options.Collation
//		Upsert *bool
//		Filter interface{}
//		Update interface{}
//		ArrayFilters *options.ArrayFilters
//		Hint interface{}
//	}
type UpdateMany = mongo.UpdateManyModel

// DeleteOneModel is used to delete at most one document in a BulkWriteOperation.
//
// Original type:
//
//	type DeleteOneModel struct {
//		Filter interface{}
//		Collation *options.Collation
//		Hint interface{}
//	}
type Delete = mongo.DeleteOneModel

// DeleteManyModel is used to delete multiple documents in a BulkWrite operation.
//
// Original type:
//
//	type DeleteManyModel struct {
//		Filter interface{}
//		Collation *options.Collation
//		Hint interface{}
//	}
type DeleteMany = mongo.DeleteManyModel

// ReplaceOneModel is used to replace at most one document in a BulkWrite operation.
//
// Original type:
//
//	type ReplaceOneModel struct {
//		Collation *options.Collation
//		Upsert *bool
//		Filter interface{}
//		Replacement interface{}
//		Hint interface{}
//	}
type Replace = mongo.ReplaceOneModel

// SearchIndexModel represents a new search index to be created.
//
// Original type:
//
//	type SearchIndexModel struct {
//		// A document describing the definition for the search index. It cannot be nil.
//		// See https://www.mongodb.com/docs/atlas/atlas-search/create-index/ for reference.
//		Definition interface{}
//		// The search index options.
//		Options *options.SearchIndexesOptions
//	}
type Search = mongo.SearchIndexModel

// IndexModel represents a new index to be created.
//
// Original type:
//
//	type IndexModel struct {
//		// A document describing which keys should be used for the index. It cannot be nil. This must be an order-preserving
//		// type such as bson.D. Map types such as bson.M are not valid. See https://www.mongodb.com/docs/manual/indexes/#indexes
//		// for examples of valid documents.
//		Keys interface{}
//		// The options to use to create the index.
//		Options *options.IndexOptions
//	}
type WriteModel mongo.IndexModel
