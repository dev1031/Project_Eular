package models

type User struct {
	Username string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
	Email    string `bson:"email,omitempty"`
	Score    int    `bson:"score,omitempty"`
	Solved   []int  `bson:"solved,omitempty"`
}
