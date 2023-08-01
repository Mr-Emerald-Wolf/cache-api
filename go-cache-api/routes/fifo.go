package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mr-emerald-wolf/go-cache-api/controllers"
)

func FifoRoutes(incomingRoutes *fiber.App) {
	fifoGroup := incomingRoutes.Group("/fifo")
	fifoGroup.Get("/get", controllers.GetFIFO)
	fifoGroup.Post("/put", controllers.PutFIFO)
	fifoGroup.Get("/size", controllers.Length)
	fifoGroup.Delete("/clear", controllers.Clear)
	fifoGroup.Get("/cache", controllers.GetCacheFIFO)
}
