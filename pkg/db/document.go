package db

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var NilTime = time.Time{}

type Document[T any] struct {
	// model
	Model *T

	// is being used for comparing differences
	initial *T

	// Meta Flags
	IsTimeSeries bool
}

type ManyDocuments[T any] struct {
	// model
	Models []Document[T]

	fix *T

	// Meta Flags
	IsTimeSeries bool
}

type Key struct {
	Key       string
	Primary   bool
	Omitempty bool
}

// Creates a new document and initializes it with the given data.
//
// This Method is needed if you want to use the .save() method
func NewDoc[T any](o *T) *Document[T] {
	c := &Document[T]{
		Model: o,
	}

	copy := *o
	c.initial = &copy

	return c
}

// Returns the name of the collection based on the struct name (mainly used internaly)
func (b *Document[T]) CollectionName() string {
	c := strings.ToLower(reflect.ValueOf(b.Model).Type().Elem().Name()) + "s"
	if b.IsTimeSeries {
		c += "_ts"
	}

	return c
}

// Returns the name of the collection based on the struct name (mainly used internaly)
func (b *ManyDocuments[T]) CollectionName() string {
	c := strings.ToLower(reflect.ValueOf(b.fix).Type().Elem().Name()) + "s"
	if b.IsTimeSeries {
		c += "_ts"
	}

	return c
}

// Sets the document to the current model
func (b *Document[T]) FindOne(filter primitive.M) (*Document[T], error) {
	m, err := FindOne[T](b.CollectionName(), filter)
	if err != nil {
		return nil, err
	}

	if m == nil {
		return nil, errors.New("document not found")
	}

	b.Model = m

	copy := *m
	b.initial = &copy

	return b, nil
}

// Sets the document to the current models
func (b *ManyDocuments[T]) Find(filter primitive.M) (*ManyDocuments[T], error) {
	m, err := Find[T](b.CollectionName(), filter)
	if err != nil {
		return nil, err
	}

	b.Models = make([]Document[T], len(*m))

	for i, v := range *m {
		b.Models[i] = *NewDoc(&v)
	}

	return b, nil
}

// [internally] checkFields checks if the fields are empty and fills them with default values
func (b *Document[T]) checkFields() {
	fields := reflect.ValueOf(b.Model).Elem()

	for j := 0; j < fields.NumField(); j++ {
		key := fields.Type().Field(j).Tag.Get("bson")

		if key == "" {
			key = fields.Type().Field(j).Name
		}

		// Set ID
		if key == "_id" && fields.Field(j).Interface() == primitive.NilObjectID {
			fields.Field(j).Set(reflect.ValueOf(primitive.NewObjectID()))
		}

		// check createdAt + updatedAt timestamps
		if (key == "createdAt" || key == "updatedAt") && fields.Field(j).Interface() == NilTime {
			fields.Field(j).Set(reflect.ValueOf(time.Now()))
		}
	}
}

// Creates a new db entry and initializes the struct
func (b *Document[T]) Create() (*mongo.InsertOneResult, error) {
	b.checkFields()

	res, err := InsertOne(b.CollectionName(), b.Model)
	if err != nil {
		return nil, err
	}

	copy := *b.Model
	b.initial = &copy

	return res, nil
}

// Save updates the document in the database
//
// Finds the document by the `_id` field or any field tagged with primary
// and updates the fields that have changed
//
// Primary tag:
//
//	type User struct {
//		Username  string         	 `bson:"username,primary"`
//		CreatedAt time.Time          `bson:"createdAt"`
//		UpdatedAt time.Time          `bson:"updatedAt"`
//	}
func (b *Document[T]) Save() (*mongo.UpdateResult, error) {
	fields := reflect.ValueOf(b.Model).Elem()
	fType := fields.Type()

	initialFields := reflect.ValueOf(b.initial).Elem()

	// check if somehow the fields are not the same
	if fType != initialFields.Type() {
		return nil, errors.New("structure mismatch. Type from initial fields is not equal to type from current fields")
	}

	filter := bson.M{}

	// * add all into bson.D
	set := bson.M{}
	unset := bson.M{}

	for j := 0; j < fields.NumField(); j++ {
		k := getKey(fType.Field(j))

		iField := initialFields.Field(j)

		// find primary key
		if k.Key == "_id" || k.Primary {
			filter = bson.M{k.Key: iField.Interface()}
		} else if k.Key == "updatedAt" {
			// set updatedAt to current time
			set[k.Key] = time.Now()
		} else {
			// compare all other fields
			field := fields.Field(j)
			deepCompare(k, &field, &iField, set, unset)
		}
	}

	if len(unset) > 0 {
		return UpdateOne(b.CollectionName(), filter, set, unset)
	}

	return UpdateOne(b.CollectionName(), filter, set, nil)
}

// *: check behaviour with array, slices
// *: handle omitempty flags etc. for bson
// *: handle unset
// TODO: check behaviour with pointer comparison (if data at pointer has changed -> maybe Elem())
// TODO: check if we can traverse the pointers too (set nested empty fields at mongo might break <- not breaking since we need to init the struct above in go anyways)
func deepCompare(key *Key, field, iField *reflect.Value, set, unset bson.M) {
	if field.Kind() == reflect.Pointer {
		// unset field if omitempty is set and pointer turned into nil
		if key.Omitempty && field.IsNil() && !iField.IsNil() {
			unset[key.Key] = "1"
		} else if !field.IsNil() /* field.Interface() != iField.Interface() */ {
			// TODO: doesnt work with pointer comparison since the copy made isn a deepcopy (yet?)
			set[key.Key] = field.Interface()
		}
	} else if field.Type().Comparable() || field.Kind() == reflect.Array {
		value := field.Interface()
		if value != iField.Interface() {
			set[key.Key] = value
		}
	} else if field.Kind() == reflect.Struct && iField.Kind() == reflect.Struct {
		nField := reflect.ValueOf(field.Interface())
		nIField := reflect.ValueOf(iField.Interface())

		for i := 0; i < nField.NumField(); i++ {
			newKey := getKey(nField.Type().Field(i))
			newKey.Key = key.Key + "." + newKey.Key

			f := nField.Field(i)
			iF := nIField.Field(i)

			deepCompare(newKey, &f, &iF, set, unset)
		}
	} else if field.Kind() == reflect.Slice {
		// check if slices are empty
		if field.Len() > 0 && iField.Len() > 0 {
			set[key.Key] = field.Interface()
		}
	} else {
		// idk, just set it i guess, could be anything
		set[key.Key] = field.Interface()
	}
}

// returns a key struct with the key and flags
// no pointer is used to prevent stuff like
//
//	t := field.Type().Field(i)
//	key := getKey(&t)
//
// so it looks cleaner ¯\_(ツ)_/¯
//
//	key := getKey(field.Type().Field(i))
func getKey(r reflect.StructField) *Key {
	k := &Key{
		Key: r.Tag.Get("bson"),
	}
	if k.Key == "" {
		k.Key = r.Name
	}

	if strings.Contains(k.Key, ",") {
		splitted := strings.Split(k.Key, ",")
		k.Key = splitted[0]

		for _, s := range splitted {
			if s == "primary" {
				k.Primary = true
			} else if s == "omitempty" {
				k.Omitempty = true
			}
		}
	}

	return k
}
