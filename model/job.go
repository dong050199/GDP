package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Job struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Company     string             `json:"company,omitempty" bson:"company,omitempty"`
	Salary      Salary             `json:"salary,omitempty" bson:"salary,omitempty"`
}
type Salary struct {
	Gross string `json:"gross,omitempty" bson:"gross,omitempty"`
	Net   string `json:"net,omitempty" bson:"net,omitempty"`
}
