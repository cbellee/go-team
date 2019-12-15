package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Team model
type Team struct {
	gorm.Model        // fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` will be added
	Name       string `gorm:"unique_index:idx_name_code"`
	Colour     Colour
	Players    []Player
}

// AllTeams returns all Teams
func (db *DB) AllTeams() ([]Team, error) {
	var teams []Team
	db.Preload("Players").Preload("Colour").Find(&teams)

	return teams, nil
}

// GetTeamByID returns a single Team
func (db *DB) GetTeamByID(id int) (Team, error) {
	var team Team
	db.Preload("Players").Preload("Colour").Find(&team, id)

	return team, nil
}

// AddTeam creates a new Team
func (db *DB) AddTeam(team Team) (Team, error) {
	db.Create(&team)

	return team, nil
}
