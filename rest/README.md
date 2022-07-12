
# rest

```go
import brocade/rest
```

## Index

- [type REST](#type-rest)
  - [NewREST() REST](#func-newrest-rest)
  - [NewRESTDefault(string, string, string) REST](#func-newrestdefaultstring-string-string-rest)
  - [NewRESTJSON(string, string, Session, RequestClient) REST](#func-newrestjsonstring-string-session-requestclient-rest)
  - [NewRESTJSONSessionless(string, string, string, string, RequestClient) REST](#func-newrestjsonsessionlessstring-string-string-string-requestclient-rest)
  - [NewRESTXML(string, string, Session, RequestClient) REST](#func-newrestxmlstring-string-session-requestclient-rest)
  - [NewRESTXMLSessionless(string, string, string, string, RequestClient) REST](#func-newrestxmlsessionlessstring-string-string-string-requestclient-rest)
- [Variables](#variables)

## Variables
```go
const (
	EndpointRESTBaseURI = "/rest"
)

```


## type [REST](<rest.go#L15>)
```go
type REST interface {
	// SetSession(sess session.Session)
	Running() running.RESTRunning
}
```

## func [NewREST() REST](<rest.go#L49>)

```go
func NewREST(config *api_interface.RESTConfig) REST
```
## func [NewRESTDefault(string, string, string) REST](<rest.go#L79>)

```go
func NewRESTDefault(host, username, password string) REST
```
## func [NewRESTJSON(string, string, Session, RequestClient) REST](<rest.go#L55>)

```go
func NewRESTJSON(host, baseURI string, sess session.Session, client utils.RequestClient) REST
```
## func [NewRESTJSONSessionless(string, string, string, string, RequestClient) REST](<rest.go#L67>)

```go
func NewRESTJSONSessionless(host, baseURI, username, password string, client utils.RequestClient) REST
```
## func [NewRESTXML(string, string, Session, RequestClient) REST](<rest.go#L61>)

```go
func NewRESTXML(host, baseURI string, sess session.Session, client utils.RequestClient) REST
```
## func [NewRESTXMLSessionless(string, string, string, string, RequestClient) REST](<rest.go#L73>)

```go
func NewRESTXMLSessionless(host, baseURI, username, password string, client utils.RequestClient) REST
```

