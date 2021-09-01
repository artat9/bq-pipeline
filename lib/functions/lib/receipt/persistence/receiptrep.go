package persreceipt

import (
	"aurora-backend/lib/functions/lib/receipt"
	"context"
	"time"
)

type (

	// Repository repository
	Repository struct {
		storage PostStorage
	}

	// PostStorage post storage
	PostStorage interface {
		Post(ctx context.Context, metadataURI string) (receipt.Post, error)
	}
)

// New new repository
func New(storage PostStorage) *Repository {
	return &Repository{
		storage: storage,
	}
}

// ToInfo convert postID to Info.
func (r *Repository) ToInfo(ctx context.Context, postID, metadataURI string) (receipt.Info, error) {
	p, err := r.storage.Post(ctx, metadataURI)
	if err != nil {
		return receipt.Info{}, err
	}
	return receipt.Info{
		Post:     p,
		SerialNo: 0,
		Issued:   time.Now(),
	}, nil
}
