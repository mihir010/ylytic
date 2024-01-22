package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	At     string             `json:"at" bson:"at"`
	Author string             `json:"author" bson:"author"`
	Like   uint               `json:"like" bson:"like"`
	Reply  uint               `json:"reply" bson:"reply"`
	Text   string             `json:"text" bson:"text"`
}
