package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"qaa_api/constant"
)

type Vote struct {
	VoteId primitive.ObjectID `bson:"_id" json:"_id"`
	AccountId primitive.ObjectID `bson:"account_id" json:"account_id"`
	IdVoted primitive.ObjectID `bson:"id_voted" json:"id_voted"`
	IsUpVote constant.VoteType `bson:"is_upvote" json:"is_upvote"`
}

