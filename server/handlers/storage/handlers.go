package storage_handler

import (
	"fmt"
	"strings"
	"valorant-agents/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("avatar")
	if err != nil {
		utils.HandleErr(err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}
    uniqueId := uuid.New()
    filename := strings.Replace(uniqueId.String(), "-", "", -1)
    fileExt := strings.Split(file.Filename, ".")[1]
    image := fmt.Sprintf("%s.%s", filename, fileExt)
	
	fmt.Println("Saving file")
    err = c.SaveFile(file, fmt.Sprintf("./storage/avatars/%s", image))

    if err != nil {
        utils.HandleErr(err)
        return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
    }

    imageUrl := fmt.Sprintf(" http://127.0.0.1:8080/storage/avatars/%s", image)
    data := map[string]interface{}{
        "name": image,
        "url":  imageUrl,
        "header":    file.Header,
        "size":      file.Size,
    }

    return c.JSON(fiber.Map{"status": 201, "image": data})
}