package main

import (
	"react-go-mysql-jwt-authentication/database"
	"react-go-mysql-jwt-authentication/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}
