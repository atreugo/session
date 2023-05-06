package mongodb

import (
	"github.com/fasthttp/session/v2/providers/mongodb"
)

// Config provider settings.
type Config mongodb.Config

// NewConfigWith returns a new configuration with especific paremters.
func NewConfigWith(connectionURL string, dbName string, collection string) Config {
	cfg := mongodb.NewConfigWith(connectionURL, dbName, collection)

	return Config(cfg)
}

// New returns a new configured postgres provider.
func New(cfg Config) (*mongodb.Provider, error) {
	provider, err := mongodb.New(mongodb.Config(cfg))

	return provider, err
}
