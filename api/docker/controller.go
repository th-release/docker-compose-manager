package docker

import (
	"strings"
	"th-release/dcm/utils"

	"github.com/gofiber/fiber/v2"
)

func Insert(c *fiber.Ctx) error {
	var dto InsertDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	config := utils.GetConfig()

	compose, err := LoadDockerCompose(config.DockerPath)

	if err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: "인식 불가",
			Data:    nil,
		})
	}

	AddService(compose, dto.Service.Name, &dto.Service.Value)

	AddVolume(compose, dto.Volume.Name, &dto.Volume.Value)

	AddNetwork(compose, dto.Network.Name, &dto.Network.Value)

	for _, v := range compose.Services {
		v.Environment["AUTH_KEY"] = dto.Service.Value.Environment["AUTH_KEY"]
	}

	err = SaveDockerCompose(compose, config.DockerPath)
	if err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	RestartDockerCompose(config.DockerPath)

	return c.Status(200).JSON(utils.BasicResponse{
		Success: true,
		Message: "Service added successfully",
		Data:    nil,
	})
}

func Delete(c *fiber.Ctx) error {
	var dto DeleteDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	config := utils.GetConfig()

	compose, err := LoadDockerCompose(config.DockerPath)

	if err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: "인식 불가",
			Data:    nil,
		})
	}

	for _, v := range compose.Services[dto.Name].Volumes {
		DeleteVolume(compose, strings.Split(v, ":")[0])
	}

	DeleteService(compose, dto.Name)

	err = SaveDockerCompose(compose, config.DockerPath)
	if err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	RestartDockerCompose(config.DockerPath)

	return c.Status(200).JSON(utils.BasicResponse{
		Success: true,
		Message: "Service added successfully",
		Data:    nil,
	})
}
