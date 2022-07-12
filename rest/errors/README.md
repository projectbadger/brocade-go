
# errors

```go
import brocade/rest/errors
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


## type [BrocadeErr](<errors.go#L16>)
```go
type BrocadeErr interface {
	Error() string
	GetErrors() *Errors
	AddError(Error)
}
```

## func [New(string) BrocadeErr](<errors.go#L142>)

```go
func New(message string, args ...interface{}) BrocadeErr
```
## func [NewFromErr(error) BrocadeErr](<errors.go#L22>)

```go
func NewFromErr(err error) BrocadeErr
```
## func [NewFromErrors() BrocadeErr](<errors.go#L39>)

```go
func NewFromErrors(err ...Error) BrocadeErr
```

## type [BrocadeErrorImpl](<errors.go#L48>)
```go
type BrocadeErrorImpl struct {
	XMLName	xml.Name	`json:"-" xml:"errors"`
	Errors	[]Error		`json:"error" xml:"error"`
}
```

## func (*BrocadeErrorImpl) [AddError(Error)](<errors.go#L73>)

```go
func (err *BrocadeErrorImpl) AddError(newErr Error)
```
## func (*BrocadeErrorImpl) [Error() string](<errors.go#L53>)

```go
func (err *BrocadeErrorImpl) Error() string
```
## func (*BrocadeErrorImpl) [GetErrors() Errors](<errors.go#L64>)

```go
func (err *BrocadeErrorImpl) GetErrors() *Errors
```

## type [Error](<errors.go#L108>)
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

## func (*Error) [Error() string](<errors.go#L135>)

```go
func (e *Error) Error() string
```
## func (*Error) [String() string](<errors.go#L124>)

```go
func (e *Error) String() string
```

## type [ErrorInfo](<errors.go#L118>)
```go
type ErrorInfo struct {
	XMLName		xml.Name	`json:"-" xml:"error-info"`
	ErrorCode	int		`json:"error-code" xml:"error-code"`
	ErrorModule	string		`json:"error-module" xml:"error-module"`
}
```

## type [Errors](<errors.go#L85>)
```go
type Errors struct {
	XMLName	xml.Name	`json:"-" xml:"errors"`
	Errors	[]Error		`json:"error" xml:"error"`
}
```

## func (*Errors) [Error() string](<errors.go#L90>)

```go
func (e *Errors) Error() string
```
## func (*Errors) [String() string](<errors.go#L97>)

```go
func (e *Errors) String() string
```

