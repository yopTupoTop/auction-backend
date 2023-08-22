package handler

import (
	"auction_backend/model"
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetHandler(c *fiber.Ctx, db *sql.DB) error {
	var contract model.ContractCreated
	var contracts []model.ContractCreated
	rows, err := db.Query("SELECT * FROM ContractCreated")
	if err != nil {
		log.Fatal(err)
		c.JSON("An error occured")
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&contract)
		contracts = append(contracts, contract)
	}

	return c.Render("contract", fiber.Map{
		"PlacedAssets": contracts,
	})
}

// TODO: connect with events to write listened events to database
func PostHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("hello")
}
