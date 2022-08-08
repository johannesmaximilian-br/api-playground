package router

import (
	songRoutes "go-voting/internal/routes/song"
	userRoutes "go-voting/internal/routes/user"
	voteRoutes "go-voting/internal/routes/vote"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	voteRoutes.SetupVoteRoutes(api)
	songRoutes.SetupSongRoutes(api)
	userRoutes.SetupUserRoutes(api)
}
