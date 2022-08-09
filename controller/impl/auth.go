package impl

import (
	"smartservice/controller"
	"smartservice/model"
	services "smartservice/services/impl"

	"github.com/gofiber/fiber/v2"
)

type Auth struct {
}

func NewAuth() controller.Auth {
	return &Auth{}
}

// define the request struct payload
type reqLogin struct {
	Username string
	Password string
}

// login get token of authentication
func (a *Auth) Login(c *fiber.Ctx) error {
	reqDataLogin := reqLogin{}

	err := c.BodyParser(&reqDataLogin)
	if err != nil {
		return err
	}

	token, err := services.NewAuth().Login(reqDataLogin.Username, reqDataLogin.Password)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"data": token,
	})
}

func (a *Auth) Register(c *fiber.Ctx) error {
	reqDataUser := model.User{}

	err := c.BodyParser(&reqDataUser)
	if err != nil {
		return err
	}

	token, err := services.NewAuth().Register(&reqDataUser)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"data": token,
	})
}
