package nft

import "strings"

const (
	baseURI               = "https://arweave.net/"
	donationImageURL      = "https://arweave.net/4gBK-mr7XeUXKemXE8AIQTsr2HfBCFZt15jPdINjvt4"
	refundRequestImageURL = "JK0imUjOSh7W07Sh6ts2dUaf83b4wOQIY0rXrOTXK2o"
)

type (
	// NFT nft
	NFT struct {
		ImageURL    string `json:"imageUrl"`
		MetadataURL string `json:"metadataUrl"`
	}
	// Metadata metadata
	Metadata struct {
		Name         string `json:"name"`
		Description  string `json:"description"`
		Image        string `json:"image"`
		AnimationURL string `json:"animation_url"`
	}

	//PostInput input
	PostInput struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Image       Image  `json:"image"`
	}

	// Image image
	Image struct {
		Data        []byte `json:"data"`
		ContentType string `json:"contentType"`
	}
)

// NewMetadata create metadata
func NewMetadata(title, description, imageURL string) Metadata {
	return newMetadata(title, description, imageURL, imageURL)
}

func newMetadata(title, description, imageURL, animationURL string) Metadata {
	if !strings.Contains(imageURL, baseURI) {
		imageURL = baseURI + imageURL
	}
	if !strings.Contains(animationURL, baseURI) {
		animationURL = baseURI + animationURL
	}
	return Metadata{
		Name:         title,
		Description:  description,
		Image:        imageURL,
		AnimationURL: animationURL,
	}
}

// NewPostMetadata new post metadata
func NewPostMetadata(title, description, imageURL string) Metadata {
	return newMetadata(title, description, imageURL, donationImageURL)
}

// NewRefundMetadata new Refundrequest metadata
func NewRefundMetadata() Metadata {
	return newMetadata("REFUND REQUEST", "The donor has applied for a refund. By refunding through the donation page, this card will be burned. Do you accept the refund request?", refundRequestImageURL, refundRequestImageURL)
}
