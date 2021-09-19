package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tweet struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId  string             `bson:"userId" json:"userId,omitempty"`
	Message string             `bson:"message" json:"message,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
}
