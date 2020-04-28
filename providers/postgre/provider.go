package postgre

import (
	"github.com/atreugo/session"
	"github.com/fasthttp/session/v2/providers/postgre"
)

type Config postgre.Config

func New(cfg Config) (session.Provider, error) {
	provider, err := postgre.New(postgre.Config(cfg))

	return provider, err
}

func NewDefaultConfig() Config {
	return Config(postgre.NewDefaultConfig())
}

func NewConfigWith(host string, port int64, username string, password string, dbName string, tableName string) Config {
	cfg := postgre.NewConfigWith(host, port, username, password, dbName, tableName)

	return Config(cfg)
}
