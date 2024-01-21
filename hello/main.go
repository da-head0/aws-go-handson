package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"sync"
)

var (
	userName      string
	userNameMutex sync.Mutex
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	if request.HTTPMethod == "GET" {
		return events.APIGatewayProxyResponse{
			Body:       "Please enter your name.",
			StatusCode: 200,
			Headers: map[string]string{
				"Content-Type": "text/plain",
			},
		}, nil
	} else if request.HTTPMethod == "POST" {
		userNameMutex.Lock()
		defer userNameMutex.Unlock()

		userName = request.Body

		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Hello, %s!", userName),
			StatusCode: 200,
			Headers: map[string]string{
				"Content-Type": "text/plain",
			},
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       "Invalid request",
		StatusCode: 400,
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
