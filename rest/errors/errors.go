package errors

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrUnmarshalFailed = errors.New("error unmarshal failed")
)

type BrocadeErr interface {
	Error() string
	GetErrors() *Errors
	AddError(Error)
}

func NewFromErr(err error) BrocadeErr {
	if err == nil {
		return nil
	}
	return &BrocadeErrorImpl{
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

func NewFromErrors(err ...Error) BrocadeErr {
	if err == nil {
		return nil
	}
	return &BrocadeErrorImpl{
		Errors: err,
	}
}

type BrocadeErrorImpl struct {
	XMLName xml.Name `json:"-" xml:"errors"`
	Errors  []Error  `json:"error" xml:"error"`
}

func (err *BrocadeErrorImpl) Error() string {
	var errStr string
	if err == nil {
		return ""
	}
	for _, brocadeErr := range err.Errors {
		errStr += strconv.Itoa(brocadeErr.ErrorInfo.ErrorCode) + ": " + brocadeErr.ErrorMessage + "; "
	}
	return strings.Trim(errStr, "; ")
}

func (err *BrocadeErrorImpl) GetErrors() *Errors {
	if err == nil {
		return nil
	}
	return &Errors{
		Errors: err.Errors,
	}
}

func (err *BrocadeErrorImpl) AddError(newErr Error) {
	if err != nil {
		err.Errors = append(err.Errors, newErr)
		return
	}
	err = &BrocadeErrorImpl{
		Errors: []Error{
			newErr,
		},
	}
}

type Errors struct {
	XMLName xml.Name `json:"-" xml:"errors"`
	Errors  []Error  `json:"error" xml:"error"`
}

func (e *Errors) Error() string {
	if e == nil {
		return ""
	}
	return e.String()
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

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return e.String()
}

func New(message string, args ...interface{}) BrocadeErr {
	if message == "" {
		return nil
	}
	return &BrocadeErrorImpl{
		Errors: []Error{
			{
				ErrorType:    "error",
				ErrorMessage: fmt.Sprintf(message, args...),
			},
		},
	}
}
