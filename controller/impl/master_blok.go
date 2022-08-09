package impl

import (
	"smartservice/controller"
	services "smartservice/services/impl"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Blok struct{}

func NewBlok() controller.Blok {
	return &Blok{}
}

func (b *Blok) Get(c *fiber.Ctx) error {
	dataBloks, err := services.NewMasterBlok().Get()
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"data": dataBloks,
	})
}

func (b *Blok) GetByGroupLocation(c *fiber.Ctx) error {
	idGroupLocation, err := strconv.Atoi(c.Query("id_group_location"))
	if err != nil {
		return err
	}

	dataBloks, err := services.NewMasterBlok().GetByGroupLocation(idGroupLocation)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"data": dataBloks,
	})
}
