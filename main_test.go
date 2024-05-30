package mongopiet

import (
	"log"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.philip.id/mongopiet/pkg/m"
)

const (
	CONNECTION_URL = "mongodb://nofy:qNl0tVerb6nf9DGY@db.slabs.io/nofy?authSource=nofy"
	DATABASE       = "nofy"
)

type NN struct {
	Value string `bson:"value"`
	Name  string `bson:"name"`
}

type Nested struct {
	Data  string `bson:"data"`
	Arr   []int  `bson:"arr"`
	NN    NN     `bson:"nn"`
	NNPtr *NN    `bson:"nnPtr"`
}

type Test struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name,primary"`
	Number    int                `bson:"number"`
	Are       []string           `bson:"are"`
	Nested    Nested             `bson:"nested"`
	NestedPtr *Nested            `bson:"nestedPtr,omitempty"`
	NestedArr []Nested           `bson:"nestedArr"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

type TestColl = m.Coll[Test]

func TestMain(t *testing.T) {
	if err := NewClient(CONNECTION_URL, DATABASE); err != nil {
		log.Fatal(err)
	}

	TestFind(t)
}

func TestManual(t *testing.T) {
	b := &TestColl{
		Model: &Test{
			ID:   primitive.NewObjectID(),
			Name: "Test",
		},
	}

	spew.Dump(b)
}

func TestNew(t *testing.T) {
	n := m.NewDoc(&Test{
		ID:        primitive.NewObjectID(),
		Name:      "Test",
		Number:    420,
		Are:       []string{},
		UpdatedAt: time.Now(),
	})

	spew.Dump(n)

	_, err := n.Create()
	if err != nil {
		log.Fatal(err)
	}
}

func TestFind(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("66513865b71f6f4b13d40503")
	if err != nil {
		log.Fatal(err)
	}

	n := &TestColl{}

	_, err = n.FindOne(bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}

	n.Model.Are = append(n.Model.Are, "Test")
	n.Model.Number = 420
	n.Model.Nested = Nested{}
	n.Model.Nested.Data = "Data"

	/* n.Model.NestedPtr = &Nested{
		Data: "Data",
	}

	n.Model.NestedPtr.NNPtr =  &NN{
		Value: "Value2",
		Name:  "Name",
	}
	*/
	n.Save()
}
