package main

import (
	"flag"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/tradmark/cmd"
	"github.com/tradmark/config"
)

func main() {

	log.Println("Loaded configs")

	app := fiber.New()

	port := os.Getenv("PORT")

	flag.StringVar(&port, "port", ":8080", "port number")
	flag.Parse()

	log.Println("Starting server...")

	config.Init()
	cmd.Execute()

	log.Fatal(app.Listen(port))

}