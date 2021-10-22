package signature

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	messagePrefix = "\x19Ethereum Signed Message:\n"
)

type (
	SignedMessage struct{}
)

// RecoverAddress recover EOA And Claim information from input
func (s SignedMessage) RecoverAddress(msg, sig string) (common.Address, error) {
	return recoverFrom(msg, sig)
}

func recoverFrom(from, sig string) (common.Address, error) {
	recovered, err := recoverAddressFromMessage(from, sig)
	if err != nil {
		return common.Address{}, err
	}
	return recovered, err
}

func recoverAddressFromMessage(message string, signature string) (common.Address, error) {
	return recoverAddress(signHash([]byte(message)), signature)
}

func recoverAddress(hash []byte, signature string) (common.Address, error) {
	sigArr, err := hexutil.Decode(signature)
	if err != nil {
		return common.Address{}, err
	}
	sigArr[64] -= 27
	rpk, err := crypto.Ecrecover(hash, sigArr)
	if err != nil {
		return common.Address{}, err
	}
	pubKey, err := crypto.UnmarshalPubkey(rpk)
	if err != nil {
		return common.Address{}, err
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return recoveredAddr, err
}

func signHash(data []byte) []byte {
	msg := fmt.Sprintf(messagePrefix+"%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}
