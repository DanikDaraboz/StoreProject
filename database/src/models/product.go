package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Price       float64            `bson:"price"`
	Description string             `bson:"description"`
	ImageURL    string             `bson:"imageUrl"`
}
