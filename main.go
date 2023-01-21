package main

import (
	_ "github.com/aldiansahm7654/go-restapi-fiber/config/database/migration"
	_ "github.com/aldiansahm7654/go-restapi-fiber/config/initialize"
	"github.com/aldiansahm7654/go-restapi-fiber/config/route"
	"github.com/aldiansahm7654/go-restapi-fiber/middleware"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	log.Println("Starting Program ...")

	configs := fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	}
	app := fiber.New(configs)
	//middleware
	middleware.FiberMiddleware(app)
	// init route
	route.SwaggerRoute(app)
	route.RouteInit(app)

	app.Listen(":2000")
}
