package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/yrk1n/backend-checkout/database"
	"github.com/yrk1n/backend-checkout/handlers"
	"github.com/yrk1n/backend-checkout/models"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App, cartHandler *handlers.CartHandler) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/cart", cartHandler.DisplayItems)

	app.Post("/addItem", cartHandler.AddItem)

	app.Get("/item/:itemId", func(c *fiber.Ctx) error {
		var item models.Item
		if err := database.DB.Db.First(&item, c.Params("itemId")).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "Item not found",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get item",
			})
		}
		return c.JSON(item)
	})

	app.Post("/:itemId/add-vas", cartHandler.AddVasItemToItem)

	// app.Delete("/item/:itemId", cartHandler.RemoveItem)

	app.Post("/reset", cartHandler.ResetCart)

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
