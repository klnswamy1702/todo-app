package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Title       string             `json:"title"`
    Description string             `json:"description"`
    Completed   bool               `json:"completed"`
}
