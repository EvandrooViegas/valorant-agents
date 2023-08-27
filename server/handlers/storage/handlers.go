package storage_handler

import (
	"valorant-agents/utils"

	"github.com/gofiber/fiber/v2"
)

func UploadImageHandler(c *fiber.Ctx) error {
	file, err := c.FormFile("avatar")
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{ Status: 500, Message: "Error getting the file", Error: err })
	}
	data, err := SaveImage(c, file)
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{ Status: 500, Message: "Error saving the file", Error: err })
	}
	return c.JSON(fiber.Map{"status": 201, "image": data})
}

func DeleteImageHandler(c *fiber.Ctx) error {
	imageName := c.Params("image_name")
	err := RemoveImage(c, imageName)
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{ Status: 500, Message: "Error removing the file", Error: err })
		return err
	}
	return nil
}
