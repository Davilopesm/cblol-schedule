package model

import "time"

type Team struct {
	Name  string
	Loses string
	Wins  string
}

type Game struct {
	FirstTeam  Team
	SecondTeam Team
	Time       time.Time
}
