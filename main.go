package main

import (
	"encoding/json"
	"fmt"
	"go-voting/database"
	"go-voting/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello Voting!")

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	database.ConnectDB()
	router.SetupRoutes(app)
	app.Listen(":3000")
}
