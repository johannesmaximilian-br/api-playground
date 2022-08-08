package userRoutes

import (
	userHandler "go-voting/internal/handler/user"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")

	user.Get("/:userId", userHandler.GetUser)
}
