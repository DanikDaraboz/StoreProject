package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"` // MongoDB ObjectID
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Price       float64            `bson:"price" json:"price"`
	Stock       int                `bson:"stock" json:"stock"`
	Category    string             `bson:"category" json:"category"`
	Images      []string           `bson:"images" json:"images"`
	CreatedAt   primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt   primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at"`
}
