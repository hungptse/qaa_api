package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
    QuestionId primitive.ObjectID    `bson:"_id" json:"question_id"`
    Name        string `bson:"name" json:"name"`
    Content     string `bson:"content" json:"content"`
    IsOpen     bool   `bson:"is_open" json:"is_open"`
    AccountId  primitive.ObjectID    `bson:"account_id" json:"account_id"`
}



