package admanager

import (
	"context"
	"kaleido-backend/pkg/ad"
	"kaleido-backend/pkg/common/log"
	contracts "kaleido-backend/pkg/contracts/adm"
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
)

// NewProvider new Provider
func NewProvider() (Provider, error) {
	client, err := ethclient.Dial(providerURL())
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
	_, err := p.newAdManager()
	if err != nil {
		log.Error("build admanager failed", err)
		return "", err
	}
	//metadata, err := ad.DisplayByMetadata(callOpts(ctx), common.HexToAddress(input.Account), input.Metadata)
	//
	//if err != nil {
	//	log.Error("display by index failed", err)
	//	return "", err
	//}
	return "", nil
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

func providerURL() string {
	if os.Getenv("EnvironmentId") == "prod" {
		return "https://polygon-mainnet.g.alchemy.com/v2/"
	}
	return "https://rinkeby.infura.io/v3/"
}
