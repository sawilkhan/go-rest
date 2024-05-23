package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sawilkhan/go-rest/database"
)

func main() {
	//init app
	err := initApp()
	if err != nil{
		panic(err)
	}

	//start app
	defer database.CloseMongoDB()


	app := generateApp()

	// app.Post("/", func(c *fiber.Ctx) error {
	// 	//write a todo to the database
	// 	sampleDoc := bson.M{"name": "sample todo"}
	// 	collection := database.GetCollection("todos")

	// 	nDoc, err := collection.InsertOne(context.TODO(), sampleDoc)
	// 	if err != nil{
	// 		return c.Status(fiber.StatusInternalServerError).SendString("Error inserting todo")
	// 	}

	// 	//send down info about the todo
	// 	return c.JSON(nDoc)
	// })

	port := os.Getenv("PORT")
	app.Listen(":"+port)
}


func initApp() error{
	//setup environment
	err := loadENV()
	if err != nil{
		return err
	}

	//setup database
	err = database.StartMongoDB()
	if err != nil{
		return err
	}
	return nil
}


func loadENV() error{
	goEnv := os.Getenv("GO_ENV")
	if goEnv == ""{
		err := godotenv.Load()
		if err != nil{
			return err
		}
	}
	return nil
}