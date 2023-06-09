package route

import "github.com/gofiber/fiber"

func RouteInit(app *fiber.App) {
	// Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Route Group
	api := app.Group("/api")
	{
		api.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!")
		})
	}
}
