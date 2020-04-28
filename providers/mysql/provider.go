package mysql

import (
	"github.com/atreugo/session"
	"github.com/fasthttp/session/v2/providers/mysql"
)

type Config mysql.Config

func New(cfg Config) (session.Provider, error) {
	provider, err := mysql.New(mysql.Config(cfg))

	return provider, err
}

func NewDefaultConfig() Config {
	return Config(mysql.NewDefaultConfig())
}

func NewConfigWith(host string, port int, user, pass, dbName, tableName string) Config {
	cfg := mysql.NewConfigWith(host, port, user, pass, dbName, tableName)

	return Config(cfg)
}
