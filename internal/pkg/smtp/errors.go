package smtp

import "errors"

var (
	BadEmail = errors.New("email has incorrect format")
)
