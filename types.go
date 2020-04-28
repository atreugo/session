package session

import (
	"github.com/fasthttp/session/v2"
)

// Config config struct
type Config session.Config

// Session session struct
type Session struct {
	*session.Session
}

type Store struct {
	*session.Store
}

// Provider interface implemented by providers
type Provider session.Provider
