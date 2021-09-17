package db

import (
	"context"
	"log"
	"time"

	"github.com/AdolfoCunquero/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	db := MongoCN.Database("DBTwitter")
	col := db.Collection("user")
	defer cancel()

	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		log.Printf("Registro no encontrado %s\n", err.Error())
		return profile, err
	}

	return profile, nil
}
