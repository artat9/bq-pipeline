package ad

import (
	"context"
)

type (

	// Input input for advertisement
	Input struct {
		Image []byte `json:"image"`
		Link  string `json:"link"`
		Alt   string `json:"alt"`
	}

	// Advertisement Advertisement
	Advertisement struct {
		ImageURL string `json:"image_url"`
		Link     string `json:"link"`
		Alt      string `json:"alt"`
	}

	// Service post service
	Service struct {
		contract SmartContract
	}

	// GetInput get input
	GetInput struct {
		SpaceMetadata string
	}

	// GetOutput get output
	GetOutput struct {
		DistributedMetadata string `json:"distributedMetadata"`
	}

	// SmartContract smart contract
	SmartContract interface {
		DisplayByMetadata(ctx context.Context, input GetInput) (string, error)
	}
)

// NewWithContract new with contract
func NewWithContract(contract SmartContract) *Service {
	return &Service{
		contract: contract,
	}
}

// Get Get metadata
func (s *Service) Get(ctx context.Context, input GetInput) (GetOutput, error) {
	meta, err := s.contract.DisplayByMetadata(ctx, input)
	return GetOutput{
		DistributedMetadata: meta,
	}, err
}
