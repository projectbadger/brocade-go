package utils

import (
	brocade_errors "github.com/projectbadger/brocade-go/rest/errors"
)

var (
	ErrContentTypeInvalid     = brocade_errors.New("content type is invalid")
	ErrContentTypeNil         = brocade_errors.New("content type is nil")
	ErrNoParsableContentFound = brocade_errors.New("no parsable content found")
)
