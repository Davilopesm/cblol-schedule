package model

const (
	PainGaming    = "Pain Gaming"
	INTZ          = "INTZ"
	LOUD          = "LOUD"
	REDCanids     = "RED Canids"
	Fluxo         = "Fluxo"
	FURIA         = "FURIA"
	KaBuMEsports  = "KaBuM! Esports"
	Liberty       = "Liberty"
	LOS           = "LOS"
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
	Time       string
	Week       string
}
