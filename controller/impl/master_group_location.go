package impl

import (
	"smartservice/controller"
	services "smartservice/services/impl"

	"github.com/gofiber/fiber/v2"
)

type GroupLocation struct{}

func NewGroupLocation() controller.GroupLocation {
	return &GroupLocation{}
}

func (b *GroupLocation) Get(c *fiber.Ctx) error {
	dataGroupLocations, err := services.NewMasterGroupLocation().Get()
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"data": dataGroupLocations,
	})
}
