package main

import (
	"context"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName         = "orders"
	collectionName = "orders"
)

type store struct {
	db *mongo.Client
}

func NewStore(db *mongo.Client) *store {
	return &store{db: db}
}

func (s *store) Create(ctx context.Context, o Order) (primitive.ObjectID, error) {
	col := s.db.Database(dbName).Collection(collectionName)
	r, err := col.InsertOne(ctx, o)
	id := r.InsertedID.(primitive.ObjectID)

	return id, err
}

func (s *store) Get(ctx context.Context, id, customerID string) (*Order, error) {
	col := s.db.Database(dbName).Collection(collectionName)

	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var o Order
	err = col.FindOne(ctx, bson.M{
		"_id":        oID,
		"customerID": customerID,
	}).Decode(&o)

	return &o, err
}

func (s *store) Update(ctx context.Context, ID string, newOrder *pb.Order) error {
	col := s.db.Database(dbName).Collection(collectionName)

	oID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	_, err = col.UpdateOne(
		ctx,
		bson.M{
			"_id": oID,
		},
		bson.M{
			"$set": bson.M{
				"paymentLink": newOrder.PaymentLink,
				"status":      newOrder.Status,
			},
		},
	)

	return err
}
