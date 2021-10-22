package accountrep

import (
	"context"
	"kaleido-backend/lib/functions/lib/account"
	"kaleido-backend/lib/functions/lib/common/log"
	"kaleido-backend/lib/functions/lib/infrastructure/ddb"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/guregu/dynamo"
)

const (
	pkPrefix = "Account:"
	sk       = "Account:Account"
)

type (
	// Repository repository.
	Repository struct {
		ddb *dynamo.DB
	}

	// DDBAccount account of ddb.
	DDBAccount struct {
		ddb.SimpleEntry
	}
)

// New new repository
func New(ddb *dynamo.DB) Repository {
	return Repository{
		ddb: ddb,
	}
}

// New new account
func (r Repository) New(ctx context.Context, ac account.Account) error {
	e := fromAccount(ac)
	if err := r.ddb.Table(ddb.Table()).Put(&e).RunWithContext(ctx); err != nil {
		log.Error("put item failed", err)
		return err
	}
	return nil
}

func fromAccount(ac account.Account) DDBAccount {
	return DDBAccount{
		SimpleEntry: ddb.NewSimpleEntry(pkPrefix+strings.ToLower(ac.Address.Hex()), sk, ac.Nonce.String()),
	}
}

// FromAddress get account from eoa
func (r Repository) FromAddress(ctx context.Context, eoa common.Address) (account.Account, error) {
	e := []DDBAccount{}
	err := r.ddb.Table(ddb.Table()).Get("PK", pkPrefix+strings.ToLower(eoa.Hex())).AllWithContext(ctx, &e)
	if err != nil {
		log.Error("query failed", err)
		return account.Account{}, err
	}
	if len(e) == 0 {
		return account.Account{}, nil
	}
	n := new(big.Int)
	n, _ = n.SetString(e[0].Data, 10)
	return account.Account{
		Address: common.HexToAddress(strings.ReplaceAll(e[0].PK, pkPrefix, "")),
		Nonce:   n,
	}, err
}
