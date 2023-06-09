package controller

import (
	"github.com/gofiber/fiber"
	"github.com/hanadap14/api-laundry-golang.git/database"
	models "github.com/hanadap14/api-laundry-golang.git/model"
)

func GetBillAll(c *fiber.Ctx) error {
	var bills []models.Bill

	database.DB.Find(&bills)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Get all bills",
		"data":    bills,
	})
}

func GetBillById(c *fiber.Ctx) error {
	id := c.Params("id")

	var bill models.Bill

	database.DB.Find(&bill, id)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Get bill by id",
		"data":    bill,
	})
}

func CreateBill(c *fiber.Ctx) error {
	bill := new(models.Bill)

	if err := c.BodyParser(bill); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	database.DB.Create(&bill)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Create bill",
		"data":    bill,
	})
}

func UpdateBill(c *fiber.Ctx) error {
	id := c.Params("id")

	bill := new(models.Bill)

	if err := c.BodyParser(bill); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	database.DB.Find(&bill, id)
	database.DB.Save(&bill)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Update bill",
		"data":    bill,
	})
}

func DeleteBill(c *fiber.Ctx) error {
	id := c.Params("id")

	var bill models.Bill

	database.DB.First(&bill, id)
	if bill.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Bill not found",
		})
	}

	database.DB.Delete(&bill)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Delete bill",
	})
}
