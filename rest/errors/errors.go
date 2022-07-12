package errors

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
)

var (
	ErrUnmarshalFailed = errors.New("error unmarshal failed")
)

// BrocadeErr represents an error that may hold brocade
// error format.
type BrocadeErr interface {
	Error() string
	GetErrors() *Errors
	AddError(Error)
}

// NewFromErr returns a new BrocadeErr from an error
func NewFromErr(err error) BrocadeErr {
	if err == nil {
		return nil
	}
	return &Errors{
		Errors: []Error{
			{
				ErrorType:    "error",
				ErrorMessage: err.Error(),
				ErrorInfo: ErrorInfo{
					ErrorCode: 500,
				},
			},
		},
	}
}

// NewFromErr returns a new BrocadeErr from a number of
// Error-s
func NewFromErrors(err ...Error) BrocadeErr {
	if err == nil {
		return nil
	}
	return &Errors{
		Errors: err,
	}
}

// Errors holds a slice []Error
type Errors struct {
	XMLName xml.Name `json:"-" xml:"errors"`
	Errors  []Error  `json:"error" xml:"error"`
}

// Error returns a string, constructed from the contained
// errors
func (e *Errors) Error() string {
	if e == nil {
		return ""
	}
	return e.String()
}

// GetErrors returns Errors
func (e *Errors) GetErrors() *Errors {
	return e
}

// AddError appends a new Error
func (err *Errors) AddError(newErr Error) {
	if err != nil {
		err.Errors = append(err.Errors, newErr)
		return
	}
	err = &Errors{
		Errors: []Error{
			newErr,
		},
	}
}

func (e *Errors) String() string {
	if e == nil {
		return ""
	}
	bytes, err := json.Marshal(e)
	if err == nil {
		return ""
	}
	return string(bytes)
}

// Error represents a Brocade REST API error
type Error struct {
	XMLName      xml.Name  `json:"-" xml:"error"`
	ErrorType    string    `json:"error-type" xml:"error-type"`
	ErrorTTag    string    `json:"error-tag" xml:"error-tag"`
	ErrorAppTag  string    `json:"error-app-tag,omitempty" xml:"error-app-tag"`
	ErrorPath    string    `json:"error-path,omitempty" xml:"error-path"`
	ErrorMessage string    `json:"error-message,omitempty" xml:"error-message"`
	ErrorInfo    ErrorInfo `json:"error-info,omitempty" xml:"error-info"`
}

type ErrorInfo struct {
	XMLName     xml.Name `json:"-" xml:"error-info"`
	ErrorCode   int      `json:"error-code" xml:"error-code"`
	ErrorModule string   `json:"error-module" xml:"error-module"`
}

// String returns JSON-marshaled Error
func (e *Error) String() string {
	if e == nil {
		return ""
	}
	bytes, err := json.Marshal(e)
	if err == nil {
		return ""
	}
	return string(bytes)
}

// Error returns JSON-marshaled Error
func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return e.String()
}

// New returns a new BrocadeErr implementation
func New(message string, args ...interface{}) BrocadeErr {
	if message == "" {
		return nil
	}
	return &Errors{
		Errors: []Error{
			{
				ErrorType:    "error",
				ErrorMessage: fmt.Sprintf(message, args...),
			},
		},
	}
}
