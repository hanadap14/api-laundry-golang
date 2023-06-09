package route

import (
	"github.com/gofiber/fiber"
	"github.com/hanadap14/api-laundry-golang.git/controller"
)

func RouteInit(app *fiber.App) {

	// Customer
	app.Get("/api/customer", func(c *fiber.Ctx) {
		controller.GetCustomerAll(c)
	})
	app.Get("/api/customer/:id", func(c *fiber.Ctx) {
		controller.GetCustomerById(c)
	})
	app.Post("/api/customer", func(c *fiber.Ctx) {
		controller.CreateCustomer(c)
	})
	app.Put("/api/customer/:id", func(c *fiber.Ctx) {
		controller.UpdateCustomer(c)
	})
	app.Delete("/api/customer/:id", func(c *fiber.Ctx) {
		controller.DeleteCustomer(c)
	})
}
