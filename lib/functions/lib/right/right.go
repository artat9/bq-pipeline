package right

import (
	"context"
	"kaleido-backend/lib/functions/lib/common/log"
)

type (

	// Output output for post
	Output struct {
		Owner string `json:"owner"`
	}
	// Service post service
	Service struct {
		contract SmartContract
	}

	// GetInput get input
	GetInput struct {
		Account  string
		Metadata string
	}

	// SmartContract smart contract
	SmartContract interface {
		Owner(ctx context.Context, postID uint64) (string, error)
	}
)

// NewService new service
func NewService(contract SmartContract) *Service {
	return &Service{
		contract: contract,
	}
}

// OwnerOf get owner of
func (s *Service) OwnerOf(ctx context.Context, postID uint64) (Output, error) {
	owner, err := s.contract.Owner(ctx, postID)
	if err != nil {
		log.Error("get owner failed", err)
		return Output{
			Owner: "",
		}, nil
	}
	return Output{
		Owner: owner,
	}, nil
}
