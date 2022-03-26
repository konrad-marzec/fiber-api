package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/konrad-marzec/fiber-api/database"
	"github.com/konrad-marzec/fiber-api/models"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func CreateResponseUser(um models.User) User {
	return User{ID: um.ID, FirstName: um.FirstName, LastName: um.LastName}
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Invalid id")
	}

	var user models.User
	if err := user.FindById(id); err != nil {
		return c.Status(400).JSON("User not found")
	}

	type UpdateUser struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

	var uu UpdateUser

	if err := c.BodyParser(&uu); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	user.FirstName = uu.FirstName
	user.LastName = uu.LastName

	database.Database.Db.Save(&user)

	return c.Status(200).JSON(CreateResponseUser(user))
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)

	return c.Status(201).JSON(CreateResponseUser(user))
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.Db.Find(&users)

	response := make([]User, len(users))

	for i, v := range users {
		ru := CreateResponseUser(v)
		response[i] = ru
	}

	return c.Status(200).JSON(response)
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Invalid id")
	}

	var user models.User
	if err := user.FindById(id); err != nil {
		return c.Status(400).JSON("User not found")
	}

	return c.Status(200).JSON(CreateResponseUser(user))
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Invalid id")
	}

	var user models.User
	if err := user.FindById(id); err != nil {
		return c.Status(400).JSON("User not found")
	}

	if err := database.Database.Db.Delete(&user); err != nil {
		return c.Status(200).JSON("User deleted")
	}

	return c.Status(400).JSON("Something went wrong")
}
