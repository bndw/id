package id

import (
	"fmt"

	"github.com/docker/docker/pkg/namesgenerator"
	"github.com/google/uuid"
	"github.com/teris-io/shortid"
)

// Short returns a short identifier.
// It's a wrapper around https://github.com/teris-io/shortid
func Short() string {
	return shortid.MustGenerate()
}

// ShortPrefix return a short identifier prefixed with the given prefix.
func ShortPrefix(pre string) string {
	return fmt.Sprintf("%s%s", pre, Short())
}

// UUID returns a version 4 UUID.
// It's a wrapper around https://godoc.org/github.com/google/uuid#UUID
func UUID() string {
	return uuid.New().String()
}

// UUIDPrefix return a version 4 UUID prefixed with the given prefix.
func UUIDPrefix(pre string) string {
	return fmt.Sprintf("%s%s", pre, UUID())
}

// Name returns a Docker-style name.
// It's a wrapper around https://github.com/moby/moby/tree/master/pkg/namesgenerator
func Name() string {
	return namesgenerator.GetRandomName(0)
}

// NamePrefix returns a Docker-style name with the given prefix.
func NamePrefix(pre string) string {
	return fmt.Sprintf("%s%s", pre, Name())
}
