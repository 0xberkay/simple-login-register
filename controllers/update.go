package controllers

import (
	"github.com/gofiber/fiber"
	"gopkg.in/ini.v1"
)

func Upgrade(c *fiber.Ctx) error {
	veri, err := ini.Load("settings.ini")
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": veri.Section("").Key("update").String(),
	})

}
