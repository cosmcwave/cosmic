package backend

import (
	"cosmic/config"
	"cosmic/registry"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"time"
)

type Backend struct {
	httpServer *http.Server
	registry   registry.Registry
	logger     *zerolog.Logger
}

func New(conf *config.Config) (*Backend, error) {
	r, err := registry.New(registry.RegistryType(conf.Registry))
	if err != nil {
		return nil, err
	}

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	l := zerolog.New(output).With().Timestamp().Logger()

	b := &Backend{
		httpServer: &http.Server{
			Addr:    conf.ServerAddress,
			Handler: routesFactory(),
		},
		registry: r,
		logger:   &l,
	}
	return b, nil
}

func (b *Backend) Start() error {
	b.logger.Info().Msg("cosmic engine started")
	return b.httpServer.ListenAndServe()
}

func routesFactory() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/namespace/{name}", namespaceHandler).Methods("GET")
	router.HandleFunc("/namespace/{name}", namespaceHandler).Methods("GET")
	router.HandleFunc("/service/{name}", serviceHandler).Methods("PUT")
	router.HandleFunc("/service/{name}", serviceHandler).Methods("PUT")
	return router
}
