package models

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ProductID string  `json:"product_id" bson:"product_id"`
	Quantity  int     `json:"quantity" bson:"quantity"`
	Price     float64 `json:"price" bson:"price"`
}

type Order struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID     string             `json:"user_id" bson:"user_id"`
	Items      []OrderItem        `json:"items" bson:"items"`
	TotalPrice float64            `json:"total_price" bson:"total_price"`
	CreatedAt  primitive.DateTime `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt  primitive.DateTime `json:"updated_at" bson:"updated_at,omitempty"`
}

// MarshalJSON ensures correct JSON serialization
func (o Order) MarshalJSON() ([]byte, error) {
	type Alias Order // Prevent recursion
	return json.Marshal(&struct {
		ID         string      `json:"id"`
		UserID     string      `json:"user_id"`
		Items      []OrderItem `json:"items"`
		TotalPrice float64     `json:"total_price"`
		CreatedAt  string      `json:"created_at"`
		UpdatedAt  string      `json:"updated_at"`
	}{
		ID:         o.ID.Hex(),
		UserID:     o.UserID,
		Items:      o.Items,
		TotalPrice: o.TotalPrice,
		CreatedAt:  o.CreatedAt.Time().Format(time.RFC3339),
		UpdatedAt:  o.UpdatedAt.Time().Format(time.RFC3339),
	})
}
