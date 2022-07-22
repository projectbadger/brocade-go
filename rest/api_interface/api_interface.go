package api_interface

import (
	"net/http"

	"github.com/projectbadger/brocade-go/session"
	"github.com/projectbadger/brocade-go/utils"
)

type RESTConfig struct {
	host        string
	baseURI     string
	session     session.Session
	contentType utils.ContentType
	client      utils.RequestClient
}

func NewRESTConfig(host, baseURI string, sess session.Session, contentType utils.ContentType, client utils.RequestClient) *RESTConfig {
	return &RESTConfig{
		host:        host,
		baseURI:     baseURI,
		session:     sess,
		contentType: contentType,
		client:      client,
	}
}

// Returns the specified host
func (c *RESTConfig) Host() string {
	return c.host
}

// Returns the specified base URI.
// Default: /rest
func (c *RESTConfig) BaseURI() string {
	return c.baseURI
}

// Returns the session interface
func (c *RESTConfig) Session() session.Session {
	return c.session
}

// Returns Content type
func (c *RESTConfig) ContentType() utils.ContentType {
	return c.contentType
}

// Sets the appropriate content headers on a request
func (c *RESTConfig) SetContentHeaders(req *http.Request) {
	c.contentType.SetRequestHeaders(req)
}

// Returns the provided HTTP client
func (c *RESTConfig) Client() utils.RequestClient {
	return c.client
}

// Sets all the appropriate authentication and content type
// headers.
func (c *RESTConfig) HandleRequest(req *http.Request) error {
	// contentType := c.ContentType()
	c.contentType.SetRequestHeaders(req)
	return c.session.HandleRequest(req)
}

// Marshal the data into the appropriate content type
func (c *RESTConfig) Marshal(v interface{}) ([]byte, error) {
	return c.contentType.Marshal(v)
}

// Unmarshal the data in the appropriate content type onto
// the interface
func (c *RESTConfig) Unmarshal(data []byte, v interface{}) error {
	return c.contentType.Unmarshal(data, v)
}
