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

	// Bill
	app.Get("/api/bill", func(c *fiber.Ctx) {
		controller.GetBillAll(c)
	})
	app.Get("/api/bill/:id", func(c *fiber.Ctx) {
		controller.GetBillById(c)
	})
	app.Post("/api/bill", func(c *fiber.Ctx) {
		controller.CreateBill(c)
	})
	app.Put("/api/bill/:id", func(c *fiber.Ctx) {
		controller.UpdateBill(c)
	})
	app.Delete("/api/bill/:id", func(c *fiber.Ctx) {
		controller.DeleteBill(c)
	})

	// Order
	app.Get("/api/order", func(c *fiber.Ctx) {
		controller.GetOrderAll(c)
	})
	app.Get("/api/order/:id", func(c *fiber.Ctx) {
		controller.GetOrderById(c)
	})
	app.Post("/api/order", func(c *fiber.Ctx) {
		controller.CreateOrder(c)
	})
	app.Put("/api/order/:id", func(c *fiber.Ctx) {
		controller.UpdateOrder(c)
	})
	app.Delete("/api/order/:id", func(c *fiber.Ctx) {
		controller.DeleteOrder(c)
	})

	//Pickup
	app.Get("/api/pickup", func(c *fiber.Ctx) {
		controller.GetPickupAll(c)
	})
	app.Get("/api/pickup/:id", func(c *fiber.Ctx) {
		controller.GetPickupById(c)
	})
	app.Post("/api/pickup", func(c *fiber.Ctx) {
		controller.CreatePickup(c)
	})
	app.Put("/api/pickup/:id", func(c *fiber.Ctx) {
		controller.UpdatePickup(c)
	})
	app.Delete("/api/pickup/:id", func(c *fiber.Ctx) {
		controller.DeletePickup(c)
	})
}
