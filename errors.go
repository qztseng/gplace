package gplace

import "fmt"

// ErrMissingAPIKey indicates a missing API key.
var ErrMissingAPIKey = fmt.Errorf("gplace: missing api key")

// ValidationError describes an invalid request payload.
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("gplace: invalid %s: %s", e.Field, e.Message)
}

// APIError represents an HTTP error from the Places API.
type APIError struct {
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	if e.Body == "" {
		return fmt.Sprintf("gplace: api error (%d)", e.StatusCode)
	}
	return fmt.Sprintf("gplace: api error (%d): %s", e.StatusCode, e.Body)
}
