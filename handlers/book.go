package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sawilkhan/go-rest/database"
	"github.com/sawilkhan/go-rest/models"
	"go.mongodb.org/mongo-driver/bson"
)

type bookDTO struct {
	Title     string `json:"title" bson:"title"`
	Author    string `json:"author" bson:"author"`
	ISBN      string `json:"isbn" bson:"isbn"`
	LibraryId string `json:"libraryId" bson:"libraryId"`
}

func CreateBook(c *fiber.Ctx) error{
	createData := bookDTO{}

	if err := c.BodyParser(&createData); err != nil{
		return err
	}

	//get the collection reference
	libraryCollection := database.GetCollection("library")

	//get the filter
	filter := bson.D{{Key: "_id", Value: createData.LibraryId}}
	nBookData := models.Book{
		Title : createData.Title,
		Author: createData.Author,
		ISBN: createData.ISBN,
	}
	updatePayload := bson.D{{Key: "$push", Value: bson.D{{Key: "books", Value: nBookData}}}}

	//update the library collection\
	_, err := libraryCollection.UpdateOne(context.TODO(), filter, updatePayload)

	if err != nil{
		return err
	}

	return c.SendString("Book created successfully")
}
