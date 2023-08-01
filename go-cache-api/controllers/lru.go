package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mr-emerald-wolf/go-cache-api/eviction"
	"github.com/mr-emerald-wolf/go-cache-api/structs"
)

var lru = structs.NewCache(3, &eviction.LRUEvictionPolicy{})

func GetLRU(ctx *fiber.Ctx) error {

	var payload structs.AccessCacheRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "err": err.Error()})
	}

	value, found := lru.Get(payload.Key)
	if !found {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": "Cache Miss", "Eviction Policy": "LRU"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Cache Hit", payload.Key: value, "Eviction Policy": "LRU"})
}

func PutLRU(ctx *fiber.Ctx) error {

	var payload structs.UpdateCacheRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "err": err.Error()})
	}

	err := lru.Put(payload.Key, payload.Value)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Cache Updated", payload.Key: payload.Value, "Eviction Policy": "LRU"})
}

func LengthLRU(ctx *fiber.Ctx) error {
	len := lru.Length()
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Cache Length", "Length": len})
}

func ClearLRU(ctx *fiber.Ctx) error {
	lru.Clear()
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Cache Cleared"})
}

func DeleteLRU(ctx *fiber.Ctx) error {
	var payload structs.AccessCacheRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "err": err.Error()})
	}

	lru.Delete(payload.Key)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Item Deleted", "Eviction Policy": "LRU"})
}

func GetCacheLRU(ctx *fiber.Ctx) error {
	queue := lru.Queue
	cache := lru.Items
	if queue.Len() == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": "Cache Empty", "Eviction Policy": "LRU"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Cache", "Least Recently Used": queue.Front(), "Most Recently Used": queue.Back(), "Items": cache})
}
