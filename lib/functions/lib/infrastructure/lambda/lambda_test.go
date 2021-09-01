package lambda

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func TestAmountImage(t *testing.T) {
	tg := invoker()
	t.Run("enable to invoke method", func(t *testing.T) {
		img, err := tg.AmountImage(context.Background(), "1", "ETH")
		if err != nil {
			t.Error(err)
		}
		if img == nil {
			t.Error("img not generated")
		}
	})
}

func TestDateImage(t *testing.T) {
	tg := invoker()
	t.Run("enable to invoke method", func(t *testing.T) {
		img, err := tg.DateImage(context.Background(), "2021/01/01")
		if err != nil {
			t.Error(err)
		}
		if img == nil {
			t.Error("img not generated")
		}
	})
}

func TestDescriptionImage(t *testing.T) {
	tg := invoker()
	t.Run("enable to invoke method", func(t *testing.T) {
		img, err := tg.DescriptionImage(context.Background(), "this is test")
		if err != nil {
			t.Error(err)
		}
		if img == nil {
			t.Error("img not generated")
		}
	})
}

func TestTitle(t *testing.T) {
	tg := invoker()
	t.Run("enable to invoke method", func(t *testing.T) {
		img, err := tg.TitleImage(context.Background(), "this is title")
		if err != nil {
			t.Error(err)
		}
		if img == nil {
			t.Error("img not generated")
		}
	})
}

func TestSerialNoImage(t *testing.T) {
	tg := invoker()
	t.Run("enable to invoke method", func(t *testing.T) {
		img, err := tg.SerialNoImage(context.Background(), 123)
		if err != nil {
			t.Error(err)
		}
		if img == nil {
			t.Error("img not generated")
		}
	})
}

func TestInvokeWithContext(t *testing.T) {
	tg := invoker()
	t.Run("invoke with context successfuly", func(t *testing.T) {
		payload := AmountInvokeInput{
			InvokeInput: InvokeInput{
				Type:  "amount",
				Value: "0",
			},
			Unit: "wei",
		}
		req, _ := toReqPayload(payload)
		res, err := tg.invokeWithContext(context.Background(), req, "textToPng")
		if err != nil {
			t.Error(err)
		}
		if res == nil {
			t.Error("response not extracted")
		}
		t.Error(string(res.Payload))
	})
}

func invoker() Invoker {
	sess := session.Must(session.NewSession())
	lmd := lambda.New(sess, aws.NewConfig().WithRegion("us-east-1"))
	return Invoker{
		svc: lmd,
	}
}
