package controllers

import "github.com/gofiber/fiber"

func GetFIFO(ctx *fiber.Ctx) error {
	
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Message":"Cache Hit",})
}