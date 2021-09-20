package models

type Relation struct {
	UserId         string `bson:"userId" json:"userId"`
	UserRelationId string `bson:"userRelationId" json:"userRelationId"`
}
