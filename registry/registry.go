package registry

import (
	"context"
	"cosmic/registry/localdb"
	"github.com/rs/zerolog"
	"os"
	"time"
)

type RegistryType string

const (
	LocalRegistry RegistryType = "localdb"
)

type Registry interface {
	LoadNamespace(ctx context.Context, ns string)
}

func New(t RegistryType) (Registry, error) {
	return registryFactory(t)
}

func registryFactory(t RegistryType) (Registry, error) {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	l := zerolog.New(output).With().Timestamp().Logger()

	switch t {
	case LocalRegistry:
		return localdb.New(&l)
	}
	return localdb.New(&l)
}