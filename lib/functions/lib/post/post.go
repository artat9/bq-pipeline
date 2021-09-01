package post

import (
	"context"
	"net/http"
)

type (
	// Image image
	Image struct {
		Data        []byte `json:"data"`
		ContentType string `json:"content_type"`
	}
	// Input input for post
	Input struct {
		Content
		KeyVisual []byte   `json:"key_visual"`
		Images    [][]byte `json:"images,omitempty"`
	}

	// Content post content
	Content struct {
		Name             string           `json:"name"`
		Description      string           `json:"description"`
		AdSpecifications AdSpecifications `json:"ad_specifications"`
	}

	// AdSpecifications AdSpecifications
	AdSpecifications struct {
		Width  int    `json:"width"`
		Height int    `json:"height"`
		Link   string `json:"link"`
	}
	// Post post
	Post struct {
		Content
		ImageURL string   `json:"image_url"`
		Images   []string `json:"images"`
	}

	// Output output for post
	Output struct {
		Metadata string `json:"metadata"`
	}

	// Service post service
	Service struct {
		uploader Uploader
	}

	// Uploader uploader
	Uploader interface {
		UploadImage(ctx context.Context, img Image) (string, error)
		UploadPostMetadata(ctx context.Context, post Post) (string, error)
		ToURL(tx string) string
	}
)

// NewService new service
func NewService(uploader Uploader) *Service {
	return &Service{
		uploader: uploader,
	}
}

func toImage(data []byte) Image {
	return Image{
		Data:        data,
		ContentType: http.DetectContentType(data),
	}
}

func (s *Service) toPost(input Input, keyVisualHash string, imageHashes []string) Post {
	return Post{
		Content:  input.Content,
		ImageURL: s.uploader.ToURL(keyVisualHash),
		Images:   s.toURLs(imageHashes),
	}
}

func (s *Service) toURLs(hashes []string) []string {
	res := []string{}
	for _, hash := range hashes {
		res = append(res, s.uploader.ToURL(hash))
	}
	return res
}

// Upload puload post
func (s *Service) Upload(ctx context.Context, input Input) (Output, error) {
	kv, err := s.uploader.UploadImage(ctx, toImage(input.KeyVisual))
	if err != nil {
		return Output{}, err
	}
	imgs, err := s.uploadImages(ctx, input.Images)
	if err != nil {
		return Output{}, err
	}
	metadata, err := s.uploader.UploadPostMetadata(ctx, s.toPost(input, kv, imgs))
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
	image := toImage(img)
	return s.uploader.UploadImage(ctx, image)
}
