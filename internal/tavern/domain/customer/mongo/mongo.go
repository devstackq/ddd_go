package mongo

import (
	"context"
	"time"

	"github.com/devstackq/tg_bot_ddd/internal/tavern/domain/customer"
	uuid "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db *mongo.Database
	// customer is used to store customers
	customer *mongo.Collection
}

type mongoCustomer struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (m mongoCustomer) toAggregate() customer.Customer {
	cst := customer.Customer{}
	cst.SetID(m.ID)
	cst.SetName(m.Name)
	return cst
}

func NewFromCustomer(cst customer.Customer) mongoCustomer {
	mngCst := mongoCustomer{}

	mngCst.ID = cst.GetID()
	mngCst.Name = cst.GetName()
	return mngCst
}

func New(ctx context.Context, conn string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn))
	if err != nil {
		return nil, err
	}
	// can use config
	db := client.Database("testDb")
	cstm := db.Collection("customers")
	return &MongoRepository{
		db:       db,
		customer: cstm,
	}, nil
}

func (mr *MongoRepository) Get(id uuid.UUID) (customer.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	// find in mongo db; by user id
	res := mr.customer.FindOne(ctx, bson.M{"id": id})
	var c mongoCustomer
	// decode mongo result -> to tempCStruct
	err := res.Decode(&c)
	if err != nil {
		return customer.Customer{}, err
	}

	return c.toAggregate(), nil
}

func (mr *MongoRepository) Add(c customer.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	cstMng := NewFromCustomer(c)
	_, err := mr.customer.InsertOne(ctx, cstMng)
	if err != nil {
		return err
	}
	// res.InsertedID
	return nil
}

func (mr *MongoRepository) Update(c customer.Customer) error {
	return nil
}
