# mongopiet

Golang MongoDB wrapper

### Readme TODO

## Primary Field for .Save()

# Use of ID

```go
type Test struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Number    int                `bson:"number"`
	Are       []string           `bson:"are"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
```

# Use of the `primary` flag

```go
type Test struct {
	Name      string             `bson:"name,primary"`
	Number    int                `bson:"number"`
	Are       []string           `bson:"are"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
```
