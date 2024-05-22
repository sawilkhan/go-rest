package main

import (
	"github.com/gofiber/fiber/v2"
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


	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		//write a todo to the database


		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
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
	err := godotenv.Load()

	if err != nil{
		return err
	}

	return nil
}