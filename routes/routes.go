
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hashmi846003/OAUTH-2.0-IMPLEMENTATION/controllers"
)

func Setup(app *fiber.App) {
	api := app.Group("/user")

	api.Get("/get-user", controllers.User)

	api.Post("/register", controllers.Register)

	api.Post("/login", controllers.Login)

	api.Post("/logout", controllers.Logout)
}
