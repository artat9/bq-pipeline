package ipfsclient

import (
	"bq-pipeline/pkg/mediaaccount"
	"context"
	"errors"
	"io"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	shell "github.com/ipfs/go-ipfs-api"
)

type (
	fakeClient struct {
		AddFunc func(r io.Reader, options ...shell.AddOpts) (string, error)
		PinFunc func(path string) error
	}
)

func (f fakeClient) Add(r io.Reader, options ...shell.AddOpts) (string, error) {
	return f.AddFunc(r, options...)
}

func (f fakeClient) Pin(path string) error {
	return f.PinFunc(path)
}

func TestService_UploadMedia(t *testing.T) {
	type fields struct {
		client Client
	}
	type args struct {
		info mediaaccount.Application
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "upload succeed",
			fields: fields{
				client: NewIpfsClient().client,
			},
			args: args{
				info: mediaaccount.Application{
					Account:     common.HexToAddress(""),
					Name:        "name",
					URL:         "url",
					Description: "description",
				},
			},
			wantErr: false,
		},
		{
			name: "fail if add failed",
			fields: fields{
				client: fakeClient{
					AddFunc: func(r io.Reader, options ...shell.AddOpts) (string, error) {
						return "", errors.New("err")
					},
					PinFunc: func(path string) error {
						return nil
					},
				},
			},
			args:    args{},
			wantErr: true,
		},
		{
			name: "fail if pin failed",
			fields: fields{
				client: fakeClient{
					AddFunc: func(r io.Reader, options ...shell.AddOpts) (string, error) {
						return "", nil
					},
					PinFunc: func(path string) error {
						return errors.New("err")
					},
				},
			},
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				client: tt.fields.client,
			}
			if _, err := s.UploadMedia(context.Background(), tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("Service.UploadMedia() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
