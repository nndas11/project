package main

import (
	controller "proj/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/", "./views")
	// stackController := controller.NewStackController(10)
	stackController := controller.StackController{}

	app.Post("/push", stackController.Push)
	app.Delete("/pop", stackController.Pop)
	app.Get("/top", stackController.Top)
	app.Get("/display", stackController.Display)

	app.Post("/declare", stackController.Declare)

	app.Listen(":3000")
}
