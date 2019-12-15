package main

import (
	"github.com/cbellee/go-team/models"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

// HTTP Event Handlers
func (env *Env) allPlayers(c echo.Context) error {
	players, err := env.db.AllPlayers()
	if err != nil {
		log.Panic(err)
	}
	if len(players) > 0 {
		return c.JSON(http.StatusOK, players)
	} else {
		return c.NoContent(http.StatusNotFound)
	}
}

func (env *Env) getPlayer(c echo.Context) error {
	var id, err = strconv.Atoi(c.Param("id"))
	player, err := env.db.GetPlayerByID(id)
	if err != nil {
		log.Panic(err)
	}
	if player.ID <= 0 {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, player)
}

func (env *Env) addPlayer(c echo.Context) error {
	p := new(models.Player)
	if err := c.Bind(p); err != nil {
		return err
	}
	player, err := env.db.AddPlayer(*p)
	if err != nil {
		log.Panic(err)
	}
	return c.JSON(http.StatusOK, player)
}
