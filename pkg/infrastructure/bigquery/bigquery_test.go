package bigquery

import (
	"bq-pipeline/pkg/infrastructure/ssm"
	"bq-pipeline/pkg/user"
	"context"
	"errors"
	"testing"

	"cloud.google.com/go/bigquery"
)

// // import (
// // 	"testing"
// // )

// func TestNewClient(t *testing.T) {
// 	t.Run("connect succeed", func(t *testing.T) {
// 		_, err := New(context.Background())
// 		if err != nil {
// 			t.Error(err)
// 		}
// 	})
// }
// func TesCredentialFileWrite
func TestNewClient(t *testing.T) {
	t.Run("connect succeded", func(t *testing.T) {
		_, err := NewClient(context.Background(), ssm.New())
		if err != nil {
			t.Error(err)
		}
	})
}

type (

	testInserter struct {
	  PutFunc func(ctx context.Context,src interface{}) error
	}
	testReader struct {
		ReadFunc func(ctx context.Context)(*bigquery.RowIterator, error)
	}
	errorInserter struct {
		testInserter testInserter
		testReader testReader
	}
	normalService struct {
		testInserter testInserter
		testReader testReader
	}
)

	func (t testInserter) Put(ctx context.Context, src interface{}) error {
		return t.PutFunc(ctx, src)
	}
	func (t testReader) Read(ctx context.Context)(*bigquery.RowIterator, error){
		return t.ReadFunc(ctx)
	}

	func (n normalService) Put(ctx context.Context, users user.User) error {
		return nil
	}

	func (n normalService) Reader(ctx context.Context)(*bigquery.RowIterator, error){
		return nil, nil
	}

func TestUpload(t *testing.T) {
	testuser := user.User{ID: "test", Name: "test"}
	testUsers := []user.User{}
	testUsers = append(testUsers, testuser)

	type fields struct {
   si SampleInserter
	}
	type args struct {
		users []user.User
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
			  si: testInserter{
					PutFunc: func(ctx context.Context, src interface{}) error {
						return nil
					},
				},
		  },
		  args: args{
		  	users: testUsers,
	    },
			wantErr:  false,
    },
		{
		  name: "upload failed",
		  fields: fields{
				  si: testInserter{
					  PutFunc: func(ctx context.Context, src interface{})error {
						  return errors.New("err")
					  },
			  },
		  },
		  args: args{
		  	users:
				  testUsers,
	      },
			wantErr: true,
    },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			s := Service{
				si: tt.fields.si,
				sr: nil,
			}
			if err := s.Upload(context.Background(), tt.args.users); (err != nil) != tt.wantErr {
				t.Errorf("Service.Put() error  = %v, wantErr %v", err, tt.wantErr)
			}
		},
  )}
}

func TestRead(t *testing.T) {
	type fields struct {
   sr SampleReader
	}
	type args struct {
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
		  name: "read succeed",
		  fields: fields{
			  sr: testReader{
					ReadFunc: func(ctx context.Context)(*bigquery.RowIterator, error) {
						return nil, nil
					},
				},
		  },
			wantErr:  false,
    },
		{
		  name: "read failed",
		  fields: fields{
				sr: testReader{
					ReadFunc: func(ctx context.Context)(*bigquery.RowIterator, error) {
						return nil, errors.New("err")
					},
			  },
		  },
			wantErr: true,
    },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			s := Service{
				si: nil,
				sr: tt.fields.sr,
			}
			if _, err := s.Download(context.Background()); (err != nil) != tt.wantErr {
				t.Errorf("Service.Put() error  = %v, wantErr %v", err, tt.wantErr)
			}
		},
  )}
}
