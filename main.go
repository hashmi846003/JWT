package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hashmi846003/jwt/routes"
    "github.com/hashmi846003/jwt/database"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.DBconn()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
