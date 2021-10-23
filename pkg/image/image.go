package image

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

type (
	// UpdateInput update input
	UpdateInput struct {
		Name string `json:"image"`
	}
	// UpdateOutput update output
	UpdateOutput struct {
		URL string `json:"url"`
	}
	// Service service
	Service struct {
		issuer Issuer
	}
	// Issuer upload url issuer
	Issuer interface {
		IssueUploadURL(ctx context.Context, filename string) (url string, err error)
	}
)

// New new service
func New(issuer Issuer) Service {
	return Service{
		issuer: issuer,
	}
}

// Update update image
func (s Service) Update(ctx context.Context, of common.Address, in UpdateInput) (UpdateOutput, error) {
	url, err := s.issuer.IssueUploadURL(ctx, imageName(of, in.Name))
	return UpdateOutput{
		URL: url,
	}, err
}

func imageName(of common.Address, filename string) string {
	return of.Hex() + "/" + filename
}
