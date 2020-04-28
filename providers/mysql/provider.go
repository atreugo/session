package mysql

import (
	"github.com/atreugo/session"
	"github.com/fasthttp/session/v2/providers/mysql"
)

// Config provider settings
type Config mysql.Config

// New returns a new configured mysql provider.
func New(cfg Config) (session.Provider, error) {
	provider, err := mysql.New(mysql.Config(cfg))

	return provider, err
}

// NewDefaultConfig returns a default configuration.
func NewDefaultConfig() Config {
	return Config(mysql.NewDefaultConfig())
}

// NewConfigWith returns a new configuration with especific paremters.
func NewConfigWith(host string, port int, user, pass, dbName, tableName string) Config {
	cfg := mysql.NewConfigWith(host, port, user, pass, dbName, tableName)

	return Config(cfg)
}
