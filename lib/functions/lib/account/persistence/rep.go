package accountrep

import (
	"context"
	"kaleido-backend/lib/functions/lib/account"
	"kaleido-backend/lib/functions/lib/common/log"
	"kaleido-backend/lib/functions/lib/infrastructure/ddb"

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
		SimpleEntry: ddb.NewSimpleEntry(pkPrefix+ac.Address.Hex(), sk, ac.Nonce.String()),
	}
}

func (r Repository) FromAddress(ctx context.Context, eoa common.Address) (ac account.Account, err error) {
	return account.Account{}, nil
}
