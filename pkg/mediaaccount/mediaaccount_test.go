package mediaaccount

import (
	"context"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestApplyForMediaInput_toOut(t *testing.T) {
	type fields struct {
		Name        string
		MailAddress string
		URL         string
	}
	tests := []struct {
		name   string
		fields fields
		want   ApplyForMediaOutput
	}{
		{
			name:   "to output success",
			fields: fields{Name: "name", MailAddress: "test@example.com", URL: "https://auroradao.org"},
			want:   ApplyForMediaOutput{Name: "name", MailAddress: "test@example.com", URL: "https://auroradao.org"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := ApplyForMediaInput{
				Name:        tt.fields.Name,
				MailAddress: tt.fields.MailAddress,
				URL:         tt.fields.URL,
			}
			if got := in.toOut(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApplyForMediaInput.toOut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_NewApplication(t *testing.T) {
	type fields struct {
		rep Repository
	}
	type args struct {
		ctx context.Context
		eoa common.Address
		in  ApplyForMediaInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    ApplyForMediaOutput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				rep: tt.fields.rep,
			}
			got, err := s.NewApplication(tt.args.ctx, tt.args.eoa, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.NewApplication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.NewApplication() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApplyForMediaInput_newApp(t *testing.T) {
	type fields struct {
		Name        string
		MailAddress string
		URL         string
	}
	type args struct {
		eoa common.Address
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Application
	}{
		{
			name: "Name",
			fields: fields{
				Name:        "Name",
				MailAddress: "test@example.com",
				URL:         "https://auroradao.org",
			},
			args: args{
				eoa: common.HexToAddress("0x8dc81f896b38167734ca4ff26b1d20c4c78e9190"),
			},
			want: Application{
				Account:     common.HexToAddress("0x8dc81f896b38167734ca4ff26b1d20c4c78e9190"),
				Name:        "Name",
				MailAddress: "test@example.com",
				URL:         "https://auroradao.org",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := ApplyForMediaInput{
				Name:        tt.fields.Name,
				MailAddress: tt.fields.MailAddress,
				URL:         tt.fields.URL,
			}
			if got := in.newApp(tt.args.eoa); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApplyForMediaInput.newApp() = %v, want %v", got, tt.want)
			}
		})
	}
}
