package main

import (
	"log"

	"github.com/go-auth/database"
	"github.com/go-auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()
	app.Use(cors.New)

	routes.Setup(app)

	log.Fatal(app.Listen(":5000"))
}