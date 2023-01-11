package route

import (
	"github.com/gofiber/fiber/v2"

	_ "github.com/aldiansahm7654/go-restapi-fiber/docs"
	swagger "github.com/arsmn/fiber-swagger/v2"
)

// SwaggerRoute func for describe group of API Docs routes.
func SwaggerRoute(a *fiber.App) {
	// Create routes group.
	route := a.Group("/swagger")

	// Routes for GET method:
	route.Get("*", swagger.HandlerDefault) // get one user by ID
}
