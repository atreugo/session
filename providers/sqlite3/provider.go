package sqlite3

import (
	"github.com/atreugo/session"
	"github.com/fasthttp/session/v2/providers/sqlite3"
)

type Config sqlite3.Config

func New(cfg Config) (session.Provider, error) {
	provider, err := sqlite3.New(sqlite3.Config(cfg))

	return provider, err
}

func NewDefaultConfig() Config {
	return Config(sqlite3.NewDefaultConfig())
}

func NewConfigWith(dbPath, tableName string) Config {
	cfg := sqlite3.NewConfigWith(dbPath, tableName)

	return Config(cfg)
}
