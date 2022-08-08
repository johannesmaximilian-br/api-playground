package voteHandler

import (
	"go-voting/database"
	"go-voting/internal/model"
	response "go-voting/util"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type VoteRequest struct {
	User   string        `json:"user"`
	Voting [5]model.Vote `json:"voting"`
}

func CreateVoteEntry(c *fiber.Ctx) error {
	response := response.CreateResponse(c)
	db := database.DB

	voteRequest := new(VoteRequest)

	if err := c.BodyParser(voteRequest); err != nil {
		return response.SendResponse(fiber.StatusInternalServerError, "Could not create vote entry", err)
	}

	votes := voteRequest.Voting

	if len(votes) < 5 {
		return response.SendResponse(fiber.StatusInternalServerError, "Could not create vote entry", nil)
	}

	var user model.User
	db.Find(&user, "username = ? ", voteRequest.User)

	if user.ID != uuid.Nil {
		return response.SendResponse(fiber.StatusForbidden, "User already voted", nil)
	}

	user.ID = uuid.New()
	user.Username = voteRequest.User
	err := db.Create(&user).Error

	if err != nil {
		return response.SendResponse(fiber.StatusInternalServerError, "Could not create user", err)
	}

	for key := range votes {
		votes[key].ID = uuid.New()
		votes[key].Score = uint(5 - key)
		votes[key].UserID = user.ID
	}

	err = db.Create(&votes).Error
	if err != nil {
		return response.SendResponse(fiber.StatusInternalServerError, "Could not create vote entry", err)
	}

	return response.SendResponse(fiber.StatusInternalServerError, "Vote entry created", votes)
}

func GetChart(c *fiber.Ctx) error {
	response := response.CreateResponse(c)
	db := database.DB
	var chart []model.Chart

	db.Find(&chart)

	if len(chart) == 0 {
		return response.SendResponse(fiber.StatusNotFound, "No Chart found", nil)
	}

	return response.SendResponse(fiber.StatusNotFound, "Chart found", chart)
}
