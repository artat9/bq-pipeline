package ethsignature

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	messagePrefix = "\x19Ethereum Signed Message:\n"
)

type (
	SignedMessage struct {
		timestamper timestamper
	}
	timestamper interface {
		Now() time.Time
	}
)

// Recover recover EOA And Claim information from input
func (s SignedMessage) Recover(msg, sig string) (common.Address, error) {
	expired, err := s.expired(msg)
	if err != nil {
		return common.Address{}, err
	}
	if expired {
		return common.Address{}, errors.New("sign expired")
	}
	eoa, err := recoverFrom(msg, sig)
	if err != nil {
		return common.Address{}, err
	}
	return eoa, err
}

func (s SignedMessage) expired(msg string) (bool, error) {
	ts, err := recoverTimestamp(msg)
	if err != nil {
		return true, err
	}
	validity := time.Unix(ts, 0)
	var now time.Time
	if s.timestamper != nil {
		now = s.timestamper.Now()
	} else {
		now = time.Now()
	}
	if validity.Before(now) {
		return true, nil
	}
	return false, nil
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

func recoverTimestamp(message string) (int64, error) {
	var val map[string]string
	err := json.Unmarshal([]byte(message), &val)
	if err != nil {
		return 0, err
	}
	ts, err := strconv.Atoi(val["validity"])
	if err != nil {
		return 0, err
	}
	return int64(ts), nil
}
