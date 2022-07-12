
# api_interface

```go
import brocade/rest/api_interface
```

## Index

- [CheckBodyForErrors(ContentType, ContentType) brocade_errors](#func-checkbodyforerrorscontenttype-contenttype-brocade-errors)
- [GetHTTPResponse() brocade_errors](#func-gethttpresponse-brocade-errors)
- [GetResponse() brocade_errors](#func-getresponse-brocade-errors)
- [GetResponseBody() error](#func-getresponsebody-error)

- [type RESTConfig](#type-restconfig)
  - [NewRESTConfig(string, string, Session, ContentType, RequestClient) RESTConfig](#func-newrestconfigstring-string-session-contenttype-requestclient-restconfig)
  - [BaseURI() string](#func-restconfig-baseuri-string)
  - [Client() utils](#func-restconfig-client-utils)
  - [ContentType() utils](#func-restconfig-contenttype-utils)
  - [HandleRequest() error](#func-restconfig-handlerequest-error)
  - [Host() string](#func-restconfig-host-string)
  - [Marshal() error](#func-restconfig-marshal-error)
  - [Session() session](#func-restconfig-session-session)
  - [SetContentHeaders()](#func-restconfig-setcontentheaders)
  - [Unmarshal() error](#func-restconfig-unmarshal-error)

## func [CheckBodyForErrors(ContentType, ContentType) brocade_errors](<handling.go#L48>)

```go
func CheckBodyForErrors(body []byte, contentType utils.ContentType) brocade_errors.BrocadeErr
```
## func [GetHTTPResponse() brocade_errors](<handling.go#L12>)

```go
func GetHTTPResponse(req *http.Request, config *RESTConfig) (*http.Response, brocade_errors.BrocadeErr)
```
## func [GetResponse() brocade_errors](<handling.go#L26>)

```go
func GetResponse(req *http.Request, config *RESTConfig) (*http.Response, []byte, brocade_errors.BrocadeErr)
```
## func [GetResponseBody() error](<handling.go#L76>)

```go
func GetResponseBody(resp *http.Response) ([]byte, error)
```


## type [RESTConfig](<api_interface.go#L19>)
```go
type RESTConfig struct {
	// contains filtered or unexported fields
}
```

## func [NewRESTConfig(string, string, Session, ContentType, RequestClient) RESTConfig](<api_interface.go#L27>)

```go
func NewRESTConfig(host, baseURI string, sess session.Session, contentType utils.ContentType, client utils.RequestClient) *RESTConfig
```

## func (*RESTConfig) [BaseURI() string](<api_interface.go#L44>)

Returns the specified base URI.
Default: /rest


```go
func (c *RESTConfig) BaseURI() string
```
## func (*RESTConfig) [Client() utils](<api_interface.go#L64>)

Returns the provided HTTP client


```go
func (c *RESTConfig) Client() utils.RequestClient
```
## func (*RESTConfig) [ContentType() utils](<api_interface.go#L54>)

Returns Content type


```go
func (c *RESTConfig) ContentType() utils.ContentType
```
## func (*RESTConfig) [HandleRequest() error](<api_interface.go#L68>)

```go
func (c *RESTConfig) HandleRequest(req *http.Request) error
```
## func (*RESTConfig) [Host() string](<api_interface.go#L38>)

Returns the specified host


```go
func (c *RESTConfig) Host() string
```
## func (*RESTConfig) [Marshal() error](<api_interface.go#L75>)

```go
func (c *RESTConfig) Marshal(v interface{}) ([]byte, error)
```
## func (*RESTConfig) [Session() session](<api_interface.go#L49>)

Returns the session interface


```go
func (c *RESTConfig) Session() session.Session
```
## func (*RESTConfig) [SetContentHeaders()](<api_interface.go#L59>)

Returns Content type


```go
func (c *RESTConfig) SetContentHeaders(req *http.Request)
```
## func (*RESTConfig) [Unmarshal() error](<api_interface.go#L79>)

```go
func (c *RESTConfig) Unmarshal(data []byte, v interface{}) error
```

