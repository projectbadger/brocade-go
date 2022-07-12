
# errors

```go
import github.com/projectbadger/brocade-go/rest/errors
```

## Index

- [type BrocadeErr](#type-brocadeerr)
  - [New(string) BrocadeErr](#func-newstring-brocadeerr)
  - [NewFromErr(error) BrocadeErr](#func-newfromerrerror-brocadeerr)
  - [NewFromErrors() BrocadeErr](#func-newfromerrors-brocadeerr)
- [type Error](#type-error)
  - [Error() string](#func-error-error-string)
  - [String() string](#func-error-string-string)
- [type ErrorInfo](#type-errorinfo)
- [type Errors](#type-errors)
  - [AddError(Error)](#func-errors-adderrorerror)
  - [Error() string](#func-errors-error-string)
  - [GetErrors() Errors](#func-errors-geterrors-errors)
  - [String() string](#func-errors-string-string)
- [Variables](#variables)

## Variables
```go
var (
	ErrUnmarshalFailed = errors.New("error unmarshal failed")
)

```


## type [BrocadeErr](<errors.go#L16>)

BrocadeErr represents an error that may hold brocade
error format.
```go
type BrocadeErr interface {
	Error() string
	GetErrors() *Errors
	AddError(Error)
}
```

## func [New(string) BrocadeErr](<errors.go#L133>)

New returns a new BrocadeErr implementation


```go
func New(message string, args ...interface{}) BrocadeErr
```
## func [NewFromErr(error) BrocadeErr](<errors.go#L23>)

NewFromErr returns a new BrocadeErr from an error


```go
func NewFromErr(err error) BrocadeErr
```
## func [NewFromErrors() BrocadeErr](<errors.go#L42>)

NewFromErr returns a new BrocadeErr from a number of
Error-s


```go
func NewFromErrors(err ...Error) BrocadeErr
```

## type [Error](<errors.go#L96>)

Error represents a Brocade REST API error
```go
type Error struct {
	XMLName		xml.Name	`json:"-" xml:"error"`
	ErrorType	string		`json:"error-type" xml:"error-type"`
	ErrorTTag	string		`json:"error-tag" xml:"error-tag"`
	ErrorAppTag	string		`json:"error-app-tag,omitempty" xml:"error-app-tag"`
	ErrorPath	string		`json:"error-path,omitempty" xml:"error-path"`
	ErrorMessage	string		`json:"error-message,omitempty" xml:"error-message"`
	ErrorInfo	ErrorInfo	`json:"error-info,omitempty" xml:"error-info"`
}
```

## func (*Error) [Error() string](<errors.go#L125>)

Error returns JSON-marshaled Error


```go
func (e *Error) Error() string
```
## func (*Error) [String() string](<errors.go#L113>)

String returns JSON-marshaled Error


```go
func (e *Error) String() string
```

## type [ErrorInfo](<errors.go#L106>)
```go
type ErrorInfo struct {
	XMLName		xml.Name	`json:"-" xml:"error-info"`
	ErrorCode	int		`json:"error-code" xml:"error-code"`
	ErrorModule	string		`json:"error-module" xml:"error-module"`
}
```

## type [Errors](<errors.go#L52>)

Errors holds a slice []Error
```go
type Errors struct {
	XMLName	xml.Name	`json:"-" xml:"errors"`
	Errors	[]Error		`json:"error" xml:"error"`
}
```

## func (*Errors) [AddError(Error)](<errors.go#L72>)

AddError appends a new Error


```go
func (err *Errors) AddError(newErr Error)
```
## func (*Errors) [Error() string](<errors.go#L59>)

Error returns a string, constructed from the contained
errors


```go
func (e *Errors) Error() string
```
## func (*Errors) [GetErrors() Errors](<errors.go#L67>)

GetErrors returns Errors


```go
func (e *Errors) GetErrors() *Errors
```
## func (*Errors) [String() string](<errors.go#L84>)

```go
func (e *Errors) String() string
```

