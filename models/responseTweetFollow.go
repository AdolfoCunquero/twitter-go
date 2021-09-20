package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResponseTweetFollow struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	UserId         string             `bson:"userId" json:"userId"`
	UserRelationId string             `bson:"userRelationId" json:"userRelationId"`
	Tweet          Tweet              `bson:"tweet" json:"tweet"`
}
