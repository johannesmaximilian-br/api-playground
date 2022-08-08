package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID       uuid.UUID `gorm:"type:uuid"`
	Username string
	Votes    []Vote `gorm:"foreignKey:user_id;references:id"`
}

type Song struct {
	gorm.Model

	ID     uuid.UUID `gorm:"type:uuid"`
	Title  string    `json:"title"`
	Artist string    `json:"artist"`
	Votes  []Vote    `gorm:"foreignKey:song_id;references:id"`
}

type Vote struct {
	gorm.Model

	ID     uuid.UUID `gorm:"type:uuid"`
	Score  uint      `json:"score"`
	UserID uuid.UUID `gorm:"type:uuid" json:"user_id"`
	SongID uuid.UUID `gorm:"type:uuid" json:"song_id"`
}

type Chart struct {
	SongID uuid.UUID `gorm:"type:uuid" json:"song_id"`
	Total  uint      `json:"total"`
	Title  string    `json:"title"`
	Artist string    `json:"artist"`
}
