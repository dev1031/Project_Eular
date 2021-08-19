package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comments struct {
	Body       string             `bson:"body, omitempty"`
	QuestionId primitive.ObjectID `bson:"questionId"`
	UserId     primitive.ObjectID `bson:"userId"`
}
