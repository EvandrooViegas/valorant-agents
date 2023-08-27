package storage_handler

import (
	"fmt"
	"mime/multipart"
	"os"
	"strings"
	"valorant-agents/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SaveImage(c *fiber.Ctx, file *multipart.FileHeader) (UploadImageResponse, error) {
	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	image := fmt.Sprintf("%s.%s", filename, fileExt)
	path := fmt.Sprintf("./storage/avatars/%s", image)
	err := c.SaveFile(file, path)
	if err != nil {
		utils.HandleErr(err)

		return UploadImageResponse{}, err
	}
	imageUrl := fmt.Sprintf("http://127.0.0.1:8080/storage/avatars/%s", image)

	return UploadImageResponse{
		Name: image,
		Url:  imageUrl,
		Size: file.Size,
	}, nil
}

func RemoveImage(c *fiber.Ctx, name string) error {
	path := fmt.Sprintf("./storage/avatars/%s", name)
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
