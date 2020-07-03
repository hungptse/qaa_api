package api

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"qaa_api/constant"
	"qaa_api/database"
	"qaa_api/helper"
	"qaa_api/models"
	"time"
)

type Account models.Account

func Login(email string, password string) (bool, string)  {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)

	defer cancel()
	account := Account{}
	col := database.GetConnection().Collection(constant.ACCOUNT_TABLE)
	if err := col.FindOne(ctx, bson.M{"email" : email}).Decode(&account); err != nil {
		log.Println(err)
	}
	if helper.ComparePassword(password, account.Password) {
		if token, err := helper.CreateToken(account.AccountId); err == nil {
			log.Println(token)
			return true, token
		}
	}
	return false, "";
}

func Register(name string, email string, password string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)

	defer cancel()
	account := Account{
		AccountId : primitive.NewObjectID(),
		Name : name,
		Email : email,
		Password : helper.HashPassword(password),
	}

	col := database.GetConnection().Collection(constant.ACCOUNT_TABLE)

	_, err := database.GetConnection().Collection(constant.ACCOUNT_TABLE).Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{"email",1}},
		Options: options.Index().SetUnique(true),
	})

	if err != nil {
		log.Println(err)
		return false
	}

	res, err := col.InsertOne(ctx, &account)
	if err != nil {
		log.Println("Email existed")
		return false
	}

	log.Printf("New account created with id: %s", res.InsertedID.(primitive.ObjectID).Hex())
	return true
}
