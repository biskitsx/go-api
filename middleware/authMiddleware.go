package middleware

import (
	"fmt"

	tokenutils "github.com/biskitsx/go-api/utils"
	"github.com/gofiber/fiber/v2"
)

func VerifyUser(c *fiber.Ctx) error {
	token := c.Cookies("access_token")

	if token == "" {
		c.SendStatus(400)
		return c.JSON(fiber.Map{"msg": "user not authenticated"})
	}
	payload, err := tokenutils.VerifyToken(token)
	if err != nil {
		c.SendStatus(400)
		return c.JSON(fiber.Map{"msg": "user not authenticated"})
	}

	id, ok := payload["id"].(float64)
	if !ok {
		c.SendStatus(400)
		return c.JSON(fiber.Map{"msg": "invalid user ID"})
	}

	userId := uint(id)
	fmt.Println(userId)
	c.Locals("userId", userId)
	return c.Next()
}
