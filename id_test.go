package id_test

import (
	"strings"
	"testing"

	"github.com/bndw/id"
)

const duplicateIterations = 1000000

func TestShort(t *testing.T) {
	var seen = make(map[string]interface{})
	for i := 0; i < duplicateIterations; i++ {
		x := id.Short()
		if _, ok := seen[x]; ok {
			t.Fatalf("Short generated duplicate id=%s after %d iterations", x, i)
		}
		if i == 0 {
			t.Log(x)
		}
		seen[x] = nil
	}
}

func TestShortPrefix(t *testing.T) {
	const prefix = "foo_"
	var seen = make(map[string]interface{})
	for i := 0; i < duplicateIterations; i++ {
		x := id.ShortPrefix(prefix)
		if !strings.HasPrefix(x, prefix) {
			t.Fatalf("Expected ID to be prefixed with %s, got %s", prefix, x)
		}
		if _, ok := seen[x]; ok {
			t.Fatalf("ShortPrefix generated duplicate id=%s after %d iterations", x, i)
		}
		if i == 0 {
			t.Log(x)
		}
		seen[x] = nil
	}
}

func TestUUID(t *testing.T) {
	var seen = make(map[string]interface{})
	for i := 0; i < duplicateIterations; i++ {
		x := id.UUID()
		if _, ok := seen[x]; ok {
			t.Fatalf("UUID generated duplicate id=%s after %d iterations", x, i)
		}
		if i == 0 {
			t.Log(x)
		}
		seen[x] = nil
	}
}

func TestUUIDPrefix(t *testing.T) {
	const prefix = "foo_"
	var seen = make(map[string]interface{})
	for i := 0; i < duplicateIterations; i++ {
		x := id.UUIDPrefix(prefix)
		if !strings.HasPrefix(x, prefix) {
			t.Fatalf("Expected ID to be prefixed with %s, got %s", prefix, x)
		}
		if _, ok := seen[x]; ok {
			t.Fatalf("UUIDPrefix generated duplicate id=%s after %d iterations", x, i)
		}
		if i == 0 {
			t.Log(x)
		}
		seen[x] = nil
	}
}

func TestName(t *testing.T) {
	for i := 0; i < 100; i++ {
		if x := id.Name(); x == "" {
			t.Fatal("Expected name, got an empty string")
		} else {
			if i == 0 {
				t.Log(x)
			}
		}
	}
}

func TestNamePrefix(t *testing.T) {
	const prefix = "foo_"
	for i := 0; i < 100; i++ {
		if x := id.NamePrefix(prefix); x == "" {
			t.Fatal("Expected name, got an empty string")
		} else {
			if !strings.HasPrefix(x, prefix) {
				t.Fatalf("Expected ID to be prefixed with %s, got %s", prefix, x)
			}
			if i == 0 {
				t.Log(x)
			}
		}
	}
}
