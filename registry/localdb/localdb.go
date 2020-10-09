package localdb

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"os"
)

type localdbregistry struct {}

func New(l *zerolog.Logger) (*localdbregistry, error) {
	var f os.FileInfo
	var err error

	f, err = os.Stat(".cosmic");
	if err != nil {
		if err := os.Mkdir(".cosmic", 0755); err != nil {
			l.Error().Msgf("failed to create new local registry: %v", err)
			return nil, err
		}
		l.Info().Msg("new local registry created")
	}

	f, err = os.Stat(".cosmic")
	if err != nil {
		l.Error().Msgf("failed to find local registry: %v", err)
		return nil, err
	}

	if !f.IsDir() {
		l.Error().Msg("local filename '.cosmic' is reserved and is not a directory")
		return nil, errors.New("local filename '.cosmic' is reserved and is not a directory")
	}
	return &localdbregistry {}, nil
}

func (l *localdbregistry) LoadNamespace(ctx context.Context, ns string) {

}