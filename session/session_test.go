package session

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"testing"

	"github.com/projectbadger/brocade-go/session/test_data"
)

type testRequestClient struct {
	headers map[string]string
	code    int
	body    []byte
}

func (c *testRequestClient) Do(*http.Request) (*http.Response, error) {
	resp := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     http.Header{},
	}
	for key, val := range c.headers {
		resp.Header.Set(key, val)
	}
	resp.Status = strconv.Itoa(c.code)
	resp.StatusCode = c.code
	resp.Body = io.NopCloser(bytes.NewReader(c.body))
	return resp, nil
}

func TestSession(t *testing.T) {
	t.Run("Authorization header handling", func(t *testing.T) {
		for _, test := range test_data.SessionImplAuthHeader {
			t.Run(test.Name, func(t *testing.T) {
				var sess Session
				switch sessionType(test.InputInt[0]) {
				case sessionTypeSession:
					sess = NewSession(test.InputStr[0], test.InputStr[1])
				case sessionTypeSessionless:
					sess = NewSessionlessSession(test.InputStr[0], test.InputStr[1])
				default:
					return
				}
				if test.ExpectedBool[0] != sess.IsSessionless() {
					t.Errorf("returned sessionless='%v' when it should have returned sessionless='%v'", sess.IsSessionless(), test.ExpectedBool[0])
				}
				req, _ := http.NewRequest("GET", "localhost", nil)
				if test.InputStr[2] == "" {
					sess.HandleRequest(req)
					authHeader := req.Header.Get("Authorization")
					if authHeader != test.ExpectedStr[0] {
						t.Errorf("before login returned header authorization='%s' when it should have returned header authorization='%s'", authHeader, test.ExpectedStr[0])
					}
					return
				}

				// set up a mock client to respond with a
				// provided header
				client := &testRequestClient{
					code: test.InputInt[1],
					headers: map[string]string{
						"Authorization": test.InputStr[2],
					},
				}
				err := sess.Login("localhost", "/rest/login", client)
				if err != nil {
					if test.ExpectedErr[0] != err.Error() {
						t.Errorf("returned error='%s' when expecting error='%s'", err.Error(), test.ExpectedErr[0])
					}
					// return
				} else if test.ExpectedErr[0] != "" {
					t.Errorf("should have thrown error='%s' at login", test.ExpectedErr[0])
				}

				req, _ = http.NewRequest("GET", "localhost", nil)
				sess.HandleRequest(req)
				authHeader := req.Header.Get("Authorization")
				if authHeader != test.ExpectedStr[1] {
					t.Errorf("after login returned header authorization='%s' when it should have returned header authorization='%s'", authHeader, test.ExpectedStr[1])
				}
				client = &testRequestClient{
					code: test.InputInt[2],
					headers: map[string]string{
						"Authorization": test.InputStr[3],
					},
				}
				err = sess.Logout("localhost", "/rest/login", client)
				if err != nil {
					if test.ExpectedErr[1] != err.Error() {
						t.Errorf("returned error='%s' when expecting error='%s'", err.Error(), test.ExpectedErr[1])
					}
					// return
				} else if test.ExpectedErr[1] != "" {
					t.Errorf("should have thrown error='%s' at login", test.ExpectedErr[1])
				}
				req, _ = http.NewRequest("GET", "localhost", nil)
				sess.HandleRequest(req)
				authHeader = req.Header.Get("Authorization")
				// Should be the same as before login
				if authHeader != test.ExpectedStr[2] {
					t.Errorf("after logout returned header authorization='%s' when it should have returned header authorization='%s'", authHeader, test.ExpectedStr[2])
				}
			})
		}
	})
}
