package main

import (
	"airport-vip-lounge/src/database"
	"airport-vip-lounge/src/handlers"
	"airport-vip-lounge/src/middleware"
	"airport-vip-lounge/src/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Airport VIP Lounge Reservation System",
	})

	app.Use(cors.New())
	app.Use(logger.New())

	database.InitDB()
	database.RunMigrations()

	api := app.Group("/api/v1")
	routes.SetupRoutes(api)

	app.Get("/health", handlers.HealthCheck)

	log.Fatal(app.Listen(":8080"))
}
