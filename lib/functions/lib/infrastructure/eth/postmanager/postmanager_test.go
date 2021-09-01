package postmanager

import (
	"context"
	"math/big"
	"testing"
)

func TestNewProvider(t *testing.T) {
	t.Run("connect succeed", func(t *testing.T) {
		_, err := NewProvider()
		if err != nil {
			t.Error(err)
		}
	})
}

func TestComputeReceiptID(t *testing.T) {
	p, _ := NewProvider()
	t.Run("enable to compute next receipt ID", func(t *testing.T) {
		res, err := p.Context(context.Background(), "0x989681", "")
		if err != nil {
			t.Error(err)
		}
		if res.ReceiptID == 0 {
			t.Error("receipt id not computed")
		}
	})
	t.Run("error if invalid postID", func(t *testing.T) {
		t.Run("addres cannot decode to big", func(t *testing.T) {
			_, err := p.Context(context.Background(), "test", "")
			if err == nil {
				t.Error("no error occured")
			}
		})
	})
}

func TestNextSerialNum(t *testing.T) {
	t.Run("enable to compute next serial num", func(t *testing.T) {
		current := big.NewInt(1)
		want := big.NewInt(2)
		got := nextSerialNum(current)
		if got.Cmp(want) != 0 {
			t.Error("want:", want, ", but got ", got)
		}
	})
}
