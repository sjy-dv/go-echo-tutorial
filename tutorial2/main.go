package main

import (
	"go-echo-tutorial/controller"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("env error : %v", err)
	}

	s := controller.Server{}

	s.DBConnection(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	s.AppLauncher(os.Getenv("APP_PORT"))
}
