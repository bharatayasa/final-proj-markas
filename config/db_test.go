package config_test

import (
	"fmt"
	"testing"

	"github.com/bharatayasa/final-project/config"
	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("env not found")
	}
}

func TestConnection(t *testing.T) {
	Init()
	config.OpenDb()
}
