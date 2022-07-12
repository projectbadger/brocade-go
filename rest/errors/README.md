
# errors

```go
import github.com/projectbadger/brocade-go/rest/errors
```

## Index

- [type BrocadeErr](#type-brocadeerr)
  - [New(string) BrocadeErr](#func-newstring-brocadeerr)
  - [NewFromErr(error) BrocadeErr](#func-newfromerrerror-brocadeerr)
  - [NewFromErrors() BrocadeErr](#func-newfromerrors-brocadeerr)
- [type BrocadeErrorImpl](#type-brocadeerrorimpl)
  - [AddError(Error)](#func-brocadeerrorimpl-adderrorerror)
  - [Error() string](#func-brocadeerrorimpl-error-string)
  - [GetErrors() Errors](#func-brocadeerrorimpl-geterrors-errors)
- [type Error](#type-error)
  - [Error() string](#func-error-error-string)
  - [String() string](#func-error-string-string)
- [type ErrorInfo](#type-errorinfo)
- [type Errors](#type-errors)
  - [Error() string](#func-errors-error-string)
  - [String() string](#func-errors-string-string)
- [Variables](#variables)

## Variables
```go
var (
	ErrUnmarshalFailed = errors.New("error unmarshal failed")
)

```


## type [BrocadeErr](<errors.go#L18>)

BrocadeErr represents an error that may hold brocade
error format
```go
type BrocadeErr interface {
	Error() string
	GetErrors() *Errors
	AddError(Error)
}
```

## func [New(string) BrocadeErr](<errors.go#L144>)

```go
func New(message string, args ...interface{}) BrocadeErr
```
## func [NewFromErr(error) BrocadeErr](<errors.go#L24>)

```go
func NewFromErr(err error) BrocadeErr
```
## func [NewFromErrors() BrocadeErr](<errors.go#L41>)

```go
func NewFromErrors(err ...Error) BrocadeErr
```

## type [BrocadeErrorImpl](<errors.go#L50>)
```go
type BrocadeErrorImpl struct {
	XMLName	xml.Name	`json:"-" xml:"errors"`
	Errors	[]Error		`json:"error" xml:"error"`
}
```

## func (*BrocadeErrorImpl) [AddError(Error)](<errors.go#L75>)

```go
func (err *BrocadeErrorImpl) AddError(newErr Error)
```
## func (*BrocadeErrorImpl) [Error() string](<errors.go#L55>)

```go
func (err *BrocadeErrorImpl) Error() string
```
## func (*BrocadeErrorImpl) [GetErrors() Errors](<errors.go#L66>)

```go
func (err *BrocadeErrorImpl) GetErrors() *Errors
```

## type [Error](<errors.go#L110>)
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

## func (*Error) [Error() string](<errors.go#L137>)

```go
func (e *Error) Error() string
```
## func (*Error) [String() string](<errors.go#L126>)

```go
func (e *Error) String() string
```

## type [ErrorInfo](<errors.go#L120>)
```go
type ErrorInfo struct {
	XMLName		xml.Name	`json:"-" xml:"error-info"`
	ErrorCode	int		`json:"error-code" xml:"error-code"`
	ErrorModule	string		`json:"error-module" xml:"error-module"`
}
```

## type [Errors](<errors.go#L87>)
```go
type Errors struct {
	XMLName	xml.Name	`json:"-" xml:"errors"`
	Errors	[]Error		`json:"error" xml:"error"`
}
```

## func (*Errors) [Error() string](<errors.go#L92>)

```go
func (e *Errors) Error() string
```
## func (*Errors) [String() string](<errors.go#L99>)

```go
func (e *Errors) String() string
```

