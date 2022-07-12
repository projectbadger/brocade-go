package session

import (
	"errors"
	"net/http"

	"github.com/projectbadger/brocade-go/utils"
)

type sessionType int

const (
	sessionTypeSessionless sessionType = iota
	sessionTypeSession
	sessionStateLoggedOut = iota << 1
)

type Session interface {
	HandleRequest(*http.Request) error
	HandleLoginResponse(*http.Response) error
	Login(host, baseURI string, client utils.RequestClient) error
	Logout(host, baseURI string, client utils.RequestClient) error
	IsSessionless() bool
}

type sessionImpl struct {
	sessionType sessionType
	username    string
	password    string
	authHeader  string
}

func NewSessionlessSession(username, password string) Session {
	return &sessionImpl{
		username:    username,
		password:    password,
		sessionType: sessionTypeSessionless,
	}
}

func NewSession(username, password string) Session {
	return &sessionImpl{
		username:    username,
		password:    password,
		sessionType: sessionTypeSession,
	}
}

func (s *sessionImpl) IsSessionless() bool {
	return s.sessionType == sessionTypeSessionless
}

func (s *sessionImpl) HandleRequest(r *http.Request) error {
	return s.setAuthHeader(r)
}

func (s *sessionImpl) IsLoggedIn() bool {
	if s.sessionType == sessionTypeSession {
		if s.authHeader != "" {
			return true
		}
	}
	return false
}

func (s *sessionImpl) Login(host, baseURI string, client utils.RequestClient) error {
	req, err := http.NewRequest("POST", host+baseURI, nil)
	if err != nil {
		return err
	}
	err = s.HandleRequest(req)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = s.HandleLoginResponse(resp)
	return err
}

func (s *sessionImpl) Logout(host, baseURI string, client utils.RequestClient) error {
	r, err := http.NewRequest("POST", host+baseURI, nil)
	if err != nil {
		return err
	}
	// setJSONContentHeaders(r)
	err = s.HandleRequest(r)
	if err != nil {
		return err
	}
	resp, err := client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return s.HandleLogoutResponse(resp)
}

func (s *sessionImpl) HandleLoginResponse(r *http.Response) error {
	if r.StatusCode != 200 {
		if r.StatusCode%400 < 100 {
			s.authHeader = ""
		}
		return errors.New("login error " + r.Status)
	}
	s.authHeader = r.Header.Get("Authorization")
	if s.authHeader == "" {
		return errors.New("authorization header is empty")
	}
	return nil
}

func (s *sessionImpl) HandleLogoutResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 299 {
		s.authHeader = ""
		return nil
	} else if resp.StatusCode >= 400 && resp.StatusCode < 499 {
		s.authHeader = ""
	}
	return errors.New("logout response " + resp.Status)
}

func (s *sessionImpl) setAuthHeader(r *http.Request) (err error) {
	// fmt.Println("Setting auth header")
	authHeader := "Basic "
	usernameB := append([]byte(s.username), ':')
	if s.sessionType == sessionTypeSession {
		if s.authHeader != "" {
			r.Header.Set("Authorization", s.authHeader)
			return
			// authHeader = a.authHeader
		}
		authHeader = "Custom_Basic " + string(utils.Base64URLEncode(append(usernameB, []byte(s.password)...)))
		r.Header.Set("Authorization", authHeader)
		return
		// Prevent session login for now
	}
	// Is sessionless
	headerB := append(usernameB, []byte(s.password)...)

	r.Header.Set("Authorization", "Basic "+string(utils.Base64URLEncode(headerB)))
	return
}
