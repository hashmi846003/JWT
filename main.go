package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hashmi846003/OAUTH-2.0-IMPLEMENTATION/routes"
)

func main() {
	app := fiber.New()

	routes.Setup(app) // A routes package/folder is created with 'Setup' function created.

	app.Listen(":8000")
}
