package v1

import "github.com/gofiber/fiber/v2"

func HandleGetHealthcheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
	})
}
