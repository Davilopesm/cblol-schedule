package handler

import (
	"context"

	"github.com/Davilopesm/cblol-schedule-go/model"
	game_view "github.com/Davilopesm/cblol-schedule-go/view/game_view"
)

type GameHandler struct{}

func (h *GameHandler) HandleGameShow(c context.Context) string {
	return renderToString(c, game_view.Show(model.Games))
}
