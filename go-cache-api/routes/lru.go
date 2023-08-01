package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mr-emerald-wolf/go-cache-api/controllers"
)

func LruRoutes(incomingRoutes *fiber.App) {
	LruGroup := incomingRoutes.Group("/lru")
	LruGroup.Get("/get", controllers.GetLRU)
	LruGroup.Post("/put", controllers.PutLRU)
	LruGroup.Get("/size", controllers.LengthLRU)
	LruGroup.Delete("/clear", controllers.ClearLRU)
	LruGroup.Get("/cache", controllers.GetCacheLRU)
}
