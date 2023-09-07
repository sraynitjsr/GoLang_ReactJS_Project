package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sraynitjsr/server/db"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Welcome to the Home Page!")
}

func GetAllData(c *fiber.Ctx) error {
	allData, err := db.GetAllData()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to read data"})
	}

	return c.JSON(allData)
}

func GetNEByType(c *fiber.Ctx) error {
	neType := c.Query("NEType")
	if neType == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "NEType query parameter is required"})
	}
	filteredData, err := db.GetNEByType(neType)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to filter data"})
	}
	return c.JSON(filteredData)
}

func GetByBuildVersion(c *fiber.Ctx) error {
	buildVersion := c.Query("BuildVersion")
	if buildVersion == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "BuildVersion query parameter is required"})
	}
	filteredData, err := db.GetByBuildVersion(buildVersion)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to filter data"})
	}
	return c.JSON(filteredData)
}

func GetByNETypeAndBuildVersion(c *fiber.Ctx) error {
	neType := c.Query("NEType")
	buildVersion := c.Query("BuildVersion")

	if neType == "" || buildVersion == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Both NEType and BuildVersion query parameters are required"})
	}

	filteredData, err := db.GetByNETypeAndBuildVersion(neType, buildVersion)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to filter data"})
	}
	return c.JSON(filteredData)
}

func GetBySupportedConfig(c *fiber.Ctx) error {
	supportedConfig := c.Query("SupportedConfiguration")
	if supportedConfig == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "SupportedConfiguration query parameter is required"})
	}
	filteredData, err := db.GetBySupportedConfig(supportedConfig)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to filter data"})
	}
	return c.JSON(filteredData)
}
