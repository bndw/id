package id_test

import (
	"testing"

	"github.com/bndw/id"
)

func TestID(t *testing.T) {
	t.Log("ID: ", id.ID())
	t.Log("IDWithPrefix: ", id.IDWithPrefix("stuff"))
	t.Log("ShortID: ", id.ShortID())
	t.Log("ShortIDWithPrefix: ", id.ShortIDWithPrefix("stuff"))
	t.Log("Name: ", id.Name())
}
