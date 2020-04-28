package memcache

import (
	"github.com/atreugo/session"
	"github.com/fasthttp/session/v2/providers/memcache"
)

type Config memcache.Config

func New(cfg Config) (session.Provider, error) {
	provider, err := memcache.New(memcache.Config(cfg))

	return provider, err
}
