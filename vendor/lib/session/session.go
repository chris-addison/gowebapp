package session

// Session is an object with methods to load and store data for the current user.
// There should be one session per user.
type Session struct {
	id string
}

// NewSession is the constructor for Session
func NewSession(id string) *Session {
	return &Session{
		id: id,
	}
}
