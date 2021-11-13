package mediaaccount

import "context"

type (
	// Uploader upload
	Uploader interface {
		UploadMedia(ctx context.Context, info PublicInfo) error
	}
	// Creator media creator
	Creator struct {
		uploader Uploader
	}
)

// NewMediaCreator new media creator
func NewMediaCreator(uploader Uploader) Creator {
	return Creator{
		uploader: uploader,
	}
}

// NewMedia create new media
func (c Creator) NewMedia(ctx context.Context, info PublicInfo) error {
	return c.uploader.UploadMedia(ctx, info)
}
