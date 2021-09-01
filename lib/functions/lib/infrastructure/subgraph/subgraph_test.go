package subgraph

import (
	"context"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("enable to initialize client", func(t *testing.T) {
		c := New()
		if c.svc == nil {
			t.Error("client not initialized")
		}
	})
}

func TestMetadataURI(t *testing.T) {
	cli := New()
	want := "https://arweave.net/1yOme2mzsjq-Q0WwFIVtoDi_zctVZcFOE36zn0lY3hQ"
	t.Run("enable to get MetadataURI from graph in case PostID exists", func(t *testing.T) {
		got, err := cli.MetadataURI(context.Background(), "0x989681")
		if err != nil {
			t.Error("unexpected error occured", err)
		}
		if got != "https://arweave.net/1yOme2mzsjq-Q0WwFIVtoDi_zctVZcFOE36zn0lY3hQ" {
			t.Error("uri unexpected. want:", want, ", got:", got)
		}
	})
	t.Run("error if PostID not found", func(t *testing.T) {
		_, err := cli.MetadataURI(context.Background(), "test")
		if err == nil {
			t.Error("no error occured")
		}
	})
}
