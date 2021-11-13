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
		rep      Repository
		notifier Notifier
	}

	// Application apply
	Application struct {
		MailAddress string
		PublicInfo
	}

	PublicInfo struct {
		Account     common.Address `json:"-"`
		Name        string         `json:"name"`
		URL         string         `json:"url"`
		Description string         `json:"description"`
	}

	// Repository repository
	Repository interface {
		OneWithEOA(ctx context.Context, eoa common.Address) (Application, error)
		NewApplication(ctx context.Context, application Application) error
	}

	// Notifier notifier
	Notifier interface {
		NotifyApplicationCreated(ctx context.Context, application Application) error
	}

	// ApplyForMediaInput input for applying a media ccount
	ApplyForMediaInput struct {
		Name        string `json:"name"`
		MailAddress string `json:"mailAddress"`
		URL         string `json:"url"`
		Description string `json:"description"`
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
func NewService(r Repository, n Notifier) Service {
	return Service{
		rep:      r,
		notifier: n,
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
	newapp := in.newApp(eoa)
	if err = s.rep.NewApplication(ctx, newapp); err != nil {
		return ApplyForMediaOutput{}, err
	}
	return in.toOut(), s.notifier.NotifyApplicationCreated(ctx, newapp)
}

func (in ApplyForMediaInput) newApp(eoa common.Address) Application {
	return Application{
		PublicInfo: PublicInfo{
			Account:     eoa,
			Name:        in.Name,
			URL:         in.URL,
			Description: in.Description,
		},
		MailAddress: in.MailAddress,
	}
}
