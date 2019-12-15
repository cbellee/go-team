package main

import (
	"github.com/cbellee/go-team/models"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

// HTTP event handlers

// allColours handler
func (env *Env) allColours(c echo.Context) error {
	colours, err := env.db.AllColours()
	if err != nil {
		log.Panic(err)
	}
	if len(colours) > 0 {
		return c.JSON(http.StatusOK, colours)
	} else {
		return c.NoContent(http.StatusNotFound)
	}
}

func (env *Env) getColour(c echo.Context) error {
	var id, err = strconv.Atoi(c.Param("id"))
	colour, err := env.db.GetColourByID(id)
	if err != nil {
		log.Panic(err)
	}
	if colour.ID > 0 {
		return c.JSON(http.StatusOK, colour)
	} else {
		return c.NoContent(http.StatusNotFound)
	}
}

func (env *Env) addColour(c echo.Context) error {
	col := new(models.Colour)
	if err := c.Bind(col); err != nil {
		return err
	}
	colour, err := env.db.AddColour(*col)
	if err != nil {
		log.Panic(err)
	}
	return c.JSON(http.StatusOK, colour)
}