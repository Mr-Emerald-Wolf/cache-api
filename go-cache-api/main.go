package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mr-emerald-wolf/go-cache-api/routes"
)

func main() {

	app := fiber.New()

	routes.FifoRoutes(app)
	routes.LifoRoutes(app)
	routes.LruRoutes(app)

	err := app.Listen(":8000")
	if err != nil {
		panic(err)
	}
}
