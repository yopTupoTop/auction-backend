package handler

import (
	"auction_backend/event"
	"auction_backend/model"
	"fmt"

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
		"ContractCreated": contracts,
	})
}

func PostHandler(c *fiber.Ctx, db *sql.DB) error {
	newEvents := event.ContractsCreated

	if err := c.BodyParser(&newEvents); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}
	fmt.Printf("%v", newEvents)

	for newEvent := 0; newEvent < len(newEvents); newEvent++ {
		_, err := db.Exec("INSERT into ContractCreated VALUES ($1)")
		if err != nil {
			log.Fatalf("An error occured while executing query: %v", err)
		}

		newEvents[newEvent] = newEvents[len(newEvents)-1]
		newEvents = newEvents[:len(newEvents)-1]
	}

	return c.Redirect("/")
}
