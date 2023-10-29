package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/yrk1n/backend-checkout/database"
	"github.com/yrk1n/backend-checkout/models"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/test-db", func(c *fiber.Ctx) error {
		var item models.Item
		if err := database.DB.Db.First(&item).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.SendString("Database connected but table is empty.")
			}
			return c.SendString("Error connecting to database.")
		}
		return c.SendString("Database connection successful!")
	})
}
