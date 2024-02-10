package handler

import (
	"time"

	"github.com/Davilopesm/cblol-schedule-go/model"
	game_view "github.com/Davilopesm/cblol-schedule-go/view/game_view"
	"github.com/labstack/echo/v4"
)

type GameHandler struct{}

func (h *GameHandler) HandleGameShow(c echo.Context) error {
	games := []model.Game{
		{
			FirstTeam: model.Team{
				Name:  "Pain Gaming",
				Loses: "0",
				Wins:  "4",
			},
			SecondTeam: model.Team{
				Name:  "INTZ",
				Loses: "4",
				Wins:  "0",
			},
			Time: time.Now(),
		},
		{
			FirstTeam: model.Team{
				Name:  "LOUD",
				Loses: "1",
				Wins:  "3",
			},
			SecondTeam: model.Team{
				Name:  "RED Canids",
				Loses: "3",
				Wins:  "1",
			},
			Time: time.Now(),
		},
	}
	return render(c, game_view.Show(games))
}
