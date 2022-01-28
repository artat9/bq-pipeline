package user

import "context"

type (
	User struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	Service struct {
		uploader Uploader
	}
	Uploader interface {
		Upload(ctx context.Context, u User) error
	}
)

func (s Service) Upload(ctx context.Context, u User) error {
	return s.uploader.Upload(ctx, u)
}
