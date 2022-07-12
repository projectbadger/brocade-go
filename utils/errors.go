package utils

import (
	brocade_errors "brocade/rest/errors"
)

var (
	ErrContentTypeInvalid     = brocade_errors.New("content type is invalid")
	ErrContentTypeNil         = brocade_errors.New("content type is nil")
	ErrNoParsableContentFound = brocade_errors.New("no parsable content found")
)
