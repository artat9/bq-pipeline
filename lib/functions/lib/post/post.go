package post

import (
	"context"
)

type (
	// Image image
	Image struct {
		Data        []byte `json:"data"`
		ContentType string `json:"contentType"`
	}
	// Input input for post
	Input struct {
		Title            string `json:"title"`
		Description      string `json:"description"`
		Image            Image  `json:"image"`
		DescriptionImage []byte `json:"descriptionImage"`
		TitleImage       []byte `json:"titleImage"`
	}

	// Output output for post
	Output struct {
		ImageURL    string `json:"image"`
		MetadataURL string `json:"metadata"`
	}

	// Service post service
	Service struct {
		uploader Uploader
	}

	// Uploader uploader
	Uploader interface {
		UploadImage(ctx context.Context, img Image) (string, error)
		UploadPostMetadata(ctx context.Context, imageTx string, request Input) (string, error)
		ToURL(tx string) string
	}
)

// NewService new service
func NewService(uploader Uploader) *Service {
	return &Service{
		uploader: uploader,
	}
}

// Upload puload post
func (s *Service) Upload(ctx context.Context, input Input) (Output, error) {
	img, err := s.uploader.UploadImage(ctx, input.Image)
	if err != nil {
		return Output{}, err
	}
	metadata, err := s.uploader.UploadPostMetadata(ctx, img, input)
	if err != nil {
		return Output{}, err
	}

	return Output{
		ImageURL:    img,
		MetadataURL: metadata,
	}, err
}
