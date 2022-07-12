
# utils

```go
import brocade/utils
```

## Index

- [Base64URLDecode()](#func-base64urldecode)
- [Base64URLEncode()](#func-base64urlencode)

- [type ContentType](#type-contenttype)
  - [Marshal() error](#func-contenttype-marshal-error)
  - [SetRequestHeaders()](#func-contenttype-setrequestheaders)
  - [String() string](#func-contenttype-string-string)
  - [Unmarshal() brocade_errors](#func-contenttype-unmarshal-brocade-errors)
  - [UnmarshalResponse() brocade_errors](#func-contenttype-unmarshalresponse-brocade-errors)
  - [Valid() bool](#func-contenttype-valid-bool)
- [type RequestClient](#type-requestclient)
- [Variables](#variables)

## Variables
```go
var (
	ErrContentTypeInvalid		= brocade_errors.New("content type is invalid")
	ErrContentTypeNil		= brocade_errors.New("content type is nil")
	ErrNoParsableContentFound	= brocade_errors.New("no parsable content found")
)

```

## func [Base64URLDecode()](<utils.go#L15>)

base64URLDecode BASE64 URL decodes the content without
any padding.


```go
func Base64URLDecode(content []byte) []byte
```
## func [Base64URLEncode()](<utils.go#L7>)

base64URLEncode BASE64 URL encodes the content without
any padding.


```go
func Base64URLEncode(content []byte) []byte
```


## type [ContentType](<types.go#L10>)
```go
type ContentType string
```

## func (*ContentType) [Marshal() error](<types.go#L133>)

```go
func (ct *ContentType) Marshal(v interface{}) ([]byte, error)
```
## func (*ContentType) [SetRequestHeaders()](<types.go#L12>)

```go
func (ct *ContentType) SetRequestHeaders(req *http.Request)
```
## func (*ContentType) [String() string](<types.go#L29>)

```go
func (ct *ContentType) String() string
```
## func (*ContentType) [Unmarshal() brocade_errors](<types.go#L36>)

```go
func (ct *ContentType) Unmarshal(data []byte, v interface{}) brocade_errors.BrocadeErr
```
## func (*ContentType) [UnmarshalResponse() brocade_errors](<types.go#L49>)

```go
func (ct *ContentType) UnmarshalResponse(data []byte, v interface{}) brocade_errors.BrocadeErr
```
## func (*ContentType) [Valid() bool](<types.go#L17>)

```go
func (ct *ContentType) Valid() bool
```

## type [RequestClient](<types.go#L146>)
```go
type RequestClient interface {
	Do(*http.Request) (*http.Response, error)
}
```

