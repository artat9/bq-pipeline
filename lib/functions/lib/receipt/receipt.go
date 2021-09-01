package receipt

import (
	"aurora-backend/lib/functions/lib/common/log"
	"aurora-backend/lib/functions/lib/nft"
	"aurora-backend/lib/functions/lib/post"
	"context"
	"fmt"
	"time"
)

type (
	// Input input
	Input struct {
		PostID  string `json:"postId"`
		Address string `json:"address"`
		Amount  string `json:"amount"`
	}
	// Output output
	Output struct {
		Metadata string `json:"metadata"`
	}

	// Service service
	Service struct {
		rep       Repository
		resolver  ContextResolver
		generator NFTGenerator
		uploader  Uploader
	}
	//Uploader uploader
	Uploader interface {
		UploadImage(ctx context.Context, img post.Image) (string, error)
		UploadNFTMetadata(ctx context.Context, imageTx string, request nft.Metadata) (string, error)
	}

	// SerialNum SerialNum of receipt
	SerialNum struct {
		Num int
	}

	// Post Post content
	Post struct {
		ProjectName string
		Description string
	}

	// Info info
	Info struct {
		Post     Post
		SerialNo int
		Issued   time.Time
	}

	// ContractContext info from contract
	ContractContext struct {
		MetadataURI         string
		ReceiptID, SerialNo int
	}

	// GenerateParameter parameter for generating receipt
	GenerateParameter struct {
		Context ContractContext
		Info    Info
		Amount  string
	}

	// ContextResolver contract ctx resolver
	ContextResolver interface {
		Context(ctx context.Context, postID string, address string) (ContractContext, error)
	}

	//NFTGenerator generator
	NFTGenerator interface {
		Generate(ctx context.Context, param GenerateParameter) ([]byte, error)
	}

	// Repository repository
	Repository interface {
		ToInfo(ctx context.Context, postID, metadataURI string) (Info, error)
	}
)

// UltraRare whether a Receipt with a SerialNum is UltraRare
func UltraRare(num int) bool {
	return num%10 == 7
}

func toGenerateParameter(in Input, context ContractContext, info Info) GenerateParameter {
	return GenerateParameter{
		Context: context,
		Info:    info,
		Amount:  in.Amount,
	}
}

// NewService new service
func NewService(rep Repository, resolver ContextResolver, generator NFTGenerator, uploader Uploader) Service {
	return Service{
		rep:       rep,
		resolver:  resolver,
		generator: generator,
		uploader:  uploader,
	}
}

// Issue issue new receipt.
func (s *Service) Issue(ctx context.Context, in Input) (string, string, error) {
	fmt.Println("postID")
	fmt.Println(in.PostID)
	contractCtx, err := s.resolver.Context(ctx, in.PostID, in.Address)
	if err != nil {
		log.Error("resolve context failed", err)
		return "", "", err
	}
	info, err := s.rep.ToInfo(ctx, in.PostID, contractCtx.MetadataURI)
	if err != nil {
		log.Error("to info failed", err)
		return "", "", err
	}
	param := toGenerateParameter(in, contractCtx, info)
	generated, err := s.generator.Generate(ctx, param)
	if err != nil {
		log.Error("generate glb failed", err)
		return "", "", err
	}
	res, err := s.uploader.UploadImage(ctx, post.Image{
		Data:        generated,
		ContentType: "model/gltf-binary",
	})
	if err != nil {
		log.Error("upload image failed", err)
		return "", "", err
	}
	metadata, err := s.uploader.UploadNFTMetadata(ctx, res, nft.NewMetadata(info.Post.ProjectName, info.Post.Description, res))
	if err != nil {
		log.Error("upload metadata failed", err)
	}
	return metadata, res, nil
}
