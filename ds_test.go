package ds_test

import (
	"context"
	"os"
	"testing"

	cds "cloud.google.com/go/datastore"
	"go.mercari.io/datastore"
	"go.mercari.io/datastore/clouddatastore"
)

type Hoge struct {
	ID   string `datastore:"-"`
	Body string
}

func TestGetMulti(t *testing.T) {
	if err := os.Setenv("DATASTORE_EMULATOR_HOST", "localhost:8081"); err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	client, err := cds.NewClient(ctx, "hoge")
	if err != nil {
		t.Fatal(err)
	}
	ds, err := clouddatastore.FromClient(ctx, client)
	if err != nil {
		t.Fatal(err)
	}

	keys := []datastore.Key{ds.NameKey("Hoge", "hoge1", nil)}
	vs := []*Hoge{
		{
			Body: "Hogeeeee",
		},
	}

	_, err = ds.PutMulti(ctx, keys, vs)
	if err != nil {
		t.Fatal(err)
	}

	rvs := make([]*Hoge, len(keys))
	if err := ds.GetMulti(ctx, keys, rvs); err != nil {
		t.Fatal(err)
	}
}
