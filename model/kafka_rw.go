package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type KafkaRw struct {
	Method string             `json:"method,omitempty"`
	Body   []Job              `json:"body,omitempty"`
	Id     primitive.ObjectID `json:"id,omitempty"`
}

type KafkaRwOrder struct {
	Method string  `json:"method,omitempty"`
	Body   []Order `json:"body,omitempty"`
	Id     int     `json:"id,omitempty"`
}
