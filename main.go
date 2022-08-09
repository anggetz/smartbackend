package main

import (
	"log"

	"smartservice/core"
	"smartservice/helpers"
	"smartservice/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	var config helpers.Config

	err := helpers.ReadConfig("./config/config.yaml", &config)
	if err != nil {
		log.Panic(err.Error())
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// run the migration
	err = core.NewDatabase().MigrationModel()
	if err != nil {
		log.Panic(err.Error())
	}

	server.Router(app)

	log.Fatal(app.Listen(config.Host))
}
