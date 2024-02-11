package main

import (
	"context"
	"net/http"

	"github.com/Davilopesm/cblol-schedule-go/handler"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	gameHandler := handler.GameHandler{}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       gameHandler.HandleGameShow(ctx),
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}, nil

}

func main() {
	lambda.Start(Handler)
}
