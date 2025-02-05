package models

type Category struct {
	ID   int    `bson:"_id,omitempty"`
	Name string `bson:"name"`
}
