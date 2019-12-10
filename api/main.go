package main

import (
	"fmt"
	"github.com/cbellee/go-team/models"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

type Env struct {
	db models.Datastore
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func main() {
	dbUserName := getEnv("DB_USER_NAME", "dbadmin@dev-go-team-mysql")
	dbPassword := getEnv("DB_PASSWORD", "M1cr0soft1234567890")
	dbHostName := getEnv("DB_HOST_NAME", "dev-go-team-mysql.mysql.database.azure.com")
	dbName := getEnv("DB_NAME", "goteamdb")

	// connection string
	dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:3606)/%s?allowNativePasswords=true&parseTime=true", dbUserName, dbPassword, dbHostName, dbName)
	
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// DB connection
	db, err := models.NewDB(dbConnectionString)
	if err != nil {
		log.Panic(err)
	}

	// drop/create DB tables
	db.DropTableIfExists(&models.Player{})
	db.AutoMigrate(&models.Player{})
	db.DropTableIfExists(&models.Team{})
	db.AutoMigrate(&models.Team{})
	db.DropTableIfExists(&models.Colour{})
	db.AutoMigrate(&models.Colour{})

	// populate seed Data
	var teams []models.Team = []models.Team{
		models.Team{Name: "the red team", Colour: models.Colour{Name: "red", Hex: "#FF0000", TeamID: 1}, Players: []models.Player{
			models.Player{FirstName: "Chris", LastName: "Bellee", Email: "cbellee@microsoft.com", Phone: "0404636252", TeamID: 1},
			models.Player{FirstName: "Dan", LastName: "Dare", Email: "ddare@gmail.com", Phone: "0405101949", TeamID: 1},
			models.Player{FirstName: "Brian", LastName: "Eno", Email: "beno@moresharkthandark.net", Phone: "0404987654", TeamID: 1},
		}},
		models.Team{Name: "the green team", Colour: models.Colour{Name: "green", Hex: "#00FF00", TeamID: 2}, Players: []models.Player{
			models.Player{FirstName: "John", LastName: "Lambert", Email: "j-lambery@microsoft.com", Phone: "0404689252", TeamID: 2},
			models.Player{FirstName: "Doug", LastName: "Taylor", Email: "d.taylor@home.net", Phone: "0415464732", TeamID: 2},
			models.Player{FirstName: "Kieth", LastName: "Jones", Email: "kjones@gmail.com", Phone: "0404004328", TeamID: 2},
		}},
		models.Team{Name: "the blue team", Colour: models.Colour{Name: "blue", Hex: "#0000FF", TeamID: 3}, Players: []models.Player{
			models.Player{FirstName: "Kelly", LastName: "Armitage", Email: "karma@yahoo.com", Phone: "0406393272", TeamID: 3},
			models.Player{FirstName: "Jason", LastName: "Spaceman", Email: "jspace@space.com.au", Phone: "0404333888", TeamID: 3},
			models.Player{FirstName: "Andy", LastName: "Pandy", Email: "apandy@gmail.com", Phone: "0404845712", TeamID: 3},
		}},
	}
	for _, team := range teams {
		db.Create(&team)
	}

	env := &Env{db}

	e.GET("/players", env.allPlayers)
	e.GET("/players/:id", env.getPlayer)
	e.POST("/players", env.addPlayer)
	e.GET("/teams", env.allTeams)
	e.GET("/teams/:id", env.getTeam)
	e.POST("/teams", env.addTeam)
	e.GET("/colours", env.allColours)
	e.GET("/colours/:id", env.getColour)
	e.POST("/colours", env.addColour)

	e.Logger.Fatal(e.Start(":8080"))
}
