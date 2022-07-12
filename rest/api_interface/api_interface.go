package api_interface

import (
	"fmt"
	"net/http"

	"github.com/projectbadger/brocade-go/session"
	"github.com/projectbadger/brocade-go/utils"
)

// type RESTInterface interface {
// 	Name() string
// 	Host() string
// 	BaseURI() string
// 	ContentType() utils.ContentType
// 	Client() utils.RequestClient
// 	Setup(string, string, session.Session, utils.ContentType, utils.RequestClient)
// }

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

// Returns Content type
func (c *RESTConfig) SetContentHeaders(req *http.Request) {
	c.contentType.SetRequestHeaders(req)
}

// Returns the provided HTTP client
func (c *RESTConfig) Client() utils.RequestClient {
	return c.client
}

func (c *RESTConfig) HandleRequest(req *http.Request) error {
	// contentType := c.ContentType()
	c.contentType.SetRequestHeaders(req)
	fmt.Printf("Content Type: '%#v' \nSession: '%#v'\n", c.contentType, c.session)
	return c.session.HandleRequest(req)
}

func (c *RESTConfig) Marshal(v interface{}) ([]byte, error) {
	return c.contentType.Marshal(v)
}

func (c *RESTConfig) Unmarshal(data []byte, v interface{}) error {
	return c.contentType.Unmarshal(data, v)
}

// type RESTImpl struct {
// 	config *RESTConfig
// }

// func (r RESTImpl) WithConfig(config *RESTConfig) *RESTImpl {
// 	return &RESTImpl{
// 		config: config,
// 	}
// }

// func (r *RESTImpl) Setup(config *RESTConfig) {
// 	r.config = config
// }

// func (r *RESTImpl) Name() string {
// 	return "rest"
// }

// func (r *RESTImpl) Host() string {
// 	return r.config.host
// }

// func (r *RESTImpl) BaseURI() string {
// 	return r.config.baseURI
// }

// func (r *RESTImpl) Session() session.Session {
// 	return r.config.session
// }

// func (r *RESTImpl) ContentType() utils.ContentType {
// 	return r.config.contentType
// }

// func (r *RESTImpl) Client() utils.RequestClient {
// 	return r.config.client
// }
