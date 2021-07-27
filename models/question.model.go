package models

type Question struct {
	Title  string `bson:"title,omitempty"`
	Body   string `bson:"body,omitempty"`
	Answer string `bson:"answer,omitempty"`
}
