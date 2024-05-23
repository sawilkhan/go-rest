package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sawilkhan/go-rest/database"
	"github.com/sawilkhan/go-rest/handlers"
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

func generateApp() *fiber.App {
	app := fiber.New()

	//ping to app for health check
	app.Get("/health", func(c *fiber.Ctx) error{
		return c.SendString("OK")
	})

	//create library group and routes
	libGroup := app.Group("/library")
	libGroup.Post("/", handlers.CreateLibrary)
	libGroup.Get("/", handlers.GetLibrary)

	bookGroup := app.Group("/book")
	bookGroup.Post("/", handlers.CreateBook)

	return app
}