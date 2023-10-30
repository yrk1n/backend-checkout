package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yrk1n/backend-checkout/database"
	"github.com/yrk1n/backend-checkout/handlers"
	"github.com/yrk1n/backend-checkout/models"
	"github.com/yrk1n/backend-checkout/repositories"
	"github.com/yrk1n/backend-checkout/services"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	db := database.DB.Db
	cartRepo := repositories.NewGormCartRepository(db)
	itemRepo := repositories.NewGORMItemRepository(db)
	vasItemRepo := repositories.NewGORMVasItemRepository(db)

	cartService := services.NewCartService(db, cartRepo, itemRepo, vasItemRepo)

	cart, err := cartService.InitializeCart()
	if err != nil {
		log.Fatalf("Failed to initialize cart: %v", err)
	}
	fmt.Printf("Initialized Cart: %+v\n", cart)

	cartHandler := handlers.NewCartHandler(cartService)
	setupRoutes(app, cartHandler)

	var carts []models.Cart
	db.Find(&carts)
	fmt.Printf("Carts: %+v\n", carts)

	app.Listen(":8080")
}
