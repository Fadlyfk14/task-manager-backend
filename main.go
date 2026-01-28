package main

import (
	"os"

	"beckend/config"
	"beckend/model"
	"beckend/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// ===== DATABASE =====
	config.InitDB()

	config.DB.AutoMigrate(
		&model.User{},
		&model.Task{},
	)

	// ===== CORS =====
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTIONS",
	}))

	// ===== LOGGER =====
	app.Use(logger.New())

	// ===== ROUTES =====
	router.SetupRoutes(app)

	// ===== START SERVER (WAJIB PAKE PORT ENV) =====
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)
}
