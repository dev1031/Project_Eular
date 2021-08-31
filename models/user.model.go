package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username,omitempty"`
	Password string             `bson:"password,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Score    int                `bson:"score"`
	Solved   []int              `bson:"solved"`
	Image    string             `bson:"image"`
}

//for more info and knowledge about omitempty please visit and read this wonderful tutorial about omitempty
//https://www.sohamkamani.com/golang/omitempty/
