
# session

```go
import github.com/projectbadger/brocade-go/session
```

## Index

- [type LoginSessionFunc](#type-loginsessionfunc)
- [type Session](#type-session)
  - [NewSession(string, string) Session](#func-newsessionstring-string-session)
  - [NewSessionlessSession(string, string) Session](#func-newsessionlesssessionstring-string-session)


## type [LoginSessionFunc](<session.go#L89>)

LoginSessionFunc represents a function to be executed
during a login session, triggered by Session.LoginSession.
```go
type LoginSessionFunc func(session Session) error
```

## type [Session](<session.go#L19>)

Session contains methods for authentication and session management
```go
type Session interface {
	HandleRequest(*http.Request) error
	HandleLoginResponse(*http.Response) error
	Login(host, baseURI string, client utils.RequestClient) error
	Logout(host, baseURI string, client utils.RequestClient) error
	LoginSession(host, baseURI string, client utils.RequestClient, sessionFunc LoginSessionFunc) error
	IsSessionless() bool
}
```

## func [NewSession(string, string) Session](<session.go#L43>)

```go
func NewSession(username, password string) Session
```
## func [NewSessionlessSession(string, string) Session](<session.go#L35>)

```go
func NewSessionlessSession(username, password string) Session
```

