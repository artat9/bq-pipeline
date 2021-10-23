package handle

import (
	"context"
	"encoding/json"
	"kaleido-backend/pkg/account"
	"kaleido-backend/pkg/infrastructure/auth"
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
	account.NewVerifyService(v).

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
