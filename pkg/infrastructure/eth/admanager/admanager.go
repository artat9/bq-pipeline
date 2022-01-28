package admanager

import (
	"bq-pipeline/pkg/ad"
	"bq-pipeline/pkg/common/log"
	contracts "bq-pipeline/pkg/contracts/adm"
	"context"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type (

	// Provider eth provider
	Provider struct {
		cli *ethclient.Client
	}
	// KeyResolver key resolver
	KeyResolver interface {
		InfuraKey(ctx context.Context) (string, error)
	}
)

// NewProvider new Provider
func NewProvider(ctx context.Context, kr KeyResolver) (Provider, error) {
	k, err := kr.InfuraKey(ctx)
	if err != nil {
		log.Error("key not found", err)
		return Provider{}, err
	}
	client, err := ethclient.Dial(providerURL(k))
	if err != nil {
		log.Error("provider create failed", err)
		return Provider{}, err
	}
	return Provider{
		cli: client,
	}, nil
}

// DisplayByMetadata get a metadata from a postmetadata
func (p Provider) DisplayByMetadata(ctx context.Context, input ad.GetInput) (string, error) {
	ad, err := p.newAdManager()
	if err != nil {
		log.Error("build admanager failed", err)
		return "", err
	}
	metadata, err := ad.Display(callOpts(ctx), input.SpaceMetadata)

	if err != nil {
		log.Error("display by index failed", err)
		return "", err
	}
	return metadata, nil
}

func callOpts(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
}

func (p Provider) newAdManager() (*contracts.Contracts, error) {
	return contracts.NewContracts(common.HexToAddress(admanagerAddress), p.cli)
}

func providerURL(infuraKey string) string {
	if os.Getenv("EnvironmentId") == "prod" {
		return "https://polygon-mainnet.g.alchemy.com/v2/" + infuraKey
	}
	return "https://rinkeby.infura.io/v3/" + infuraKey
}
