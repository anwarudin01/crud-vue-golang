package main

import (
	"github.com/anwarudin01/go-fiber-crud/controllers/bbmcontroller"
	"github.com/anwarudin01/go-fiber-crud/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()
	app.Use(cors.New())
	api := app.Group("/api")
	bbm := api.Group("/bbm")

	

	bbm.Get("/", bbmcontroller.Index)
	bbm.Get("/:id", bbmcontroller.Show)
	bbm.Post("/", bbmcontroller.Create)
	bbm.Put("/:id", bbmcontroller.Update)
	bbm.Delete("/:id", bbmcontroller.Delete)

	

	app.Listen(":8080")
}