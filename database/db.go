package database

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "os"
    "qaa_api/constant"
)
func GetConnection() *mongo.Database {
    ctx := context.Background()
    if client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv(constant.DB_URL))); err != nil{
        log.Println(err)
    } else {
        if err := client.Connect(ctx); err != nil {
            log.Println(err)
        }

        return client.Database(constant.DATABASE_NAME);
    }
    return nil
}
