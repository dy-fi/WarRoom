package repos

// ClientError is an error formatted for the client error template
type ClientError struct {
	Msg string		`json:"msg"`
	Err error		`json:"err"` 
}

// internal critical error
var special = &ClientError{
	Msg: "You shouldn't be getting this message and we're sorry you are.  Give us a moment.",
	Err: nil,
}

// FormatError formats a custom message and an error to be rendered to the error template
func FormatError(m string, e error) ClientError {
	if e == nil {
		return *special
	}

	newErr := &ClientError{
		m,
		e,
	}

	return *newErr
}