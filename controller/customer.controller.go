package controller

import (
	"github.com/gofiber/fiber"
	"github.com/hanadap14/api-laundry-golang.git/database"
	models "github.com/hanadap14/api-laundry-golang.git/model"
)

func GetCustomerAll(c *fiber.Ctx) error {
	var customers []models.Customer

	database.DB.Find(&customers)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Get all customers",
		"data":    customers,
	})
}

func GetCustomerById(c *fiber.Ctx) error {
	id := c.Params("id")

	var customer models.Customer

	database.DB.Find(&customer, id)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Get customer by id",
		"data":    customer,
	})
}

func CreateCustomer(c *fiber.Ctx) error {
	customer := new(models.Customer)

	if err := c.BodyParser(customer); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	database.DB.Create(&customer)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Create customer",
		"data":    customer,
	})
}

func UpdateCustomer(c *fiber.Ctx) error {
	id := c.Params("id")

	customer := new(models.Customer)

	if err := c.BodyParser(customer); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	var customerUpdate models.Customer

	database.DB.First(&customerUpdate, id)

	customerUpdate.Name = customer.Name
	customerUpdate.Address = customer.Address

	database.DB.Save(&customerUpdate)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Update customer",
		"data":    customerUpdate,
	})
}

func DeleteCustomer(c *fiber.Ctx) error {
	id := c.Params("id")

	var customer models.Customer

	database.DB.First(&customer, id)
	if customer.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Customer not found",
		})
	}

	database.DB.Delete(&customer)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Delete customer",
	})
}
