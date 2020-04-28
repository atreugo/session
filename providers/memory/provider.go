package memory

import (
	"github.com/atreugo/session"
	"github.com/fasthttp/session/v2/providers/memory"
)

type Config memory.Config

func New(cfg Config) (session.Provider, error) {
	provider, err := memory.New(memory.Config(cfg))

	return provider, err
}
