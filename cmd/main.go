package main

import (
	"fmt"

	"github.com/Davilopesm/cblol-schedule-go/handler"
	"github.com/labstack/echo/v4"
)

type DB struct{}

func main() {
	app := echo.New()

	gameHandler := handler.GameHandler{}
	app.GET("/games", gameHandler.HandleGameShow)
	app.Start(":3000")
	fmt.Println("Working")
}
