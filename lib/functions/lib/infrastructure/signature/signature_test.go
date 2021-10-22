package signature

import (
	"testing"
)

const (
	messageSignature = "0x03574b2456ba7e3e4fc9f62004c9160dbebbf7731b9b666372ab0d62dc08e5497ba3c6864b6cdadf8211bac776edeeda3426abbdfa2f01851d79af51a38cfcfd1c"
	message          = "{\"propertyType\":\"Discord User ID\",\"propertyId\":\"895234765119168522\",\"evidence\":\"\",\"method\":\"Claime Discord App\"}"
	expectedAddress  = "0x8dc81F896B38167734ca4ff26b1D20C4c78e9190"
)

func TestRecover(t *testing.T) {
	t.Run("recover address", func(t *testing.T) {
		addr, err := SignedMessage{}.Recover(message, messageSignature)
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
