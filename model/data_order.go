package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Attribute_Value struct {
	Product_attribute_id     int    `json:"Product_attribute_id,omitempty" bson:"Product_attribute_id,omitempty"`
	Attribute_Value_id       int    `json:"Attribute_Value_id,omitempty" bson:"Attribute_Value_id,omitempty"`
	Attribute_Value          string `json:"Attribute_Value,omitempty" bson:"Attribute_Value,omitempty"`
	Attribute_Value_Property string `json:"Attribute_Value_Property,omitempty" bson:"Attribute_Value_Property,omitempty"`
}

type Order struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Order_ID        int                `json:"Order_ID,omitempty" bson:"Order_ID,omitempty"`
	User_ID         int                `json:"User_ID,omitempty" bson:"User_ID,omitempty"`
	Phone           string             `json:"Phone,omitempty" bson:"Phone,omitempty"`
	Name            string             `json:"Name,omitempty" bson:"Name,omitempty"`
	Address         string             `json:"Address,omitempty" bson:"Address,omitempty"`
	Product_ID      int                `json:"Product_ID,omitempty" bson:"Product_ID,omitempty"`
	Quantity        int                `json:"Quantity,omitempty" bson:"Quantity,omitempty"`
	Price           float64            `json:"Price,omitempty" bson:"Price,omitempty"`
	Attribute_Value []Attribute_Value  `json:"Attribute_Value,omitempty" bson:"Attribute_Value,omitempty"`
}
