package api

import (
	"th-release/dcm/utils"

	"github.com/gofiber/fiber/v2"
)

type PasswordDto struct {
	PAssword string `json:"password"`
}

func ApiMiddleware(c *fiber.Ctx) error {
	var dto PasswordDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Next()
}
