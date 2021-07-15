package main

import (
	"./databases"
	"./routes"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"

)

func main() {

	databases.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
