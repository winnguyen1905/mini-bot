// Package mongodb implements mongodb database functionalities.
package mongodb

import (
	"context"
	"fmt"

	"blitzbot/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateOrder creates a new order.
func (m *mongodb) CreateOrder(ctx context.Context, order *models.Order) error {
	collection := m.db.Collection("orders")

	order.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

// GetOrder returns a order.
func (m *mongodb) GetOrder(ctx context.Context, _id string) (*models.Order, error) {
	order := models.Order{}

	collection := m.db.Collection("orders")

	filter := bson.D{{Key: "_id", Value: _id}}

	err := collection.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		return nil, fmt.Errorf("can not fetch order %v from database: %v", _id, err)
	}

	return &order, nil
}

// UpdateOrder changes of order
func (m *mongodb) UpdateOrder(ctx context.Context, _id string, order *models.Order) error {
	collection := m.db.Collection("orders")

	filter := bson.D{{Key: "_id", Value: _id}}
	update := bson.D{{Key: "$set", Value: order}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
