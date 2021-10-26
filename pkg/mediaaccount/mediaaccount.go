package mediaaccount

import (
	"context"
	"errors"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
)

type (
	// Service account service
	Service struct {
		rep Repository
	}

	// Application apply
	Application struct {
		Account     common.Address
		Name        string
		MailAddress string
		URL         string
	}

	// Repository repository
	Repository interface {
		OneWithEOA(ctx context.Context, eoa common.Address) (Application, error)
		NewApplication(ctx context.Context, application Application) error
	}

	// ApplyForMediaInput input for applying a media ccount
	ApplyForMediaInput struct {
		Name        string `json:"name"`
		MailAddress string `json:"mailAddress"`
		URL         string `json:"url"`
	}
	// ApplyForMediaOutput output for applying a media ccount
	ApplyForMediaOutput struct {
		Name        string `json:"name"`
		MailAddress string `json:"mailAddress"`
		URL         string `json:"url"`
	}
)

func (in ApplyForMediaInput) toOut() ApplyForMediaOutput {
	return ApplyForMediaOutput{
		Name:        in.Name,
		MailAddress: in.MailAddress,
		URL:         in.URL,
	}
}

// NewService new service
func NewService(r Repository) Service {
	return Service{
		rep: r,
	}
}

// NewApplication apply for a media account
func (s Service) NewApplication(ctx context.Context, eoa common.Address, in ApplyForMediaInput) (ApplyForMediaOutput, error) {
	ap, err := s.rep.OneWithEOA(ctx, eoa)
	if err != nil {
		return ApplyForMediaOutput{}, err
	}
	if !reflect.DeepEqual(ap, Application{}) {
		return ApplyForMediaOutput{}, errors.New("application already registered")
	}
	return in.toOut(), s.rep.NewApplication(ctx, in.newApp(eoa))
}

func (in ApplyForMediaInput) newApp(eoa common.Address) Application {
	return Application{
		Account:     eoa,
		Name:        in.Name,
		MailAddress: in.MailAddress,
		URL:         in.URL,
	}
}
