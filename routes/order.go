package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/konrad-marzec/fiber-api/database"
	"github.com/konrad-marzec/fiber-api/models"
)

type Order struct {
	ID      uint    `json:"id"`
	Product Product `json:"product"`
	User    User    `json:"user"`
}

func CreateResponseOrder(order models.Order, user models.User, product models.Product) Order {
	return Order{ID: order.ID, User: CreateResponseUser(user), Product: CreateResponseProduct(product)}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := user.FindById(int(order.UserRefer)); err != nil {
		return c.Status(400).JSON("User not found")
	}

	var product models.Product
	if err := product.FindById(int(order.ProductRefer)); err != nil {
		return c.Status(400).JSON("Product not found")
	}

	database.Database.Db.Create(&order)

	return c.Status(201).JSON(CreateResponseOrder(order, user, product))
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}

	database.Database.Db.Joins("User").Joins("Product").Find(&orders)

	response := make([]Order, len(orders))

	for i, v := range orders {
		rp := CreateResponseOrder(v, v.User, v.Product)
		response[i] = rp
	}

	return c.Status(200).JSON(response)
}

func GetUserOrders(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Invalid id")
	}

	var user models.User
	if err := user.FindById(id); err != nil {
		return c.Status(400).JSON("User not found")
	}

	orders := []models.Order{}
	database.Database.Db.Joins("User").Joins("Product").Find(&orders, "orders.user_refer = ?", id)

	response := make([]Order, len(orders))

	for i, v := range orders {
		rp := CreateResponseOrder(v, v.User, v.Product)
		response[i] = rp
	}

	return c.Status(200).JSON(response)
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Invalid id")
	}

	var order models.Order
	if err := order.FindById(id); err != nil {
		return c.Status(400).JSON("Order not found")
	}

	return c.Status(200).JSON(CreateResponseOrder(order, order.User, order.Product))
}

func DeleteOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Invalid id")
	}

	var order models.Order
	if err := order.FindById(id); err != nil {
		return c.Status(400).JSON("Order not found")
	}

	if err := database.Database.Db.Delete(&order); err != nil {
		return c.Status(200).JSON("Order deleted")
	}

	return c.Status(400).JSON("Something went wrong")
}
