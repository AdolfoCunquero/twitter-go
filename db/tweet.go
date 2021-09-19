package db

import (
	"context"
	"log"
	"time"

	"github.com/AdolfoCunquero/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const database string = "DBTwitter"
const col_name string = "tweet"
const pagination int64 = 20

func InsertTweet(tweet models.Tweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(database)
	col := db.Collection(col_name)

	document := bson.M{
		"userId":  tweet.UserId,
		"message": tweet.Message,
		"date":    tweet.Date,
	}

	result, err := col.InsertOne(ctx, document)

	if err != nil {
		return "", false, err
	}

	objId, _ := result.InsertedID.(primitive.ObjectID)

	return objId.String(), true, nil
}

func ReadTweets(userId string, page int64) ([]*models.Tweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(database)
	col := db.Collection(col_name)

	//var results []*models.Tweet
	var results = make([]*models.Tweet, 0, 25)

	filter := bson.M{
		"userId": userId,
	}

	opt := options.Find()

	opt.SetLimit(pagination)
	opt.SetSort(bson.D{{Key: "date", Value: -1}})
	opt.SetSkip((page - 1) * pagination)

	cursor, err := col.Find(ctx, filter, opt)

	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var document models.Tweet
		err := cursor.Decode(&document)
		if err != nil {
			return results, false
		}
		results = append(results, &document)
	}

	return results, true
}

func DeleteTweet(tweetId string, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(database)
	col := db.Collection(col_name)

	objId, _ := primitive.ObjectIDFromHex(tweetId)

	filter := bson.M{
		"_id":    objId,
		"userId": userId,
	}

	_, err := col.DeleteOne(ctx, filter)

	return err
}
