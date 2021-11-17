package media

import "github.com/ethereum/go-ethereum/common"

type (
	// Proxy proxy
	Proxy common.Address
	// EOA eoa
	EOA common.Address

	// Media media
	Media struct {
		Proxy       Proxy
		EOA         EOA
		MailAddress string
	}
)

// NewMedia create a media
func NewMedia(p Proxy, e EOA, mailAddress string) Media {
	return Media{
		Proxy:       p,
		EOA:         e,
		MailAddress: mailAddress,
	}
}
