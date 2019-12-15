package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Player model
type Player struct {
	gorm.Model // fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` will be added
	TeamID     int
	FirstName  string
	LastName   string
	Phone      string
	Email      string `gorm:"unique_index:idx_name_code"`
}

// AllPlayers returns all Players
func (db *DB) AllPlayers() ([]Player, error) {
	var player Player
	var team Team
	//var colour Colour
	db.Model(&player).Related(&team)
	//db.Find(&players).Find(&colour)

	return player, nil
}

// GetPlayerByID returns a single player
func (db *DB) GetPlayerByID(id int) (Player, error) {
	var player Player
	db.First(&player, id)

	return player, nil
}

// AddPlayer creates a new Player
func (db *DB) AddPlayer(player Player) (Player, error) {
	db.Create(&player)

	return player, nil
}
