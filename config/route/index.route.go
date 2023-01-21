package route

import (
	auth_controller "github.com/aldiansahm7654/go-restapi-fiber/controller/auth.controller"
	"github.com/aldiansahm7654/go-restapi-fiber/controller/book.controller"
	"github.com/aldiansahm7654/go-restapi-fiber/middleware"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	api := r.Group("/api")
	auth := api.Group("/auth")
	book := api.Group("/book")

	auth.Post("/", auth_controller.Login)

	book.Get("/", book_controller.Index)
	book.Put("/:id", book_controller.Update)
	book.Delete("/:id", book_controller.Delete)
	//with jwt protected
	book.Post("/", middleware.JWTProtected(), book_controller.Create)
}
