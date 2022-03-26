package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/konrad-marzec/fiber-api/database"
	"github.com/konrad-marzec/fiber-api/models"
	"github.com/konrad-marzec/fiber-api/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome api")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)

	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Post("/api/users", routes.CreateUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	app.Get("/api/users/:id/orders", routes.GetUserOrders)

	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Post("/api/products", routes.CreateProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)

	app.Get("/api/orders", routes.GetOrders)
	app.Get("/api/orders/:id", routes.GetOrder)
	app.Post("/api/orders", routes.CreateOrder)
	app.Delete("/api/orders/:id", routes.DeleteOrder)
}

func main() {
	database.ConnectDb()

	database.Database.Db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
