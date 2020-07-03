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

func CreateAnswer(questionId primitive.ObjectID, accountId primitive.ObjectID, content string) (bool, string)  {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)

	defer cancel()

	col := database.GetConnection().Collection(constant.QUESTION_TABLE)
	if res, _ := col.Find(ctx,bson.M{"_id" : questionId}); !res.Next(ctx) {
		return false, ""
	}


	question := models.Answer{
		AnswerId: primitive.NewObjectID(),
		QuestionId: questionId,
		Content : content,
		AccountId: accountId,
	}
	col = database.GetConnection().Collection(constant.ANSWER_TABLE)

	res, err := col.InsertOne(ctx, &question)
	if err != nil {
		log.Println(err)
		return false, ""
	}

	log.Printf("New answer created with id: %s", res.InsertedID.(primitive.ObjectID).Hex())
	return true, res.InsertedID.(primitive.ObjectID).Hex()
}

func ViewAnswerByQuestion(questionId primitive.ObjectID) []models.AnswerWithVote {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)

	defer cancel()
	var vote models.Vote
	var answers []models.AnswerWithVote
	col := database.GetConnection().Collection(constant.ANSWER_TABLE)
	colVoted := database.GetConnection().Collection(constant.VOTE_TABLE)
	res, _ := col.Find(ctx, bson.M{"question_id": questionId});
	for res.Next(ctx) {
		var answer = models.AnswerWithVote{ Voted : 0 }
		if err := res.Decode(&answer); err != nil {
			log.Println(err)
		}
		votedAnswer, _ := colVoted.Find(ctx,bson.M{"id_voted" : answer.AnswerId})
		for votedAnswer.Next(ctx){
			if err := votedAnswer.Decode(&vote); err == nil {
				if vote.IsUpVote {
					answer.Voted++
				} else {
					answer.Voted--
				}
			}
		}
		answers = append(answers,answer)
	}
	if len(answers) == 0 {
		return []models.AnswerWithVote{}
	}
	return answers
}