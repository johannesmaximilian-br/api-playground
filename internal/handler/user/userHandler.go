package userHandler

import (
	"go-voting/database"
	"go-voting/internal/model"
	response "go-voting/util"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserVoteResult struct {
	ID       uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	SongID   uuid.UUID `json:"sing_id"`
	Title    string    `json:"title"`
	Artist   string    `json:"artist"`
	Score    uint      `json:"score"`
}

func GetUser(c *fiber.Ctx) error {
	response := response.CreateResponse(c)
	db := database.DB

	// var user model.User
	var result []UserVoteResult

	id := c.Params("userId")

	db.Debug().Model(&model.User{}).Select("users.id", "users.username", "songs.id as song_id", "songs.title", "songs.artist", "votes.score").Joins("JOIN votes ON votes.user_id=users.id").Joins("JOIN songs ON songs.id=votes.song_id").Where("users.id = ?", id).Order("votes.score DESC").Scan(&result)
	// db.Debug().Preload("Votes").Find(&user, "users.id = ?", id)

	// if user.ID == uuid.Nil {
	// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found!", "data": nil})
	// }

	if len(result) == 0 {
		return response.SendResponse(fiber.StatusNotFound, "User not found!", nil)
	}

	return response.SendResponse(fiber.StatusOK, "User found", result)
}
