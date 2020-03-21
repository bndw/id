package id_test

import (
	"testing"

	"github.com/bndw/id"
)

func TestID(t *testing.T) {
	t.Log("ID: ", id.ID())
	t.Log("IDWithPrefix: ", id.PrefixID("stuff"))
	t.Log("ShortID: ", id.ShortID())
	t.Log("ShortIDWithPrefix: ", id.ShortPrefixID("stuff"))
	t.Log("Name: ", id.Name())
}
