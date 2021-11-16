package mediarep

import (
	"context"
	"kaleido-backend/pkg/common/log"
	"kaleido-backend/pkg/infrastructure/ddb"
	"kaleido-backend/pkg/mediaaccount"
	"strconv"

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

func (r Repository) OneWithEOAAndAppliedAt(ctx context.Context, eoa common.Address, timestamp string) (mediaaccount.Application, error) {
	apps := []DDBApplication{}
	if err := r.db.Table(ddb.Table()).Get("PK", toPK(eoa, timestamp)).AllWithContext(ctx, &apps); err != nil {
		log.Error("query failed", err)
		return mediaaccount.Application{}, err
	}
	res := mediaaccount.Application{}
	if len(apps) == 0 {
		return res, nil
	}
	res.Account = eoa
	res.AppliedAt = timestamp
	for _, ap := range apps {
		d := ap.Data
		switch sk := ap.SK; sk {
		case nameSK():
			res.Name = d
		case mailSK():
			res.MailAddress = d
		case urlSK():
			res.URL = d
		case descriptionSK():
			res.Description = d
		case documentURLSK():
			res.URL = d
		case pvMonthSK():
			i, _ := strconv.Atoi(d)
			res.PVMonth = i
		case primaryCustomersSK():
			res.PrimaryCustomers = d
		}
	}
	return res, nil
}

func fromApplication(ap mediaaccount.Application) []DDBApplication {
	apps := []DDBApplication{}
	pk := toPK(ap.Account, ap.AppliedAt)
	entries := map[string]string{
		nameSK():             ap.Name,
		mailSK():             ap.MailAddress,
		urlSK():              ap.URL,
		descriptionSK():      ap.Description,
		documentURLSK():      ap.DocumentURL,
		pvMonthSK():          strconv.Itoa(ap.PVMonth),
		primaryCustomersSK(): ap.PrimaryCustomers,
	}
	for k, v := range entries {
		apps = append(apps, DDBApplication{ddb.NewSimpleEntry(pk, k, v)})
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

func descriptionSK() string {
	return toSK("Description")
}

func documentURLSK() string {
	return toSK("DocumentURL")
}

func pvMonthSK() string {
	return toSK("PVMonth")
}

func primaryCustomersSK() string {
	return toSK("PrimaryCustomers")
}

func toSK(suf string) string {
	return skPrefix + suf
}

func toPK(eoa common.Address, at string) string {
	return pkPrefix + eoa.Hex() + ":" + at
}
