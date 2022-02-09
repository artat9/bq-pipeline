package user

import (
	"context"
)

type (
	User struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	Service struct {
		uploader   Uploader
		downloader Downloader
	}
	Uploader interface {
		Upload(ctx context.Context, users []User) error
	}
	Downloader interface {
		Download(ctx context.Context, it interface{}) ([]*User, error)
	}
)

//TODO FIX
func NewService(u Uploader, d Downloader) *Service {
	return &Service{
		uploader:   u,
		downloader: d,
	}
}

func (s *Service) Upload(ctx context.Context, users []User) error {
	return s.uploader.Upload(ctx, users)
}

func (s *Service) Downlaod(ctx context.Context, it interface{}) ([]*User, error) {
	return s.downloader.Download(ctx, it)
}
