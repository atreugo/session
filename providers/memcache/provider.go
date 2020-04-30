package memcache

import (
	"github.com/fasthttp/session/v2/providers/memcache"
)

// Config provider settings.
type Config memcache.Config

// New returns a new memcache provider configured.
func New(cfg Config) (*memcache.Provider, error) {
	provider, err := memcache.New(memcache.Config(cfg))

	return provider, err
}
