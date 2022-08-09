package controller

import "github.com/gofiber/fiber/v2"

type Auth interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}
