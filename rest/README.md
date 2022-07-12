
# rest

```go
import github.com/projectbadger/brocade-go/rest
```

Brocade REST interface

To use the REST interface, instantiate a new one and
interact with its components.

```go
rest := NewDefaultREST("10.0.0.10", "username", "password")
str := rest.Running().FRU().Name()
fmt.Println("FRU module name:", str)

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


## type [REST](<rest.go#L27>)

Brocade REST interface
```go
type REST interface {
	// Returns Running interface
	Running() running.RESTRunning
}
```

## func [NewREST() REST](<rest.go#L45>)

NewREST returns a new REST interface from config


```go
func NewREST(config *api_interface.RESTConfig) REST
```
## func [NewRESTDefault(string, string, string) REST](<rest.go#L88>)

NewRESTDefault returns a new REST interface
from config with predefined ContentTypeJSON,
sessionless session and http.DefaultClient as client


```go
func NewRESTDefault(host, username, password string) REST
```
## func [NewRESTJSON(string, string, Session, RequestClient) REST](<rest.go#L53>)

NewRESTJSON returns a new REST interface from config
with predefined ContentTypeJSON


```go
func NewRESTJSON(host, baseURI string, sess session.Session, client utils.RequestClient) REST
```
## func [NewRESTJSONSessionless(string, string, string, string, RequestClient) REST](<rest.go#L70>)

NewRESTJSONSessionless returns a new REST interface
from config with predefined ContentTypeJSON and
sessionless session


```go
func NewRESTJSONSessionless(host, baseURI, username, password string, client utils.RequestClient) REST
```
## func [NewRESTXML(string, string, Session, RequestClient) REST](<rest.go#L61>)

NewRESTXML returns a new REST interface from config
with predefined ContentTypeXML


```go
func NewRESTXML(host, baseURI string, sess session.Session, client utils.RequestClient) REST
```
## func [NewRESTXMLSessionless(string, string, string, string, RequestClient) REST](<rest.go#L79>)

NewRESTXMLSessionless returns a new REST interface
from config with predefined ContentTypeXML and
sessionless session


```go
func NewRESTXMLSessionless(host, baseURI, username, password string, client utils.RequestClient) REST
```

