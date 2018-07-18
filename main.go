package main

import (
	"fmt"
	"net/url"
	"strings"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	// SlackBotUserID is ID of SlackBot 
	SlackBotUserID = "USLACKBOT"
)

func parseRequestParams(reqBody string) (map[string]string, error) {

	params := strings.Split(reqBody, "&")
	paramsMap := make(map[string]string)

	for _, p := range params {
		param := strings.Split(p, "=")
		paramsMap[param[0]], _ = url.QueryUnescape(param[1])
	}

	return paramsMap, nil
}

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	params, _ := parseRequestParams(request.Body)
	fmt.Println(params["text"])
	var bytes []byte

	if SlackBotUserID != params["user_id"] {
		text, err := ProcessCommand(params["text"])
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}

		response := Response{
			Text: text,
		}
	
		bytes, _ = json.Marshal(&response)
	}

	bodyText := ""
	if len(bytes) != 0 {
		bodyText = string(bytes)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       bodyText,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil

}

func main() {
	lambda.Start(Handler)
}
