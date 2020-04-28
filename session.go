package session

import (
	"github.com/fasthttp/session/v2"
	"github.com/savsgio/atreugo/v11"
)

func New(cfg Config) *Session {
	s := session.New(session.Config(cfg))

	return &Session{Session: s}
}

func NewDefaultConfig() Config {
	return Config(session.NewDefaultConfig())
}

func (s *Session) SetProvider(provider Provider) error {
	return s.Session.SetProvider(provider)
}

// TODO: Use pool
func (s *Session) Get(ctx *atreugo.RequestCtx) (*Store, error) {
	store, err := s.Session.Get(ctx.RequestCtx)

	return &Store{Store: store}, err
}

func (s *Session) Save(ctx *atreugo.RequestCtx, store *Store) error {
	return s.Session.Save(ctx.RequestCtx, store.Store)
}

func (s *Session) Regenerate(ctx *atreugo.RequestCtx) error {
	return s.Session.Regenerate(ctx.RequestCtx)
}

func (s *Session) Destroy(ctx *atreugo.RequestCtx) error {
	return s.Session.Destroy(ctx.RequestCtx)
}
