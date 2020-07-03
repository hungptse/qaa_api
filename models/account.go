package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct{
    AccountId primitive.ObjectID `bson:"_id"`
    Email string `bson:"email"`
    Password string `bson:"password"`
    Name string `bson:"name"`
}
