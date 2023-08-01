package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mr-emerald-wolf/go-cache-api/eviction"
	"github.com/mr-emerald-wolf/go-cache-api/structs"
)

var fifo = structs.NewCache(3, &eviction.FIFOEvictionPolicy{})

func GetFIFO(ctx *fiber.Ctx) error {

	var payload structs.AccessCacheRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "err": err.Error()})
	}

	value, found := fifo.Get(payload.Key)
	if !found {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": "Cache Miss", "Eviction Policy": "FIFO"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Cache Hit", payload.Key: value, "Eviction Policy": "FIFO"})
}

func PutFIFO(ctx *fiber.Ctx) error {

	var payload structs.UpdateCacheRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "err": err.Error()})
	}

	err := fifo.Put(payload.Key, payload.Value)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Cache Updated", payload.Key: payload.Value, "Eviction Policy": "FIFO"})
}

func Length(ctx *fiber.Ctx) error {
	len := fifo.Length()
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Cache Length", "Length": len})
}

func Clear(ctx *fiber.Ctx) error {
	fifo.Clear()
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Cache Cleared"})
}

func DeleteFIFO(ctx *fiber.Ctx) error {
	var payload structs.AccessCacheRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "err": err.Error()})
	}

	fifo.Delete(payload.Key)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Item Deleted", "Eviction Policy": "FIFO"})
}

func GetCacheFIFO(ctx *fiber.Ctx) error {
	queue := fifo.Queue
	cache := fifo.Items
	if (queue.Len() == 0) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": "Cache Empty", "Eviction Policy": "FIFO"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Cache", "Front": queue.Front(), "Items": cache})
}
