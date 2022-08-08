package songRoutes

import (
	songHandler "go-voting/internal/handler/song"

	"github.com/gofiber/fiber/v2"
)

func SetupSongRoutes(router fiber.Router) {
	song := router.Group("/song")

	song.Post("/", songHandler.CreateSong)
	song.Get("/search", songHandler.SearchSong)
	song.Get("/:songId", songHandler.GetSong)
}
