package models

import ()

type Player struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Colour struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Team struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Players []Player `json:"players"`
	Colour  Colour   `json:"colour"`
}
