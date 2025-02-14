package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CartItem struct {
	ProductID primitive.ObjectID `bson:"product_id" json:"product_id"`
	Quantity  int                `bson:"quantity" json:"quantity"`
	Price     float64            `bson:"price" json:"price"`
}

type Cart struct {
	UserID     primitive.ObjectID `bson:"user_id" json:"user_id"`
	Items      []CartItem         `bson:"items" json:"items"`
	TotalPrice float64            `json:"total_price" bson:"total_price"`
}
