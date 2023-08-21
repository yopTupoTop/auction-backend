package handler

import (
	"auction_backend/model"
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetHandler(c *fiber.Ctx, db *sql.DB) error {
	var asset model.ContractCreated
	var assets []model.ContractCreated
	rows, err := db.Query("SELECT * FROM placedAssets")
	if err != nil {
		log.Fatal(err)
		c.JSON("An error occured")
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&asset)
		assets = append(assets, asset)
	}

	return c.Render("asset", fiber.Map{
		"PlacedAssets": assets,
	})
}

// TODO: connect with events to write listened events to database
func PostHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("hello")
}
