package controllers

import (
	"github.com/biskitsx/go-api/db"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	users := new([]db.User)
	db.Db.Find(&users)
	db.Db.Preload("Posts").Find(&users)

	return c.JSON(users)
}
func GetUser(c *fiber.Ctx) error {
	user := new([]db.User)
	db.Db.First(&user, "id = ?", c.Params("id"))
	return c.JSON(user)

}

// func GetUser(c *fiber.Ctx) {

// }

// func GetUser(c *fiber.Ctx) {

// }
