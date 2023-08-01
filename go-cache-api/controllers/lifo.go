package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mr-emerald-wolf/go-cache-api/eviction"
	"github.com/mr-emerald-wolf/go-cache-api/structs"
)

var lifo = structs.NewCache(3, &eviction.LIFOEvictionPolicy{})

func GetLIFO(ctx *fiber.Ctx) error {

	var payload structs.AccessCacheRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "err": err.Error()})
	}

	value, found := lifo.Get(payload.Key)
	if !found {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": "Cache Miss", "Eviction Policy": "LIFO"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Cache Hit", payload.Key: value, "Eviction Policy": "LIFO"})
}

func PutLIFO(ctx *fiber.Ctx) error {

	var payload structs.UpdateCacheRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "err": err.Error()})
	}

	err := lifo.Put(payload.Key, payload.Value)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Cache Updated", payload.Key: payload.Value, "Eviction Policy": "LIFO"})
}

func LengthLIFO(ctx *fiber.Ctx) error {
	len := lifo.Length()
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Cache Length", "Length": len})
}

func ClearLIFO(ctx *fiber.Ctx) error {
	lifo.Clear()
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Cache Cleared"})
}

func DeleteLIFO(ctx *fiber.Ctx) error {
	var payload structs.AccessCacheRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "err": err.Error()})
	}

	lifo.Delete(payload.Key)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Item Deleted", "Eviction Policy": "LIFO"})
}

func GetCacheLIFO(ctx *fiber.Ctx) error {
	queue := lifo.Queue
	cache := lifo.Items
	if (queue.Len() == 0) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": "Cache Empty", "Eviction Policy": "LIFO"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Cache", "Front": queue.Front(), "Items": cache})
}
