package handle

import (
	"encoding/json"
	"kaleido-backend/pkg/account"
	"kaleido-backend/pkg/infrastructure/signature"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ethereum/go-ethereum/common"
)

// Headers with headers
func Headers(request events.APIGatewayProxyRequest) map[string]string {
	return map[string]string{
		"Access-Control-Allow-Headers":     "*",
		"Access-Control-Allow-Methods":     "GET,POST,PUT,DELETE",
		"Access-Control-Allow-Credentials": "true",
		"Access-Control-Allow-Origin":      "*",
	}
}

// EOAFromSign get eoa from sign
// TODO: inputのeoaと検証
func EOAFromSign(event interface{}) (common.Address, error) {
	sign := Sign(event)
	return signature.SignedMessage{}.Recover(sign.Msg, sign.Sig)
}

func withoutProtocol(origin string) string {
	return strings.ReplaceAll(strings.ReplaceAll(origin, "https://", ""), "http://", "")
}

// AppsyncInput input
type AppsyncInput struct {
	Arguments Arguments `json:"arguments"`
}

// Arguments args
type Arguments struct {
	Input interface{} `json:"input"`
}

// UnauthorizedResponse unauthorized response
func UnauthorizedResponse(request events.APIGatewayProxyRequest, err error) events.APIGatewayProxyResponse {
	return ErrorResponse(request, 401, err)
}

// ErrorResponse error response
func ErrorResponse(request events.APIGatewayProxyRequest, code int, err error) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Headers:    Headers(request),
		Body:       err.Error(),
	}
}

// ClientErrorResponse error response
func ClientErrorResponse(request events.APIGatewayProxyRequest, err error) events.APIGatewayProxyResponse {
	return ErrorResponse(request, 400, err)
}

// UnexpectedErrorResponse error response
func UnexpectedErrorResponse(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Headers:    Headers(request),
		Body:       "unknown error",
	}
}

// NotFoundErrorResponse error response
func NotFoundErrorResponse(request events.APIGatewayProxyRequest, err error) events.APIGatewayProxyResponse {
	return ErrorResponse(request, 404, err)
}

// NormalResponse success response
func NormalResponse(request events.APIGatewayProxyRequest, res interface{}) events.APIGatewayProxyResponse {
	resJSON, _ := json.Marshal(&res)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    Headers(request),
		Body:       string(resJSON),
	}
}

// Authorization auth value
func Authorization(event interface{}) string {
	req := map[string]interface{}{}
	b, _ := json.Marshal(&event)
	json.Unmarshal(b, &req)
	input := req["request"]
	headers := input.(map[string]interface{})
	auth := headers["authorization"]
	if auth != nil {
		return auth.(string)
	}
	return ""
}

// Argument argument
func Argument(event interface{}, key string) string {
	k := input(event)
	if k == nil {
		return ""
	}
	return k[key].(string)
}

// IntArgument int argument
func IntArgument(event interface{}, key string) int {
	k := input(event)
	if k == nil {
		return 0
	}
	value := k[key]
	switch value.(type) {
	case int:
		b := value.(int)
		return b
	case float64:
		n := value.(float64)
		return int(n)
	}
	return 0
}

func input(event interface{}) map[string]interface{} {
	req := map[string]interface{}{}
	b, _ := json.Marshal(&event)
	json.Unmarshal(b, &req)
	arguments := req["arguments"]
	kv, ok := arguments.(map[string]interface{})
	if !ok {
		return nil
	}
	input, ok := kv["input"]
	in, ok := input.(map[string]interface{})
	if !ok {
		return nil
	}
	return in
}

// Sign sign
func Sign(event interface{}) account.SignInInput {
	args := input(event)
	if args == nil {
		return account.SignInInput{}
	}
	sign := args["sign"]
	input := account.SignInInput{}
	b, _ := json.Marshal(&sign)
	json.Unmarshal(b, &input)
	return input
}
