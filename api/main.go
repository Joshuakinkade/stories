package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/joshuakinkade/web-journal/db"
	"github.com/joshuakinkade/web-journal/handlers"
)

func main() {
	godotenv.Load(".env")

	userHandler := handlers.UserHandlers{}

	fmt.Println(os.Getenv("DB_PORT"))
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Panic("port not valid", err)
	}

	err = db.Init(db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		DBName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Panic("Could not initialize database pool: ", err)
	}

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Login)

	app.Listen(":8080")
}
