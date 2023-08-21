package storage_handler

import (
	"valorant-agents/utils"

	"github.com/gofiber/fiber/v2"
)

func UploadImageHandler(c *fiber.Ctx) error {
	file, err := c.FormFile("avatar")
	if err != nil {
		return utils.HandleJSONerror(c, fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}
	data, err := SaveImage(c, file)
	if err != nil {
		utils.HandleErr(err)
	}
	return c.JSON(fiber.Map{"status": 201, "image": data})
}

func DeleteImageHandler(c *fiber.Ctx) error {
	imageName := c.Params("image_name")
	err := RemoveImage(c, imageName)
	if err != nil {
		utils.HandleJSONerror(c, fiber.Map{"status": 500, "message": "Error removing image", "data": nil})
		return err
	}
	return nil
}
