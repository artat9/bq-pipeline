package receipt

import (
	"aurora-backend/lib/functions/lib/nft"
	"aurora-backend/lib/functions/lib/post"
	"context"
	"errors"
	"testing"
)

type (
	fakeUploader struct {
		uploadImageFunc       func(ctx context.Context, img post.Image) (string, error)
		uploadNFTMetadataFunc func(ctx context.Context, imageTx string, request nft.Metadata) (string, error)
	}
	fakeContextResolver struct {
		contextFunc func(ctx context.Context, postID string, address string) (ContractContext, error)
	}
	fakeNFTGenerator struct {
		generateFunc func(ctx context.Context, param GenerateParameter) ([]byte, error)
	}
	fakeRepository struct {
		toInfoFunc func(ctx context.Context, postID, metadataURI string) (Info, error)
	}
)

func (upld fakeUploader) UploadImage(ctx context.Context, img post.Image) (string, error) {
	return upld.uploadImageFunc(ctx, img)
}

func (upld fakeUploader) UploadNFTMetadata(ctx context.Context, imageTx string, request nft.Metadata) (string, error) {
	return upld.uploadNFTMetadataFunc(ctx, imageTx, request)
}

func (res fakeContextResolver) Context(ctx context.Context, postID string, address string) (ContractContext, error) {
	return res.contextFunc(ctx, postID, address)
}

func (gen fakeNFTGenerator) Generate(ctx context.Context, param GenerateParameter) ([]byte, error) {
	return gen.generateFunc(ctx, param)
}

func (rep fakeRepository) ToInfo(ctx context.Context, postID, metadataURI string) (Info, error) {
	return rep.toInfoFunc(ctx, postID, metadataURI)
}

func newDefaultService() Service {
	return Service{
		rep:       defaultRep(),
		resolver:  defaultContextResolver(),
		generator: defaultNFTGenerator(),
		uploader:  defaultUploader(),
	}
}

func defaultUploader() Uploader {
	return fakeUploader{
		uploadImageFunc: func(ctx context.Context, img post.Image) (string, error) {
			return "", nil
		},
		uploadNFTMetadataFunc: func(ctx context.Context, imageTx string, request nft.Metadata) (string, error) {
			return "", nil
		},
	}
}

func defaultContextResolver() ContextResolver {
	return fakeContextResolver{
		contextFunc: func(ctx context.Context, postID, address string) (ContractContext, error) {
			return ContractContext{}, nil
		},
	}
}

func defaultNFTGenerator() NFTGenerator {
	return fakeNFTGenerator{
		generateFunc: func(ctx context.Context, param GenerateParameter) ([]byte, error) {
			return []byte{}, nil
		},
	}
}

func defaultRep() Repository {
	return fakeRepository{
		toInfoFunc: func(ctx context.Context, postID, metadataURI string) (Info, error) {
			return Info{}, nil
		},
	}
}

func TestUltraRare(t *testing.T) {
	t.Run("ultrarare receipt in case ReceiptID is endswith 7", func(t *testing.T) {
		inputs := []SerialNum{
			{Num: 99907},
			{Num: 99817},
			{Num: 99727},
			{Num: 99637},
			{Num: 99547},
			{Num: 99457},
			{Num: 99367},
			{Num: 99277},
			{Num: 99187},
			{Num: 99097},
		}
		for _, in := range inputs {
			if !in.UltraRare() {
				t.Error(in.Num, " is not ultrarare")
			}
		}
	})
	t.Run("normal receipt in case ReceiptID is not endswith 7", func(t *testing.T) {
		inputs := []SerialNum{
			{Num: 99909},
			{Num: 99818},
			{Num: 99726},
			{Num: 99635},
			{Num: 99544},
			{Num: 99453},
			{Num: 99362},
			{Num: 99271},
			{Num: 99180},
		}
		for _, in := range inputs {
			if in.UltraRare() {
				t.Error(in.Num, " is ultrarare")
			}
		}
	})
}

func TestIssue(t *testing.T) {
	t.Run("metadata created", func(t *testing.T) {
		target := newDefaultService()
		if _, err := target.Issue(context.Background(), Input{}); err != nil {
			t.Error("unexpected error occured", err)
		}
	})
	t.Run("error if context cannot resolved", func(t *testing.T) {
		target := newDefaultService()
		target.resolver = fakeContextResolver{
			contextFunc: func(ctx context.Context, postID, address string) (ContractContext, error) {
				return ContractContext{}, errors.New("")
			},
		}
		assertError(t, target)
	})
	t.Run("error if resolve info failed", func(t *testing.T) {
		target := newDefaultService()
		target.rep = fakeRepository{
			toInfoFunc: func(ctx context.Context, postID, metadataURI string) (Info, error) {
				return Info{}, errors.New("")
			},
		}
		assertError(t, target)
	})
	t.Run("error if glb generate failed", func(t *testing.T) {
		target := newDefaultService()
		target.generator = fakeNFTGenerator{
			generateFunc: func(ctx context.Context, param GenerateParameter) ([]byte, error) {
				return []byte{}, errors.New("")
			},
		}
		assertError(t, target)
	})
	t.Run("error if upload image failed", func(t *testing.T) {
		target := newDefaultService()
		target.uploader = fakeUploader{
			uploadImageFunc: func(ctx context.Context, img post.Image) (string, error) {
				return "", errors.New("")
			},
		}
		assertError(t, target)
	})
}

func assertError(t *testing.T, s Service) {
	if _, err := s.Issue(context.Background(), Input{}); err == nil {
		t.Error("no error occured")
	}
}
