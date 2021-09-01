package postmanager

import (
	"aurora-backend/lib/functions/lib/common/log"
	"aurora-backend/lib/functions/lib/contracts"
	"aurora-backend/lib/functions/lib/receipt"
	"context"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

// Context retrive context from contract.
func (p Provider) Context(ctx context.Context, postID string, address string) (receipt.ContractContext, error) {
	pm, err := p.newPostManager()
	if err != nil {
		log.Error("postmanager create failed", err)

		return receipt.ContractContext{}, err
	}

	posts, err := pm.AllPosts(callOpts(ctx), common.HexToHash(postID).Big())
	if err != nil {
		log.Error("get posts failed", err)
		return receipt.ContractContext{}, err
	}

	seed, err := pm.NextReceiptSeed(callOpts(ctx))
	if err != nil {
		log.Error("retrive next receipt id failed", err)

		return receipt.ContractContext{}, err
	}
	postIDBig, err := hexutil.DecodeBig(postID)
	if err != nil {
		log.Error("decode fialed", err)

		return receipt.ContractContext{}, err
	}
	nextSerialNum := nextSerialNum(posts.DonatedCount)
	res, err := pm.ComputeDonationReceiptId(callOpts(ctx), common.BigToAddress(postIDBig), seed, nextSerialNum)
	if err != nil {
		log.Error("get next receipt id  failed", err)
	}
	return receipt.ContractContext{
		MetadataURI: posts.MetadataURI,
		ReceiptID:   int(res.Uint64()),
		SerialNo:    int(nextSerialNum.Int64()),
	}, nil
}

func nextSerialNum(donatedCount *big.Int) *big.Int {
	return big.NewInt(0).Add(donatedCount, big.NewInt(1))
}

func callOpts(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
}

func (p Provider) newPostManager() (*contracts.Contracts, error) {
	return contracts.NewContracts(common.HexToAddress(postManagerAddress), p.cli)
}

func providerURL() string {
	if os.Getenv("EnvironmentId") == "prod" {
		return "https://polygon-mainnet.g.alchemy.com/v2/9apanUOHDQhhzl4RJHtcxY8SpMbH3QWJ"
	}
	return "https://rinkeby.infura.io/v3/eb8face5ffc84dc7bd8f2ddaef07dc34"
}
