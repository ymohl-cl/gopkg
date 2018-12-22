package errorx

const (
	separator = " -> "
)

// Errorx take a status code to describe the message
// Implement a wrapper to embeded repport error
type Errorx struct {
	code    uint64
	message string
}

// New return an errorx
func New(code uint64, message string) *Errorx {
	return &Errorx{code: code, message: message}
}

// Error return the content and implement the error standard type
func (e Errorx) Error() string {
	return e.message
}

// Wrap add a content error on the current error
func (e *Errorx) Wrap(message string) {
	e.message = message + separator + e.message
	return
}
