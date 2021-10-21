package account

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang-jwt/jwt"
)

func TestNew(t *testing.T) {
	addr := common.HexToAddress("0x668F22f015BF2c91Bf4fb19a03619B8Ff593E8A8")
	u1 := New(addr)
	u2 := New(addr)
	t.Run("address should be matched", func(t *testing.T) {
		if addr != u1.Address {
			t.Error("got:", u1.Address)
		}
	})
	t.Run("nonce should be a random value", func(t *testing.T) {
		if u1.Nonce.Cmp(u2.Nonce) == 0 {
			t.Error("got:", u1.Nonce)
		}
	})
	jwt.NewWithClaims(&jwt.SigningMethodEd25519{}, jwt.MapClaims{})
}

func TestNewService(t *testing.T) {

}
