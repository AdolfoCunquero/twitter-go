package db

import (
	"context"
	"time"

	"github.com/AdolfoCunquero/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("DBTwitter")
	col := db.Collection("user")

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
