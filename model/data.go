package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Data struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Company     string             `json:"company,omitempty" bson:"company,omitempty"`
	Salary      string             `json:"salary,omitempty" bson:"salary,omitempty"`
}
