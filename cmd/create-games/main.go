package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Davilopesm/cblol-schedule-go/internal/database"
	"github.com/Davilopesm/cblol-schedule-go/model"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	database := database.TableBasics{}
	database.Client()
	body := []byte(request.Body)

	var games []model.Game
	err := json.Unmarshal(body, &games)
	if err != nil {
		println(err)
		panic("Failed parsing game from Request")
	}

	err = database.AddGames(games)
	if err != nil {
		panic("Failed adding game to database")
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       "OK",
	}, nil

}

func main() {
	lambda.Start(Handler)
}
