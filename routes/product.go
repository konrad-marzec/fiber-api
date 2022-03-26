package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/konrad-marzec/fiber-api/database"
	"github.com/konrad-marzec/fiber-api/models"
)

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serialNumber"`
}

func CreateResponseProduct(pm models.Product) Product {
	return Product{ID: pm.ID, Name: pm.Name, SerialNumber: pm.SerialNumber}
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Invalid id")
	}

	var product models.Product
	if err := product.FindById(id); err != nil {
		return c.Status(400).JSON("Product not found")
	}

	type UpdateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serialNumber"`
	}

	var up UpdateProduct

	if err := c.BodyParser(&up); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	product.Name = up.Name
	product.SerialNumber = up.SerialNumber

	database.Database.Db.Save(&product)

	return c.Status(200).JSON(CreateResponseProduct(product))
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)

	return c.Status(201).JSON(CreateResponseProduct(product))
}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}

	database.Database.Db.Find(&products)

	response := make([]Product, len(products))

	for i, v := range products {
		rp := CreateResponseProduct(v)
		response[i] = rp
	}

	return c.Status(200).JSON(response)
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Invalid id")
	}

	var product models.Product
	if err := product.FindById(id); err != nil {
		return c.Status(400).JSON("Product not found")
	}

	return c.Status(200).JSON(CreateResponseProduct(product))
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Invalid id")
	}

	var product models.Product
	if err := product.FindById(id); err != nil {
		return c.Status(400).JSON("Product not found")
	}

	if err := database.Database.Db.Delete(&product); err != nil {
		return c.Status(200).JSON("Product deleted")
	}

	return c.Status(400).JSON("Something went wrong")
}
