package main

import (
	"context"
	"encoding/json"
	"kaleido-backend/pkg/account"
	accountrep "kaleido-backend/pkg/account/persistence"
	"kaleido-backend/pkg/common/log"
	"kaleido-backend/pkg/handle"
	"kaleido-backend/pkg/infrastructure/auth"
	"kaleido-backend/pkg/infrastructure/ddb"
	"kaleido-backend/pkg/infrastructure/signature"
	"kaleido-backend/pkg/infrastructure/ssm"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req, err := toInput(request)
	if err != nil {
		return handle.UnauthorizedResponse(request, err), nil
	}
	app, err := newApp(ctx)
	if err != nil {
		log.Error("initialization failed", err)
		return handle.UnauthorizedResponse(request, err), err
	}
	res, err := app.Sign(ctx, req)
	if err != nil {
		return handle.UnauthorizedResponse(request, err), nil
	}
	return handle.NormalResponse(request, res), nil
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
