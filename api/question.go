package api

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "qaa_api/constant"
    "qaa_api/database"
    "qaa_api/models"
    "time"
)

func CreateQuestion(name string, content string, isOpen bool, accountId primitive.ObjectID) (bool, string)  {
    ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)

    defer cancel()
    question := models.Question{
        QuestionId: primitive.NewObjectID(),
        Name: name,
        Content : content,
        IsOpen:isOpen,
        AccountId: accountId,
    }
    col := database.GetConnection().Collection(constant.QUESTION_TABLE)

    res, err := col.InsertOne(ctx, &question)
    if err != nil {
        log.Println(err)
        return false, ""
    }

    log.Printf("New question created with id: %s", res.InsertedID.(primitive.ObjectID).Hex())
    return true, res.InsertedID.(primitive.ObjectID).Hex()
}

func ViewQuestion(page int64, size int64, isOpen bool) []models.Question {
    ctx, cancel := context.WithTimeout(context.Background(), 10  * time.Second)
    defer cancel()
    var questions []models.Question
    var skip int64
    skip = (page - 1) * size
    col := database.GetConnection().Collection(constant.QUESTION_TABLE)

    if res, err := col.Find(ctx, bson.M{"is_open": isOpen}, &options.FindOptions{
        Limit: &size,
        Skip:  &skip,
    }); err != nil {
        log.Println(err)
    } else{
        for res.Next(ctx) {
            var question models.Question
            if err := res.Decode(&question); err != nil {
                log.Println(err)
            }
            questions = append(questions,question)
        }
        log.Println(questions)
    }
    if len(questions) > 0 {
        return questions
    }
    return []models.Question{}
}
