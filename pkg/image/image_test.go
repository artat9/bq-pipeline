package image

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func Test_imageName(t *testing.T) {
	t.Run("image name should be has prefix address hex", func(t *testing.T) {
		account := common.HexToAddress("0x50414Ac6431279824df9968855181474c919a94B")
		filename := "test.txt"
		want := "0x50414Ac6431279824df9968855181474c919a94B/test.txt"
		got := imageName(account, filename)
		if want != got {
			t.Error("got:", got)
		}
	})
}
