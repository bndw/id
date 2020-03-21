package id

import (
	"fmt"

	"github.com/docker/docker/pkg/namesgenerator"
	"github.com/google/uuid"
)

const shortLength = 8

func id(length int) string {
	_id := uuid.New().String()
	if length != -1 {
		return _id[:length]
	}

	return _id
}

func ID() string {
	return id(-1)
}

func PrefixID(pre string) string {
	return fmt.Sprintf("%s_%s", pre, id(-1))
}

func ShortID() string {
	return id(shortLength)
}

func ShortPrefixID(pre string) string {
	return fmt.Sprintf("%s_%s", pre, id(shortLength))
}

func Name() string {
	return namesgenerator.GetRandomName(0)
}
