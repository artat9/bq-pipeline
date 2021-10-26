package mediarep

import (
	"context"
	"kaleido-backend/pkg/common/log"
	"kaleido-backend/pkg/infrastructure/ddb"
	"kaleido-backend/pkg/mediaaccount"

	"github.com/ethereum/go-ethereum/common"
	"github.com/guregu/dynamo"
)

const (
	pkPrefix = "MediaApplication:"
	skPrefix = pkPrefix
)

type (
	// Repository repository
	Repository struct{ db *dynamo.DB }
	// DDBApplication application
	DDBApplication struct {
		ddb.SimpleEntry
	}
)

// New new repository
func New(db *dynamo.DB) Repository { return Repository{db: db} }

// NewApplication new application
func (r Repository) NewApplication(ctx context.Context, application mediaaccount.Application) error {
	for _, app := range fromApplication(application) {
		err := r.db.Table(ddb.Table()).Put(&app).RunWithContext(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

// OneWithEOA from eoa
func (r Repository) OneWithEOA(ctx context.Context, eoa common.Address) (mediaaccount.Application, error) {
	apps := []DDBApplication{}
	if err := r.db.Table(ddb.Table()).Get("PK", toPK(eoa)).AllWithContext(ctx, &apps); err != nil {
		log.Error("query failed", err)
		return mediaaccount.Application{}, err
	}
	res := mediaaccount.Application{}
	if len(apps) == 0 {
		return res, nil
	}
	res.Account = eoa
	for _, ap := range apps {
		switch sk := ap.SK; sk {
		case nameSK():
			res.Name = ap.Data
		case mailSK():
			res.MailAddress = ap.Data
		case urlSK():
			res.URL = ap.Data
		}
	}
	return res, nil
}

func fromApplication(ap mediaaccount.Application) []DDBApplication {
	apps := []DDBApplication{}
	if ap.Name != "" {
		apps = append(apps, DDBApplication{SimpleEntry: ddb.NewSimpleEntry(toPK(ap.Account), nameSK(), ap.Name)})
	}
	if ap.MailAddress != "" {
		apps = append(apps, DDBApplication{SimpleEntry: ddb.NewSimpleEntry(toPK(ap.Account), mailSK(), ap.MailAddress)})
	}
	if ap.URL != "" {
		apps = append(apps, DDBApplication{SimpleEntry: ddb.NewSimpleEntry(toPK(ap.Account), urlSK(), ap.URL)})
	}
	return apps
}

func nameSK() string {
	return toSK("Name")
}
func mailSK() string {
	return toSK("MailAddress")
}
func urlSK() string {
	return toSK("URL")
}

func toSK(suf string) string {
	return skPrefix + suf
}

func toPK(eoa common.Address) string {
	return pkPrefix + eoa.Hex()
}
