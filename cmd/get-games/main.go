package main

import (
	"context"
	"net/http"

	"github.com/Davilopesm/cblol-schedule-go/handler"
	"github.com/Davilopesm/cblol-schedule-go/internal/database"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	gameHandler := handler.GameHandler{}

	database := database.TableBasics{}
	database.Client()
	games, err := database.GetAllGames()
	if err != nil {
		panic("Failed getting games from database")
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       gameHandler.HandleGameShow(ctx, games),
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}, nil

}

func main() {
	lambda.Start(Handler)
}
