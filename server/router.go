package server

import (
	"smartservice/controller/impl"

	"github.com/gofiber/fiber/v2"
)

func Router(f *fiber.App) {
	routerAuth(f)
	routerMaster(f)
}

func routerAuth(f *fiber.App) {
	apiAuth := f.Group("/v1")
	apiAuth.Post("/token", impl.NewAuth().Login)
	apiAuth.Post("/register", impl.NewAuth().Register)
}

func routerMaster(f *fiber.App) {
	apiMaster := f.Group("/v1")
	apiMaster.Get("/master-blok/get", impl.NewBlok().Get)
	apiMaster.Get("/master-blok/get-by-group-location", impl.NewBlok().GetByGroupLocation)

	apiMaster.Get("/master-group-location/get", impl.NewGroupLocation().Get)

}
