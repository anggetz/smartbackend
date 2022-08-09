package controller

import "github.com/gofiber/fiber/v2"

type GroupLocation interface {
	Get(c *fiber.Ctx) error
}
