package signature

import (
	"testing"
	"time"
)

const (
	messageSignature = "0xfdf923c4588ba8bbb305bae866dbd4f13dcef982a1b1e5700e8a53da37b29d59655ac6a734c7a26fd861cc54f24897cc6fa32d52aa97584b96de2beaf49b5a7c1c"
	message          = "{\"validity\":\"1636951865\"}"
	expectedAddress  = "0xCdfc500F7f0FCe1278aECb0340b523cD55b3EBbb"
)

var target = SignedMessage{}

type (
	fakeTimestamper struct{}
)

func (ft fakeTimestamper) Now() time.Time {
	return time.Date(2000, 1, 1, 1, 1, 1, 1, time.Local)
}

func init() {
	target.timestamper = fakeTimestamper{}
}

func TestRecover(t *testing.T) {
	t.Run("recover address", func(t *testing.T) {
		addr, err := target.Recover(message, messageSignature)
		if err != nil {
			t.Error(err)
		}
		if expectedAddress != addr.Hex() {
			t.Error("got:", addr.Hex())
		}
	})
}

func TestRecoverFrom(t *testing.T) {
	t.Run("error if decode msg failed", func(t *testing.T) {
		if _, err := recoverFrom("aaa", ""); err == nil {
			t.Error("no error occured")
		}
	})
}
func TestRecoverAddress(t *testing.T) {
	t.Run("error if decode msg failed", func(t *testing.T) {
		if _, err := recoverAddress([]byte{}, messageSignature); err == nil {
			t.Error("no error occured")
		}
	})
}
