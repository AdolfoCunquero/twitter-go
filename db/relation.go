package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AdolfoCunquero/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const col_name_rel string = "relation"

//const pagination int64 = 20

func InsertRelation(rel models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(database)
	col := db.Collection(col_name_rel)
	_, err := col.InsertOne(ctx, rel)

	if err != nil {
		return false, err
	}

	return true, nil
}

func DeleteRelation(rel models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(database)
	col := db.Collection(col_name_rel)

	_, err := col.DeleteOne(ctx, rel)

	if err != nil {
		return false, err
	}

	return true, nil
}

func SearchRelation(rel models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(database)
	col := db.Collection(col_name_rel)

	filter := bson.M{
		"userId":         rel.UserId,
		"userRelationId": rel.UserRelationId,
	}

	var result models.Relation

	fmt.Println(result)

	err := col.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		return false, err
	}

	return true, nil
}

func ReadAllUsers(ID string, page int64, search string, tipo string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(database)
	col := db.Collection("user")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"firstName": bson.M{
			"$regex": `(?i)` + search,
		},
	}

	cursor, err := col.Find(ctx, query, findOptions)

	if err != nil {
		return results, false
	}

	var exists, include bool

	for cursor.Next(ctx) {
		var usr models.User
		err := cursor.Decode(&usr)
		if err != nil {
			log.Println(err.Error())
			return results, false
		}

		var rel models.Relation
		rel.UserId = ID
		rel.UserRelationId = usr.ID.Hex()

		include = false
		exists, err = SearchRelation(rel)

		if tipo == "new" && !exists {
			include = true
		}

		if tipo == "follow" && exists {
			include = true
		}

		if rel.UserRelationId == ID {
			include = false
		}

		if include == true {
			usr.Password = ""
			usr.Biography = ""
			usr.WebSite = ""
			usr.Location = ""
			usr.Banner = ""
			usr.Email = ""

			results = append(results, &usr)
		}
	}
	err = cursor.Err()

	if err != nil {
		return results, false
	}

	cursor.Close(ctx)
	return results, true
}
