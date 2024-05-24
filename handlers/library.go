package handlers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sawilkhan/go-rest/database"
	"github.com/sawilkhan/go-rest/models"
	"go.mongodb.org/mongo-driver/bson"
)


type libraryDTO struct{
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
	Empty []string `json:"no_exists" bson:"books"`
}

func CreateLibrary(c *fiber.Ctx) error {
	nLibrary := libraryDTO{}
	if err:= c.BodyParser(&nLibrary); err != nil{
		return err
	}
	
	nLibrary.Empty = make([]string, 0)

	libraryCollection := database.GetCollection("library")
	nDoc, err := libraryCollection.InsertOne(context.TODO(), nLibrary)
	if err != nil{
		return err
	}

	return c.JSON(fiber.Map{"id": nDoc.InsertedID})
}

func GetLibrary(c *fiber.Ctx) error{
	libraryCollection := database.GetCollection("library")
	cursor, err := libraryCollection.Find(context.TODO(), bson.M{})

	if err != nil{
		return err
	}

	var libraries []models.Library

	if err = cursor.All(context.TODO(), &libraries); err != nil{
		return err
	}

	return c.JSON(libraries)
}

func DeleteLibrary(c *fiber.Ctx) error{	
	libraryCollection := database.GetCollection("library")
	dId := c.Params("id")
	fmt.Println(dId)

	_, err := libraryCollection.DeleteOne(context.TODO(), bson.D{{Key: "id", Value: dId}})

	if err != nil{
		return err
	}

	return c.SendString("Library deleted successfully")
}

