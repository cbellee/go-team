package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Datastore interface
type Datastore interface {
	AllPlayers() ([]Player, error)
	GetPlayerByID(id int) (Player, error)
	AddPlayer(player Player) (Player, error)
	AllTeams() ([]Team, error)
	GetTeamByID(id int) (Team, error)
	AddTeam(team Team) (Team, error)
	AllColours() ([]Colour, error)
	GetColourByID(id int) (Colour, error)
	AddColour(colour Colour) (Colour, error)
}

// DB struct
type DB struct {
	*gorm.DB
}

// NewDB creates a new database connection
func NewDB(dataSourceName string) (*DB, error) {
	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
