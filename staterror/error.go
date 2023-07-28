package staterror

import "strings"

const (
	ErrorCodeAccessTokenInvalid = "access_token_invalid"
	ErrorCodeAccessTokenExpiry  = "access_token_expiry"

	ErrorCodeLinkNotExists = "link_not_exists"
)

// domainError
type domainError string

func (e domainError) Error() string {
	return strings.ReplaceAll(string(e), "_", " ")
}

func (e domainError) ErrorCode() string {
	return string(e)
}

// New
func New(v string) domainError {
	return domainError(v)
}

// notfoudError
type notfoudError struct {
	domainError
}

func (e notfoudError) NotFound() {}

// NewNotFound
func NewNotFound(v string) notfoudError {
	return notfoudError{domainError(v)}
}
