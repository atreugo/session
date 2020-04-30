package sqlite3

import (
	"github.com/fasthttp/session/v2/providers/sqlite3"
)

// Config provider settings.
type Config sqlite3.Config

// New returns a new configured sqlite3 provider.
func New(cfg Config) (*sqlite3.Provider, error) {
	provider, err := sqlite3.New(sqlite3.Config(cfg))

	return provider, err
}

// NewDefaultConfig returns a default configuration.
func NewDefaultConfig() Config {
	return Config(sqlite3.NewDefaultConfig())
}

// NewConfigWith returns a new configuration with especific paremters.
func NewConfigWith(dbPath, tableName string) Config {
	cfg := sqlite3.NewConfigWith(dbPath, tableName)

	return Config(cfg)
}
