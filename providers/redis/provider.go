package redis

import (
	"github.com/atreugo/session"
	"github.com/fasthttp/session/v2/providers/redis"
)

type Config redis.Config

func New(cfg Config) (session.Provider, error) {
	provider, err := redis.New(redis.Config(cfg))

	return provider, err
}
