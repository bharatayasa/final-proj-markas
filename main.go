package main

import (
	"fmt"
	"log"

	"github.com/bharatayasa/final-project/config"
	"github.com/bharatayasa/final-project/router"
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

	router.RouterDatabaseBackup(app)

	log.Fatal(app.Listen(":3000"))
}
