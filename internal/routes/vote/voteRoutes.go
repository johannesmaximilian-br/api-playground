package voteRoutes

import (
	voteHandler "go-voting/internal/handler/vote"

	"github.com/gofiber/fiber/v2"
)

func SetupVoteRoutes(router fiber.Router) {
	vote := router.Group("/vote")

	vote.Post("/", voteHandler.CreateVoteEntry)
	vote.Get("/", voteHandler.GetChart)
}
