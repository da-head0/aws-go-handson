package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// AWS lambda의 log로 남는 내용
	fmt.Println("Received body: ", request.Body)

	return events.APIGatewayProxyResponse{
		Body: request.Body,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
