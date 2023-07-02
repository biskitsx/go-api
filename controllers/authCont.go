package controllers

import (
	"github.com/biskitsx/go-api/db"
	tokenutils "github.com/biskitsx/go-api/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	user := new(UserRegister)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	// check user existed
	userExisted := new(db.User)
	db.Db.First(&userExisted, "username = ?", user.Username)
	if userExisted.ID > 0 {
		return c.JSON(fiber.Map{"err": "user already registered"})
	}

	// hash
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	userExisted.Username = user.Username
	userExisted.Password = string(hash)

	// create user
	db.Db.Create(&userExisted)
	return c.JSON(userExisted)
}
func Login(c *fiber.Ctx) error {
	user := new(UserRegister)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	// check user existed
	userExisted := new(db.User)
	db.Db.First(&userExisted, "username = ?", user.Username)
	if userExisted.ID == 0 {
		return c.JSON(fiber.Map{"err": "user does not existed"})
	}

	// hash
	err := bcrypt.CompareHashAndPassword([]byte(userExisted.Password), []byte(user.Password))
	if err != nil {
		return c.JSON(fiber.Map{"err": "wrong password"})
	}
	token, _ := tokenutils.CreateToken(userExisted.ID)
	cookieToken := new(fiber.Cookie)
	cookieToken.Name = "access_token"
	cookieToken.Value = token

	c.Cookie(cookieToken)
	// create user
	return c.JSON(fiber.Map{
		"user": userExisted,
	})
}
