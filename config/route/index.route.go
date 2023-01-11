package route

import (
	"github.com/aldiansahm7654/go-restapi-fiber/controller"
	auth_controller "github.com/aldiansahm7654/go-restapi-fiber/controller/auth.controller"
	"github.com/aldiansahm7654/go-restapi-fiber/middleware"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	api := r.Group("/api")
	auth := r.Group("/auth")
	book := api.Group("/book")

	auth.Post("/", auth_controller.Login)

	book.Get("/", middleware.JWTProtected(), BookController.Index)
	book.Post("/", middleware.JWTProtected(), BookController.Create)
	book.Put("/:id", BookController.Update)
	book.Delete("/:id", BookController.Delete)
}
