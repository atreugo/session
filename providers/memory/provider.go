package memory

import (
	"github.com/fasthttp/session/v2/providers/memory"
)

// Config provider settings.
type Config memory.Config

// New returns a new configured memory provider.
func New(cfg Config) (*memory.Provider, error) {
	provider, err := memory.New(memory.Config(cfg))

	return provider, err
}
