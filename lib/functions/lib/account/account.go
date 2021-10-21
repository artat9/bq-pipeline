package account

import (
	"crypto/rand"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type (
	// Account user
	Account struct {
		Address common.Address `json:"address"`
		Nonce   big.Int        `json:"nonce"`
	}
)

// New new user
func New(address common.Address) Account {
	//Max random value, a 130-bits integer, i.e 2^130 - 1
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))

	//Generate cryptographically strong pseudo-random between 0 - max
	n, _ := rand.Int(rand.Reader, max)

	return Account{
		Address: address,
		Nonce:   *n,
	}
}
