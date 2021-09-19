package db

import (
	"context"
	"time"

	"github.com/AdolfoCunquero/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModifyRegister(usr models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("DBTwitter")
	col := db.Collection("user")

	register := make(map[string]interface{})

	if len(usr.FirstName) > 0 {
		register["firstName"] = usr.FirstName
	}

	if len(usr.LastName) > 0 {
		register["lastName"] = usr.LastName
	}

	if len(usr.BirthDate.String()) > 0 {
		register["birthDate"] = usr.BirthDate
	}

	if len(usr.Avatar) > 0 {
		register["avatar"] = usr.Avatar
	}

	if len(usr.Banner) > 0 {
		register["banner"] = usr.Banner
	}

	if len(usr.Biography) > 0 {
		register["biography"] = usr.Biography
	}

	if len(usr.Location) > 0 {
		register["location"] = usr.Location
	}

	if len(usr.WebSite) > 0 {
		register["webSite"] = usr.WebSite
	}

	updateString := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{
		"_id": bson.M{"$eq": objID},
	}

	_, err := col.UpdateOne(ctx, filter, updateString)

	if err != nil {
		return false, err
	}

	return true, nil
}
