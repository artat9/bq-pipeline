package klimg

import "net/http"

type (
	// Image image
	Image struct {
		Data        []byte `json:"data"`
		ContentType string `json:"content_type"`
	}
)

// ToImage byte to image
func ToImage(data []byte) Image {
	return Image{
		Data:        data,
		ContentType: http.DetectContentType(data),
	}
}
