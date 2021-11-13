package ipfsclient

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"kaleido-backend/pkg/common/log"
	"kaleido-backend/pkg/mediaaccount"

	shell "github.com/ipfs/go-ipfs-api"
)

type (
	// Client client
	Client interface {
		Add(r io.Reader, options ...shell.AddOpts) (string, error)
		Pin(path string) error
	}
	// Service service
	Service struct {
		client Client
	}
)

// New new cilent
func New(c Client) Service {
	return Service{
		client: c,
	}
}

// NewIpfsClient new ipfs client
func NewIpfsClient() Service {
	return New(shell.NewShell("api.thegraph.com/ipfs/"))
}

// UploadMedia upload media info
func (s Service) UploadMedia(ctx context.Context, info mediaaccount.PublicInfo) error {
	return s.upload(ctx, info)
}

func (s Service) upload(ctx context.Context, target interface{}) error {
	b, err := json.Marshal(&target)
	if err != nil {
		log.Error("json marshall faile", err)
		return err
	}
	cid, err := s.client.Add(bytes.NewReader(b))
	if err != nil {
		log.Error("add failed", err)
		return err
	}
	return s.client.Pin(cid)

}
