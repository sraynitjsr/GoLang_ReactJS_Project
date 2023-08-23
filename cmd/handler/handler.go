package handler

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func HelloHandler(c *fiber.Ctx) error {
	response := HelloResponse{
		Message: "Welcome to Students Go Fiber Service",
	}
	return c.JSON(response)
}

func GetAllStudents(c *fiber.Ctx) error {
	targetDir := "STUDENTS"

	var jsonFiles []map[string]interface{}
	err := collectJSONFiles(targetDir, &jsonFiles)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(jsonFiles)
}

func collectJSONFiles(dirPath string, jsonFiles *[]map[string]interface{}) error {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return errors.New("Error reading directory: " + err.Error())
	}

	for _, file := range files {
		filePath := filepath.Join(dirPath, file.Name())

		if file.IsDir() {
			err := collectJSONFiles(filePath, jsonFiles)
			if err != nil {
				return err
			}
		} else if filepath.Ext(file.Name()) == ".json" {
			content, err := os.ReadFile(filePath)
			if err != nil {
				return errors.New("Error reading file: " + err.Error())
			}

			var jsonData map[string]interface{}
			err = json.Unmarshal(content, &jsonData)
			if err != nil {
				return errors.New("Error parsing JSON: " + err.Error())
			}

			*jsonFiles = append(*jsonFiles, jsonData)
		}
	}

	return nil
}

func GetStudentByRollNumber(c *fiber.Ctx) error {
	return nil
}
