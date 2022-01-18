package handler

import (
	"filaments/config"

	"github.com/gofiber/fiber/v2"
)

// GetAllProjects query all Projects
func GetAllProjects(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All Projects",
		"data":    config.Config("MY_ACCOUNT_ID"),
	})
}

// GetProject query Project
func GetProject(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Project found",
		"data":    id,
	})
}

// CreateProject new Project
func CreateProject(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Created Project",
		"data":    "",
	})
}

// DeleteProject delete Project
func DeleteProject(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Project successfully deleted",
		"data":    id,
	})
}
