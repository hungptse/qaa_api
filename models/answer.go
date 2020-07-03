package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Answer struct {
	AnswerId primitive.ObjectID `bson:"_id" json:"_id"`
	QuestionId primitive.ObjectID `bson:"question_id" json:"question_id"`
	AccountId primitive.ObjectID `bson:"account_id" json:"account_id"`
	Content string `bson:"content" json:"content"`
}

type AnswerWithVote struct {
	AnswerId primitive.ObjectID `bson:"_id" json:"_id"`
	QuestionId primitive.ObjectID `bson:"question_id" json:"question_id"`
	AccountId primitive.ObjectID `bson:"account_id" json:"account_id"`
	Content string `bson:"content" json:"content"`
	Voted int `json:"voted"`
}