package lambda

import (
	"aurora-backend/lib/functions/lib/common/log"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"image"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type (
	// InvokeInput input
	InvokeInput struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}

	// AmountInvokeInput input
	AmountInvokeInput struct {
		InvokeInput
		Unit string `json:"unit"`
	}
	Invoker struct {
		svc *lambda.Lambda
	}
)

func NewInvoker() Invoker {
	return Invoker{
		svc: lambda.New(session.New()),
	}
}

func (i Invoker) AmountImage(ctx context.Context, amt, unit string) (image.Image, error) {
	in := AmountInvokeInput{
		InvokeInput: InvokeInput{
			Type:  "amount",
			Value: amt,
		},
		Unit: unit,
	}
	return i.invoke(ctx, in)
}

func (i Invoker) DateImage(ctx context.Context, dt string) (image.Image, error) {
	in := InvokeInput{
		Type:  "date",
		Value: dt,
	}
	return i.invoke(ctx, in)
}

func (i Invoker) DescriptionImage(ctx context.Context, description string) (image.Image, error) {
	in := InvokeInput{
		Type:  "description",
		Value: description,
	}
	return i.invoke(ctx, in)
}

func (i Invoker) TitleImage(ctx context.Context, title string) (image.Image, error) {
	in := InvokeInput{
		Type:  "title",
		Value: title,
	}
	return i.invoke(ctx, in)
}

func (i Invoker) SerialNoImage(ctx context.Context, serial string) (image.Image, error) {
	in := InvokeInput{
		Type:  "serialNo",
		Value: serial,
	}
	return i.invoke(ctx, in)
}

func (i Invoker) invokeWithContext(ctx context.Context, payload []byte, funcName string) (*lambda.InvokeOutput, error) {
	return i.svc.InvokeWithContext(ctx, &lambda.InvokeInput{
		FunctionName:   aws.String("arn:aws:lambda:us-east-1:495476032358:function:" + funcName),
		Payload:        payload,
		InvocationType: aws.String("RequestResponse"),
	})
}

func (i Invoker) executeFunction(ctx context.Context, payload []byte, functionName string) ([]byte, error) {
	resp, err := i.invokeWithContext(ctx, payload, functionName)
	if err != nil {
		log.Error("invoke function failed", err)
		return nil, err
	}
	decoded, err := decodePayload(resp.Payload)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func (i Invoker) invoke(ctx context.Context, in interface{}) (image.Image, error) {
	payload, err := toReqPayload(in)
	resp, err := i.executeFunction(ctx, payload, "textToPng")
	if err != nil {
		log.Error("invoke function failed", err)
		return nil, err
	}
	img, _, err := image.Decode(bytes.NewReader(resp))
	if err != nil {
		log.Error("decode 2 image failed", err)
		return nil, err
	}
	return img, nil
}

func toReqPayload(input interface{}) ([]byte, error) {
	payload, err := json.Marshal(&input)
	if err != nil {
		log.Error("json marshal failed", err)
		return nil, err
	}
	return payload, nil
}

func decodePayload(resp []byte) ([]byte, error) {
	base64Text, err := base64.StdEncoding.DecodeString(strings.ReplaceAll(string(resp), `"`, ""))
	if err != nil {
		log.Error("decode failed", err)
	}
	return base64Text, err
}
