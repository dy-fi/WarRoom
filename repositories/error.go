package repos

import (
	"os"
	"errors"
)

// ClientError is an error formatted for the client error template
type ClientError struct {
	msg string			`json:"msg"`
	err errors.Error	`json:"err"`
}

var special = &ClientError{
	msg: "You shouldn't be getting this message and we're sorry you are.  Give us a moment.",
	err: "",
}

// FormatError formats a custom message and an error to be rendered to error template
func FormatError(m string, e error) ClientError {
	if errors.Is(e, os.ErrExist) {
		m := &ClientError{
			m,
			e.Error,
		}
		return *m
	}
	
	return *special
}