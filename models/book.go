package models

type Book struct {
	ID     string `json:"id" bson:"id"`
	Title  string `json:"title" bson:"title"`
	Author string `json:"author" bson:"author"`
	ISBN   string `json:"isbn" bson:"isbn"`
}