package session

import (
	"sync"
	"time"
)

// Session is an object with methods to load and store data for the current user.
// There should be one session per user.
type Session struct {
	Lock     sync.Mutex
	id       string
	lastUsed time.Time
	values   map[string]interface{}
}

// NewSession is the constructor for Session
func NewSession(id string) *Session {
	return &Session{
		id:       id,
		lastUsed: time.Now(),
		values:   make(map[string]interface{}),
	}
}

// Read returns value associated with the given key
func (session *Session) Read(key string) interface{} {
	session.lastUsed = time.Now()
	return session.values[key]
}

// Write associates the given key and value
func (session *Session) Write(key string, value interface{}) {
	session.lastUsed = time.Now()
	session.values[key] = value
}

// Delete removes the given key
func (session *Session) Delete(key string) {
	session.lastUsed = time.Now()
	delete(session.values, key)
}
