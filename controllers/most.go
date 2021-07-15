package controllers

import (
	"../databases"
	"../models"
	"github.com/gofiber/fiber"
)

func MostUser(c *fiber.Ctx) error {

	var user models.User

	var most []models.User
	databases.DB.Model(&user).Order("points DESC").Limit(5).Find(&most)

	return c.JSON(fiber.Map{

		"mostUsers": most,
	})
}
