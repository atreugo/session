package redis

import (
	"github.com/atreugo/session"
	"github.com/fasthttp/session/v2/providers/redis"
)

// Config provider settings.
type Config redis.Config

// New returns a new configured redis provider.
func New(cfg Config) (session.Provider, error) {
	provider, err := redis.New(redis.Config(cfg))

	return provider, err
}
