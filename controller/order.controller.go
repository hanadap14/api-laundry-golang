package controller

import (
	"github.com/gofiber/fiber"
	"github.com/hanadap14/api-laundry-golang.git/database"
	models "github.com/hanadap14/api-laundry-golang.git/model"
)

func GetOrderAll(c *fiber.Ctx) error {
	var orders []models.Order

	database.DB.Find(&orders)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Get all orders",
		"data":    orders,
	})
}

func GetOrderById(c *fiber.Ctx) error {
	id := c.Params("id")

	var order models.Order

	database.DB.Find(&order, id)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Get order by id",
		"data":    order,
	})
}

func CreateOrder(c *fiber.Ctx) error {
	order := new(models.Order)

	if err := c.BodyParser(order); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	database.DB.Create(&order)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Create order",
		"data":    order,
	})
}

func UpdateOrder(c *fiber.Ctx) error {
	id := c.Params("id")

	order := new(models.Order)

	if err := c.BodyParser(order); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	database.DB.Find(&order, id)
	database.DB.Save(&order)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Update order",
		"data":    order,
	})
}

func DeleteOrder(c *fiber.Ctx) error {
	id := c.Params("id")

	var order models.Order

	database.DB.First(&order, id)
	if order.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Order not found",
		})
	}

	database.DB.Delete(&order)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Delete order",
	})
}
