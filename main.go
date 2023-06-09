package main

import (
	"github.com/gofiber/fiber"
	"github.com/hanadap14/api-laundry-golang.git/database"
	"github.com/hanadap14/api-laundry-golang.git/route"
)

func main() {
	database.ConnectDatabase()

	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":3000")
}
