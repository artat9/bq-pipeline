package handle

import (
	"context"
	"encoding/json"
	"kaleido-backend/pkg/account"
	"kaleido-backend/pkg/infrastructure/auth"
	"kaleido-backend/pkg/infrastructure/signature"
	"kaleido-backend/pkg/infrastructure/ssm"
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

// EOA retrive eoa from string
func EOA(ctx context.Context, jwt string) (common.Address, error) {
	v, err := auth.NewVerifier(ctx, ssm.New())
	if err != nil {
		return common.Address{}, err
	}
	return account.NewVerifyService(v).Verify(ctx, jwt)
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
	k := arguments(event)
	if k == nil {
		return ""
	}
	return k[key].(string)
}

func IntArgument(event interface{}, key string) int {
	k := arguments(event)
	if k == nil {
		return 0
	}
	return k[key].(int)
}

func arguments(event interface{}) map[string]interface{} {
	req := map[string]interface{}{}
	b, _ := json.Marshal(&event)
	json.Unmarshal(b, &req)
	input := req["arguments"]
	kv, ok := input.(map[string]interface{})
	if !ok {
		return nil
	}
	return kv
}

// Sign sign
func Sign(event interface{}) account.SignInInput {
	args := arguments(event)
	if args == nil {
		return account.SignInInput{}
	}
	sign := args["sign"]
	input := account.SignInInput{}
	b, _ := json.Marshal(&sign)
	json.Unmarshal(b, &input)
	return input
}
