package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yrk1n/backend-checkout/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":8080")
}
