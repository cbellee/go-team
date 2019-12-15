package main

import (
	"github.com/cbellee/go-team/models"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

// HTTP Event Handlers
func (env *Env) allTeams(c echo.Context) error {
	teams, err := env.db.AllTeams()
	if err != nil {
		log.Panic(err)
	}
	if len(teams) > 0 {
		return c.JSON(http.StatusOK, teams)
	} else {
		return c.NoContent(http.StatusNotFound)
	}
}

func (env *Env) getTeam(c echo.Context) error {
	var id, err = strconv.Atoi(c.Param("id"))
	team, err := env.db.GetTeamByID(id)
	if err != nil {
		log.Panic(err)
	}
	if team.ID > 0 {
		return c.JSON(http.StatusOK, team)
	} else {
		return c.NoContent(http.StatusNotFound)
	}
}

func (env *Env) addTeam(c echo.Context) error {
	t := new(models.Team)
	if err := c.Bind(t); err != nil {
		return err
	}
	team, err := env.db.AddTeam(*t)
	if err != nil {
		log.Panic(err)
	}
	return c.JSON(http.StatusOK, team)
}
