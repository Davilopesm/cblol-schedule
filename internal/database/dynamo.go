package database

import (
	"context"
	"log"

	"github.com/Davilopesm/cblol-schedule-go/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

// RUN LOCAL
// func (basics *TableBasics) Client() error {
// 	config, err := config.LoadDefaultConfig(context.TODO(),
// 		config.WithRegion("eu-west-1"),
// 		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")),
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	basics.DynamoDbClient = dynamodb.NewFromConfig(config)
// 	return nil
// }

func (basics *TableBasics) Client() error {
	config, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-west-1"),
	)
	if err != nil {
		return err
	}

	basics.DynamoDbClient = dynamodb.NewFromConfig(config)
	return nil
}

func (basics *TableBasics) AddGames(games []model.Game) error {
	for _, game := range games {
		firstTeam, err := attributevalue.MarshalMap(game.FirstTeam)
		if err != nil {
			panic(err)
		}
		secondTeam, err := attributevalue.MarshalMap(game.SecondTeam)
		if err != nil {
			panic(err)
		}
		_, err = basics.DynamoDbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
			TableName: aws.String("games"), Item: map[string]types.AttributeValue{
				"Week":       &types.AttributeValueMemberS{Value: game.Week},
				"Time":       &types.AttributeValueMemberS{Value: game.Time},
				"FirstTeam":  &types.AttributeValueMemberM{Value: firstTeam},
				"SecondTeam": &types.AttributeValueMemberM{Value: secondTeam},
			},
		})
		if err != nil {
			log.Printf("Couldn't add item to table %v\n", err)
			break
		}
	}
	return nil
}

func (basics *TableBasics) GetAllGames() ([]model.Game, error) {
	out, err := basics.DynamoDbClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("games"),
	})
	if err != nil {
		return nil, err
	}

	var games []model.Game
	for _, item := range out.Items {
		var game model.Game
		if err := attributevalue.UnmarshalMap(item, &game); err != nil {
			log.Printf("Error unmarshalling item: %v\n", err)
			continue
		}
		games = append(games, game)
	}
	return games, nil
}
