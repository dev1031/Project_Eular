package models

type User struct {
	Username string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
	Email    string `bson:"email,omitempty"`
	Score    int    `bson:"score"`
	Solved   []int  `bson:"solved"`
}

//for more info and knowledge about omitempty please visit and read this wonderful tutorial about omitempty
//https://www.sohamkamani.com/golang/omitempty/
