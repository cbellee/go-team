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
<<<<<<< HEAD
	var player Player
	var team Team
	//var colour Colour
	db.Model(&player).Related(&team)
	//db.Find(&players).Find(&colour)

	return player, nil
=======
	var players []Player
	var colour Colour
	db.Find(&players).Find(&colour)

	return players, nil
>>>>>>> 87ce71de57f6fac2dadffc2996d73276cb9da3a4
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
