package main

import (
	"loginApp/database"
	"loginApp/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:8000, https://prakul.work",
	}))

	route.Setup(app)

	app.Listen(":8000")
}
