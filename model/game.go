package model

import "time"

const (
	PainGaming    = "Pain Gaming"
	INTZ          = "INTZ"
	LOUD          = "LOUD"
	REDCanids     = "RED Canids"
	Fluxo         = "Fluxo"
	FURIA         = "FURIA"
	KaBuMEsports  = "KaBuM! Esports"
	Liberty       = "Liberty"
	LOS           = "LOUD"
	VivoKeydStars = "Vivo Keyd Stars"
)

type Team struct {
	Name  string
	Loses string
	Wins  string
}

type Game struct {
	FirstTeam  Team
	SecondTeam Team
	Time       time.Time
	Week       string
}

var Games = []Game{
	{
		FirstTeam: Team{
			Name:  Liberty,
			Loses: "4",
			Wins:  "2",
		},
		SecondTeam: Team{
			Name:  LOS,
			Loses: "4",
			Wins:  "2",
		},
		Time: time.Now(),
		Week: "4",
	},
	{
		FirstTeam: Team{
			Name:  Fluxo,
			Loses: "1",
			Wins:  "4",
		},
		SecondTeam: Team{
			Name:  PainGaming,
			Loses: "1",
			Wins:  "3",
		},
		Time: time.Now(),
		Week: "4",
	},
	{
		FirstTeam: Team{
			Name:  LOUD,
			Loses: "1",
			Wins:  "4",
		},
		SecondTeam: Team{
			Name:  REDCanids,
			Loses: "1",
			Wins:  "3",
		},
		Time: time.Now(),
		Week: "4",
	},
	{
		FirstTeam: Team{
			Name:  FURIA,
			Loses: "1",
			Wins:  "4",
		},
		SecondTeam: Team{
			Name:  VivoKeydStars,
			Loses: "1",
			Wins:  "3",
		},
		Time: time.Now(),
		Week: "4",
	},
	{
		FirstTeam: Team{
			Name:  KaBuMEsports,
			Loses: "1",
			Wins:  "4",
		},
		SecondTeam: Team{
			Name:  INTZ,
			Loses: "1",
			Wins:  "3",
		},
		Time: time.Now(),
		Week: "4",
	},
	{
		FirstTeam: Team{
			Name:  PainGaming,
			Loses: "1",
			Wins:  "4",
		},
		SecondTeam: Team{
			Name:  VivoKeydStars,
			Loses: "1",
			Wins:  "3",
		},
		Time: time.Now(),
		Week: "4",
	},
}

func FilterGamesByWeek(week string) []Game {
	var filteredGames []Game
	for _, game := range Games {
		if game.Week == week {
			filteredGames = append(filteredGames, game)
		}
	}
	return filteredGames
}
