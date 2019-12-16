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

type TeamPlayer struct {
	ID int
	FirstName string
	LastName string
	Phone string
	Email string
	TeamID int
	TeamName string
	TeamColour string
	TeamColourID int
	TeamColourHex string
}

// AllPlayers returns all Players
func (db *DB) AllPlayers() ([]Player, error) {
	var players []Player
	var colour Colour
	db.Find(&players).Find(&colour)
	return players, nil
}

// AllTeamPlayers returns player & team information
func (db *DB) AllTeamPlayers() ([]TeamPlayer, error) {
	var teamPlayers []TeamPlayer
	db.Table("players").Select("players.id, players.first_name, players.last_name, players.phone, players.email, teams.id as team_id, teams.name as team_name, colours.name as team_colour, colours.id as team_colour_id, colours.hex as team_colour_hex").Joins("join teams on teams.id = players.team_id").Joins("join colours on colours.team_id = teams.id").Scan(&teamPlayers)
	return teamPlayers, nil
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
