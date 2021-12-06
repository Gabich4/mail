package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Text      string             `bson:"text"`
	Receivers []string           `bson:"receivers"`
	Status    string             `bson:"status"`
}
