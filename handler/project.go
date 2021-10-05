package handler

import (
	"filaments/database"
	"filaments/model"

	"github.com/gofiber/fiber/v2"
)

// GetAllProjects query all Projects
func GetAllProjects(c *fiber.Ctx) error {
	db := database.DB
	var Projects []model.Project
	db.Find(&Projects)
	return c.JSON(fiber.Map{"status": "success", "message": "All Projects", "data": Projects})
}

// GetProject query Project
func GetProject(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var Project model.Project
	db.Find(&Project, id)
	if Project.ProjectName == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Project found with ID", "data": nil})

	}
	return c.JSON(fiber.Map{"status": "success", "message": "Project found", "data": Project})
}

// CreateProject new Project
func CreateProject(c *fiber.Ctx) error {
	db := database.DB
	Project := new(model.Project)
	if err := c.BodyParser(Project); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create Project", "data": err})
	}
	db.Create(&Project)
	return c.JSON(fiber.Map{"status": "success", "message": "Created Project", "data": Project})
}

// DeleteProject delete Project
func DeleteProject(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var Project model.Project
	db.First(&Project, id)
	if Project.ProjectName == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Project found with ID", "data": nil})

	}
	db.Delete(&Project)
	return c.JSON(fiber.Map{"status": "success", "message": "Project successfully deleted", "data": nil})
}
