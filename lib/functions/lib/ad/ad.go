package ad

import (
	"aurora-backend/lib/functions/lib/klimg"
	"context"
)

type (

	// Input input for advertisement
	Input struct {
		Image []byte `json:"image"`
		Link  string `json:"link"`
		Alt   string `json:"alt"`
	}
	// Output output for post
	Output struct {
		Metadata string `json:"metadata"`
	}

	// Advertisement Advertisement
	Advertisement struct {
		ImageURL string `json:"image_url"`
		Link     string `json:"link"`
		Alt      string `json:"alt"`
	}

	// Service post service
	Service struct {
		uploader Uploader
	}

	// Uploader uploader
	Uploader interface {
		UploadImage(ctx context.Context, img klimg.Image) (string, error)
		UploadAdvertisementMetadata(ctx context.Context, advertisement Advertisement) (string, error)
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
	img := klimg.ToImage(input.Image)
	imgTx, err := s.uploader.UploadImage(ctx, img)
	if err != nil {
		return Output{}, err
	}
	metadata, err := s.uploader.UploadAdvertisementMetadata(ctx, Advertisement{
		ImageURL: s.uploader.ToURL(imgTx),
		Link:     input.Link,
		Alt:      input.Alt,
	})
	if err != nil {
		return Output{}, err
	}
	return Output{
		Metadata: metadata,
	}, err
}

func (s *Service) uploadImages(ctx context.Context, imgs [][]byte) ([]string, error) {
	res := []string{}
	for _, img := range imgs {
		hash, err := s.uploadImage(ctx, img)
		if err != nil {
			return nil, err
		}
		res = append(res, hash)
	}
	return res, nil
}

func (s *Service) uploadImage(ctx context.Context, img []byte) (string, error) {
	image := klimg.ToImage(img)

	return s.uploader.UploadImage(ctx, image)
}
