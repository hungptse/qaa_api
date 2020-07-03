package api

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"qaa_api/constant"
	"qaa_api/database"
	"qaa_api/models"
	"time"
)

func CreateVote(accountId primitive.ObjectID, id primitive.ObjectID, voteType constant.VoteType) (bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 *  time.Second)

	defer cancel()
	var vote = models.Vote{ VoteId: primitive.NewObjectID(), IdVoted: id, AccountId: accountId, IsUpVote : voteType }

	col := database.GetConnection().Collection(constant.ANSWER_TABLE)
	if res, _ := col.Find(ctx,bson.M{"_id" : id}); !res.Next(ctx) {
		col := database.GetConnection().Collection(constant.QUESTION_TABLE)
		if res, _ := col.Find(ctx,bson.M{"_id" : id}); !res.Next(ctx) {
			return false, ""
		}
	}
	col = database.GetConnection().Collection(constant.VOTE_TABLE)
	if isExisted, _ := col.Find(ctx, bson.M{"account_id": accountId, "id_voted": id}); isExisted.Next(ctx) {
		return false, ""
	}
	res, err := col.InsertOne(ctx, &vote)
	if err != nil {
		log.Println(err)
		return false, ""
	}
	log.Printf("New vote created with id: %s", res.InsertedID.(primitive.ObjectID).Hex())
	return true, res.InsertedID.(primitive.ObjectID).Hex()
}
