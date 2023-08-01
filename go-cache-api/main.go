package main

import (
	"github.com/gofiber/fiber"
	"github.com/mr-emerald-wolf/go-cache-api/initializers"
)

func main() {

	cache := initializers.InitCache()

	app := fiber.New()
}
