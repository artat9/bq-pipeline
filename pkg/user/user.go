package user

import (
	"context"
	"os/user"
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
		Upload(ctx context.Context) error
	}
	Downloader interface {
		Download(ctx context.Context, it []interface{}) ([]*user.User, error)
	}
)

//TODO FIX
func NewService(u Uploader, d Downloader) *Service {
	return &Service{
		uploader:   u,
		downloader: d,
	}
}

func (s Service) Upload(ctx context.Context) error {
	return s.uploader.Upload(ctx)
}

func (s Service) Downlaod(ctx context.Context, it []interface{}) ([]*user.User, error) {
	return s.downloader.Download(ctx, it)
}
