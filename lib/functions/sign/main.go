package main

import (
	"context"
	"encoding/json"
	"kaleido-backend/lib/functions/lib"
	"kaleido-backend/lib/functions/lib/account"
	accountrep "kaleido-backend/lib/functions/lib/account/persistence"
	"kaleido-backend/lib/functions/lib/common/log"
	"kaleido-backend/lib/functions/lib/infrastructure/auth"
	"kaleido-backend/lib/functions/lib/infrastructure/ddb"
	"kaleido-backend/lib/functions/lib/infrastructure/signature"
	"kaleido-backend/lib/functions/lib/infrastructure/ssm"

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
	res, err := app.Sign(ctx, req)
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

func toInput(request events.APIGatewayProxyRequest) (account.SignInInput, error) {
	req := account.SignInInput{}
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		log.Error("json unmarshal faile", err)
		return account.SignInInput{}, err
	}
	return req, nil
}

func newApp(ctx context.Context) (account.SignService, error) {
	sig, err := auth.NewSigner(ctx, ssm.New())
	if err != nil {
		return account.SignService{}, err
	}
	return account.NewSignService(signature.SignedMessage{}, accountrep.New(ddb.New()), sig), nil
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
