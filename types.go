package session

import (
	"github.com/fasthttp/session/v2"
)

// Config configuration of session manager.
type Config session.Config

// Session manages the users sessions.
type Session struct {
	*session.Session
}

// Store represents the user session.
type Store struct {
	*session.Store
}

// Provider interface implemented by providers.
type Provider session.Provider
