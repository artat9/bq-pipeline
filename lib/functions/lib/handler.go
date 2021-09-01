package lib

import (
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

// Headers with headers
func Headers(request events.APIGatewayProxyRequest) map[string]string {
	return map[string]string{
		"Access-Control-Allow-Headers":     "*",
		"Access-Control-Allow-Methods":     "GET,POST,PUT,DELETE",
		"Access-Control-Allow-Credentials": "true",
		"Access-Control-Allow-Origin":      withoutProtocol(request.Headers["origin"]),
	}
}

func withoutProtocol(origin string) string {
	return strings.ReplaceAll(strings.ReplaceAll(origin, "https://", ""), "http://", "")
}
