package controller

import "github.com/gofiber/fiber/v2"

type Blok interface {
	Get(c *fiber.Ctx) error
	GetByGroupLocation(c *fiber.Ctx) error
}
