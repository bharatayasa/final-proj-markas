package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error, env not found")
	}
}

func AuthMiddleware(c *fiber.Ctx) error {
	Init()

	authHeader := c.Get("Authorization")
	if authHeader != os.Getenv("SECRET_KEY") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	return c.Next()
}
