package main

import (
	"log"
	"project/database"
	"project/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors" // Import CORS middleware
)

func main() {
	// Connect to the database
	database.Connect()

	// Initialize Fiber app
	app := fiber.New()

	// Enable CORS middleware to allow requests from frontend (e.g., React on localhost:3000)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // Replace with the URL of your frontend
		AllowMethods: "GET,POST,PUT,DELETE",   // Allowed methods
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Routes
	app.Post("/submit", handlers.SubmitFormHandler)

	// Start server
	log.Fatal(app.Listen(":8000"))
}

