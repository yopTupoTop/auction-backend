package main

import (
	"auction_backend/event"
	"auction_backend/handler"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	connStr := "postgresql://user:yoptupotop@localhost/auction?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	go event.SubscribeToEvents()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return handler.GetHandler(c, db)
	})
	app.Post("/", func(c *fiber.Ctx) error {
		return handler.PostHandler(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
