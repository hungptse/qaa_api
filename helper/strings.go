package helper

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

func ParseInt(numberInString string) int64 {
	number, _ :=strconv.ParseInt(numberInString,10,64)
	return number
}

func ParseBool(boolInString string) bool  {
	boolean, _ := strconv.ParseBool(boolInString)
	return boolean
}
func ParseObjectID(objectIdString string) primitive.ObjectID  {
	objectId, _ := primitive.ObjectIDFromHex(objectIdString)
	return objectId
}