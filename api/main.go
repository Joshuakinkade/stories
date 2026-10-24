package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/joshuakinkade/web-journal/auth"
	"github.com/joshuakinkade/web-journal/db"
	"github.com/joshuakinkade/web-journal/users"
	"github.com/joshuakinkade/web-journal/web"
)

func main() {
	godotenv.Load(".env")

	authKey := os.Getenv("AUTH_KEY")
	if len(authKey) < 32 {
		log.Panic("Invalid AUTH_KEY")
	}
	refreshKey := os.Getenv("REFRESH_KEY")
	if len(refreshKey) < 32 {
		log.Panic("Invalid REFRESH_KEY")
	}
	authService := auth.New([]byte(authKey), []byte(refreshKey))

	u := users.NewService(authService, users.NewRepository())
	userHandler := web.NewUserHandlers(u)

	m := web.NewMiddleware(authService)

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

	app.Use(m.AuthMiddleware)

	app.Get("/user-details", userHandler.GetUser)
	app.Patch("/user", userHandler.UpdateUser)

	app.Post("/settings", func(ctx fiber.Ctx) error {
		ctx.WriteString("Hello, World!")
		ctx.Status(200)
		return nil
	})

	app.Listen(":8080")
}
