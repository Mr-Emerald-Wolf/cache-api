package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mr-emerald-wolf/go-cache-api/controllers"
)

func LifoRoutes(incomingRoutes *fiber.App) {
	lifoGroup := incomingRoutes.Group("/lifo")
	lifoGroup.Get("/get", controllers.GetLIFO)
	lifoGroup.Post("/put", controllers.PutLIFO)
	lifoGroup.Get("/size", controllers.LengthLIFO)
	lifoGroup.Delete("/clear", controllers.ClearLIFO)
	lifoGroup.Get("/cache", controllers.GetCacheLIFO)
}
