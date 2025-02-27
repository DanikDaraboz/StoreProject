package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	CreatedAt   primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt   primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at"`
}
