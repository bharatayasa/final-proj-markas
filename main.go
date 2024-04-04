package main

import (
	"fmt"
	"log"

	"github.com/bharatayasa/mini-project3-markas/config"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("env not found")
	}
}

func main() {
	Init()
	config.OpenDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hallo world")
	})

	log.Fatal(app.Listen(":3000"))
}
