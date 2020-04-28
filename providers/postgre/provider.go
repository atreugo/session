package postgre

import (
	"github.com/atreugo/session"
	"github.com/fasthttp/session/v2/providers/postgre"
)

// Config provider settings.
type Config postgre.Config

// New returns a new configured postgres provider.
func New(cfg Config) (session.Provider, error) {
	provider, err := postgre.New(postgre.Config(cfg))

	return provider, err
}

// NewDefaultConfig returns a default configuration.
func NewDefaultConfig() Config {
	return Config(postgre.NewDefaultConfig())
}

// NewConfigWith returns a new configuration with especific paremters.
func NewConfigWith(host string, port int64, username string, password string, dbName string, tableName string) Config {
	cfg := postgre.NewConfigWith(host, port, username, password, dbName, tableName)

	return Config(cfg)
}
