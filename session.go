package session

import (
	"github.com/fasthttp/session/v2"
	"github.com/savsgio/atreugo/v11"
)

// New returns a configured manager.
func New(cfg Config) *Session {
	s := session.New(session.Config(cfg))

	return &Session{Session: s}
}

// NewDefaultConfig returns a new default configuration.
func NewDefaultConfig() Config {
	return Config(session.NewDefaultConfig())
}

// SetProvider sets the session provider used by the sessions manager.
func (s *Session) SetProvider(provider Provider) error {
	return s.Session.SetProvider(provider)
}

// Get returns the user session
// if it does not exist, it will be generated.
func (s *Session) Get(ctx *atreugo.RequestCtx) (*Store, error) {
	store, err := s.Session.Get(ctx.RequestCtx)

	return &Store{Store: store}, err
}

// Save saves the user session
//
// Warning: Don't use the store after exec this function, because, you will lose the after data
// For avoid it, defer this function in your request handler.
func (s *Session) Save(ctx *atreugo.RequestCtx, store *Store) error {
	return s.Session.Save(ctx.RequestCtx, store.Store)
}

// Regenerate generates a new session id to the current user.
func (s *Session) Regenerate(ctx *atreugo.RequestCtx) error {
	return s.Session.Regenerate(ctx.RequestCtx)
}

// Destroy destroys the session of the current user.
func (s *Session) Destroy(ctx *atreugo.RequestCtx) error {
	return s.Session.Destroy(ctx.RequestCtx)
}
