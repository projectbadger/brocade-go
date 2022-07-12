
# session

```go
import github.com/projectbadger/brocade-go/session
```

## Index

- [type Session](#type-session)
  - [NewSession(string, string) Session](#func-newsessionstring-string-session)
  - [NewSessionlessSession(string, string) Session](#func-newsessionlesssessionstring-string-session)


## type [Session](<session.go#L19>)

Session contains methods for authentication and session management
```go
type Session interface {
	HandleRequest(*http.Request) error
	HandleLoginResponse(*http.Response) error
	Login(host, baseURI string, client utils.RequestClient) error
	Logout(host, baseURI string, client utils.RequestClient) error
	IsSessionless() bool
}
```

## func [NewSession(string, string) Session](<session.go#L42>)

```go
func NewSession(username, password string) Session
```
## func [NewSessionlessSession(string, string) Session](<session.go#L34>)

```go
func NewSessionlessSession(username, password string) Session
```

