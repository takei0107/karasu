package httperror

import (
	"errors"
	"fmt"
	"io/fs"
	"net/http"
)

type httpError struct {
	err        error
	statusCode int
	message    string
}

func New(e error) httpError {
	var statusCode int
	switch {
	case errors.Is(e, fs.ErrNotExist):
		statusCode = http.StatusNotFound
	default:
		statusCode = http.StatusInternalServerError
	}
	return httpError{e, statusCode, http.StatusText(statusCode)}
}

func (e httpError) StatusCode() int {
	return e.statusCode
}

func (e httpError) Message() string {
	return e.message
}

func (e httpError) Error() string {
	return fmt.Sprint(fmt.Errorf("statusCode = %d, message = %s, detail = [%w]", e.statusCode, e.message, e.err))
}
