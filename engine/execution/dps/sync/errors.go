package sync

import "errors"

// Sentinels used by the server to handle sync request errors.
var (
	ErrBadRequest = errors.New("bad request")
	ErrInternal = errors.New("internal server error")
)
