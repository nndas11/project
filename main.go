package main

import (
	controller "proj/controllers"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Static("/", "./views")
	stackController := controller.NewStackController(10)

	app.Post("/push", stackController.Push)
	app.Post("/pop", stackController.Pop)
	app.Get("/top", stackController.Top)
	app.Get("/display", stackController.Display)
	app.Listen(":3000")
}
