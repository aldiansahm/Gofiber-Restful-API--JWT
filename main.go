package main

import (
	_ "github.com/aldiansahm7654/go-restapi-fiber/config/database/migration"
	_ "github.com/aldiansahm7654/go-restapi-fiber/config/initialize"
	"github.com/aldiansahm7654/go-restapi-fiber/config/route"
	"github.com/aldiansahm7654/go-restapi-fiber/middleware"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	log.Println("Starting Program ...")
	app := fiber.New()
	//middleware
	middleware.FiberMiddleware(app)
	// init route
	route.SwaggerRoute(app)
	route.RouteInit(app)

	app.Listen(":2000")
}
