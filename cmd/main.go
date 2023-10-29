package main

import (
	"github.com/yrk1n/backend-checkout/database"
)

func main() {
	database.ConnectDb()

	setupRoutes(app)
	app.Listen(":8080")
}
