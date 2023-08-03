package handler

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func GetHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("hello")
}

func PostHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("hello")
}

func PutHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("hello")
}

func DeleteHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("hello")
}
