package main

import (
	"context"
	"encoding/json"
	lib "kaleido-backend/pkg"
	"kaleido-backend/pkg/account"
	"kaleido-backend/pkg/common/log"
	"kaleido-backend/pkg/infrastructure/auth"
	"kaleido-backend/pkg/infrastructure/ssm"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req, err := toInput(request)
	if err != nil {
		return unauthorized(request, err), nil
	}
	app, err := newApp(ctx)
	if err != nil {
		log.Error("initialization failed", err)
		return unauthorized(request, err), err
	}
	res, err := app.Reflesh(ctx, req)
	if err != nil {
		return unauthorized(request, err), nil
	}
	resJSON, _ := json.Marshal(&res)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    lib.Headers(request),
		Body:       string(resJSON),
	}, nil
}

func main() {
	lambda.Start(handler)
}

func toInput(request events.APIGatewayProxyRequest) (account.RefleshInput, error) {
	req := account.RefleshInput{}
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		log.Error("json unmarshal faile", err)
		return account.RefleshInput{}, err
	}
	return req, nil
}

func newApp(ctx context.Context) (account.VerifyService, error) {
	sig, err := auth.NewVerifier(ctx, ssm.New())
	if err != nil {
		return account.VerifyService{}, err
	}
	return account.NewVerifyService(sig), nil
}

func unauthorized(request events.APIGatewayProxyRequest, err error) events.APIGatewayProxyResponse {
	return anyError(request, 401, err)
}

func anyError(request events.APIGatewayProxyRequest, code int, err error) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Headers:    lib.Headers(request),
		Body:       err.Error(),
	}
}
