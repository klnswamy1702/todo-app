package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Todo struct represents a todo item
type Todo struct {
    ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Title     string             `json:"title,omitempty" bson:"title,omitempty"`
    Completed bool               `json:"completed,omitempty" bson:"completed,omitempty"`
}

