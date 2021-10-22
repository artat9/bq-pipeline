package ddb

import (
	"context"
	"errors"
	"kaleido-backend/lib/functions/lib/common/log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/guregu/dynamo"
)

// New New Dynamodb client
func New() *dynamo.DB {
	ddb := dynamo.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
	client := ddb.Client().(*dynamodb.DynamoDB).Client
	xray.AWS(client)
	return ddb
}

// Table Resolve table name
func Table() string {
	return "kaleido-backend-main-" + os.Getenv("EnvironmentId")
}

// SimpleEntry simple entry
type SimpleEntry struct {
	PK   string `dynamo:"PK"`
	SK   string `dynamo:"SK"`
	Data string `dynamo:"Data"`
}

//NewSimpleEntry new SimpleEntry
func NewSimpleEntry(id, sk, data string) SimpleEntry {
	return SimpleEntry{
		PK:   id,
		SK:   sk,
		Data: data,
	}
}

// ResolveSimpleEntry resolve SimpleEntry from DDB.
func ResolveSimpleEntry(ctx context.Context, pk, id string, tbl dynamo.Table) (SimpleEntry, error) {
	entry := SimpleEntry{}
	err := tbl.Get("PK", pk).Range("SK", dynamo.Equal, id).OneWithContext(ctx, &entry)
	if err != nil {
		log.Error("", err)
		return entry, errors.New("entry not found. id:" + id)
	}
	return entry, nil
}

// ResolveSimpleEntries resolve SimpleEntry slice from DDB.
func ResolveSimpleEntries(ctx context.Context, pk string, ids []string, tbl dynamo.Table) ([]SimpleEntry, error) {
	entries := []SimpleEntry{}
	for _, id := range ids {
		ent, err := ResolveSimpleEntry(ctx, pk, id, tbl)
		if err != nil {
			return entries, err
		}
		entries = append(entries, ent)
	}
	return entries, nil
}
