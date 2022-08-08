package songHandler

import (
	"go-voting/database"
	"go-voting/internal/model"
	response "go-voting/util"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type SongSearchResult struct {
	ID     uuid.UUID `json:"song_id"`
	Artist string    `json:"artist"`
	Title  string    `json:"title"`
}

func CreateSong(c *fiber.Ctx) error {
	response := response.CreateResponse(c)
	db := database.DB
	song := new(model.Song)
	err := c.BodyParser(song)

	if err != nil {
		return response.SendResponse(fiber.StatusInternalServerError, "Could not create song", err)
	}

	song.ID = uuid.New()
	err = db.Create(&song).Error
	if err != nil {
		return response.SendResponse(fiber.StatusInternalServerError, "Could not create song", err)
	}

	return response.SendResponse(fiber.StatusCreated, "Song created!", song)
}

func GetSong(c *fiber.Ctx) error {
	response := response.CreateResponse(c)
	db := database.DB
	var song model.Song
	id := c.Params("songId")

	db.Preload("Votes").Find(&song, "id = ?", id)

	if song.ID == uuid.Nil {
		return response.SendResponse(fiber.StatusNotFound, "Song not found", nil)
	}

	return response.SendResponse(fiber.StatusOK, "Song found", song)
}

func SearchSong(c *fiber.Ctx) error {
	response := response.CreateResponse(c)
	db := database.DB
	var searchResult []SongSearchResult
	query := c.Query("q")

	// Simple similarity match.
	// db.Debug().Model(&model.Song{}).Select("id", "artist", "title").Where("SIMILARITY(artist,?) > 0.2", query).Or("SIMILARITY(title,?) > 0.1", query).Scan(&searchResult)

	// Sorting to metaphone match.
	db.Debug().Model(&model.Song{}).Select("id", "artist", "title").Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "SIMILARITY(METAPHONE(artist,10),METAPHONE(?,10)) DESC",
			Vars:               []interface{}{query},
			WithoutParentheses: true},
	}).Limit(5).Scan(&searchResult)

	if len(searchResult) == 0 {
		return response.SendResponse(fiber.StatusNotFound, "Search has no result!", searchResult)
	}

	return response.SendResponse(fiber.StatusOK, "Search returns matching results.", searchResult)
}
