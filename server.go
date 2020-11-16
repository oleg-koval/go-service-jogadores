package main

import (
	"log"
	"os"

	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/oleg-koval/go-service-jogadores/controllers"
	"github.com/oleg-koval/go-service-jogadores/seeder"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	if len(connectionString) == 0 {
		connectionString = "mongodb://localhost:27017"
	}

	err := mgm.SetDefaultConfig(nil, "players", options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	seeder.GoDrop()
	seeder.GoFake()

	app := fiber.New()
	// Default config
	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	app.Get("/api/players", controllers.GetAllPlayers)
	app.Get("/api/players/:id", controllers.GetPlayerByID)
	app.Post("/api/players", controllers.CreatePlayer)
	app.Delete("/api/players/:id", controllers.DeletePlayer)

	app.Post("/api/requests/", controllers.CreateRequest)
	app.Get("/api/requests/", controllers.GetAllPoints)
	app.Get("/api/requests/closest", controllers.GetAllCloserRequests)

	port := os.Getenv("PORT")

	app.Listen(port)
}
