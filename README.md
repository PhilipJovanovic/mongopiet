# mongopiet

Golang MongoDB wrapper

### Readme TODO

## How to init

```go
import (
	"go.philip.id/mongopiet"
)

func main() {
	if err := mongopiet.NewClient(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DATABASE")); err != nil {
		log.Fatal(err)
	}


	err = mongopiet.RegisterIndexes([]db.CollectionIndex{{
		Collection: "users",
		Index: []mongo.IndexModel{
			{
				Keys:    bson.D{{Key: "username", Value: 1}},
				Options: options.Index().SetUnique(true),
			},
		},
	}})
	if err != nil {
		log.Fatal(err)
	}
}
```

## Basic usage

```go
import (
	"go.philip.id/mongopiet/db"
)

func main() {
	newUser := &User{
		ID: primitive.NewObjectID(),
		Name: "SNWZY",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	// Create
	_, err := db.InsertOne("user", newUser)
	if err != nil {
		log.Fatal(err)
	}

	// Find
	user, err := db.FindOne[User]("user", bson.M{"_id": newUser.ID})
	if err != nil {
		log.Fatal(err)
	}

	// Find Many
	users, err := db.Find[User]("user", bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(users)

	// Update
	_, err := db.UpdateOne("user", bson.M{"_id": newUser.ID}, bson.M{"name": "SNWZY1"})
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	_, err := db.DeleteOne("user", bson.M{"_id": newUser.ID})
	if err != nil {
		log.Fatal(err)
	}

	// ... and a few more: InsertMany, CountDocuments, UpdateMany, DeleteMany, Aggregate, BulkWrite
}
```

## Use of document struct (still experimental)

### Primary Field for .Save()

#### Use of ID

```go
type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
```

#### Use of the `primary` flag

```go
type User struct {
	Username  string         	 `bson:"username,primary"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
```

### Create Model

User will generate a collection named `users` in the database at creation of the first database entry

```go
type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	CreatedAt time.Time          `bson:"createdAt"`
}

type UserDoc = db.Document[UserDocument]
type UserDocs = db.ManyDocuments[UserDocument]
```

### Create

```go
func main() {
	newUser := db.NewDoc(&User{
		ID:        primitive.NewObjectID(),
		Name:      "SNWZY",
	})

	_, err := newUser.Create()
	if err != nil {
		log.Fatal(err)
	}
}
```

### FindOne + Update

```go
func main() {
	id, _ := primitive.ObjectIDFromHex("XXXXX")
	user := &UserDoc{}

	_, err = user.FindOne(bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}

	user.Model.Name = "SNWZY1"

	user.Save()
}
```

## TODO

- Add propagation
- Limits etc
