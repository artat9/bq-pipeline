package refund

import (
	"aurora-backend/lib/functions/lib/nft"
	"context"
)

type (
	// Service post service
	Service struct {
		uploader Uploader
	}

	// Uploader uploader
	Uploader interface {
		UploadNFTMetadata(ctx context.Context, imageTx string, request nft.Metadata) (string, error)
	}
)

// NewService new service
func NewService(uploader Uploader) *Service {
	return &Service{
		uploader: uploader,
	}
}

// Upload puload post
func (s *Service) Upload(ctx context.Context) (string, string, error) {
	metadata := nft.NewRefundMetadata()
	meta, err := s.uploader.UploadNFTMetadata(ctx, metadata.Image, metadata)
	return meta, metadata.Image, err
}
