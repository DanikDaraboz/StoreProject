package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Username string             `bson:"username" json:"username"`
    Email    string             `bson:"email" json:"email"`
    Password string             `bson:"password" json:"password"`
    Role     string             `bson:"role" json:"role"` 			// "user" or "admin"
    Age      int                `bson:"age,omitempty" json:"age,omitempty"`
    Phone    string             `bson:"phone,omitempty" json:"phone,omitempty"`
    Address  string             `bson:"address,omitempty" json:"address,omitempty"`
}
