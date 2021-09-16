package db

import (
	"context"
	"time"

	"github.com/AdolfoCunquero/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRegister(usr models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("DBTwitter")
	col := db.Collection("user")

	usr.Password, _ = EncryptPassword(usr.Password)

	result, err := col.InsertOne(ctx, usr)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
