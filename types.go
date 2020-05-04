package session

import (
	"github.com/fasthttp/session/v2"
)

// Config configuration of session manager.
type Config session.Config

// Session manages the users sessions.
type Session struct {
	session *session.Session
}

// Provider interface implemented by providers.
type Provider session.Provider
