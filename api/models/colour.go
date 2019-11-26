package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Colour stores team colours
type Colour struct {
	gorm.Model        // fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` will be added
	TeamID     int    
	Name       string `gorm:"unique_index:idx_name_code"`
	Hex        string `gorm:"unique_index:idx_name_code"`
}

// AllColours returns all Colours
func (db *DB) AllColours() ([]Colour, error) {
	var colours []Colour
	db.Find(&colours)

	return colours, nil
}

// GetColourByID returns a single colour given an ID (int)
func (db *DB) GetColourByID(id int) (Colour, error) {
	var colour Colour
	db.First(&colour, id)

	return colour, nil
}

// AddColour creates a new Colour
func (db *DB) AddColour(colour Colour) (Colour, error) {
	db.Create(&colour)

	return colour, nil
}
