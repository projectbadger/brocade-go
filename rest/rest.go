// Brocade REST interface
/*
To use the REST interface, instantiate a new one and
interact with its components.

 rest := NewDefaultREST("10.0.0.10", "username", "password")
 str := rest.Running().FRU().Name()
 fmt.Println("FRU module name:", str)

*/
package rest

import (
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
	"github.com/projectbadger/brocade-go/rest/running"
	"github.com/projectbadger/brocade-go/session"
	"github.com/projectbadger/brocade-go/utils"
)

const (
	EndpointRESTBaseURI = "/rest"
)

// Brocade REST interface
type REST interface {
	// Returns Running interface
	Running() running.RESTRunning
}

type restImpl struct {
	config *api_interface.RESTConfig
}

func (i *restImpl) Name() string {
	return "rest"
}

func (i *restImpl) Running() running.RESTRunning {
	return running.NewRunning(i.config)
}

// NewREST returns a new REST interface from config
func NewREST(config *api_interface.RESTConfig) REST {
	return &restImpl{
		config: config,
	}
}

// NewRESTJSON returns a new REST interface from config
// with predefined ContentTypeJSON
func NewRESTJSON(host, baseURI string, sess session.Session, client utils.RequestClient) REST {
	return &restImpl{
		config: api_interface.NewRESTConfig(host, baseURI, sess, utils.ContentTypeJSON, client),
	}
}

// NewRESTXML returns a new REST interface from config
// with predefined ContentTypeXML
func NewRESTXML(host, baseURI string, sess session.Session, client utils.RequestClient) REST {
	return &restImpl{
		config: api_interface.NewRESTConfig(host, baseURI, sess, utils.ContentTypeXML, client),
	}
}

// NewRESTJSONSessionless returns a new REST interface
// from config with predefined ContentTypeJSON and
// sessionless session
func NewRESTJSONSessionless(host, baseURI, username, password string, client utils.RequestClient) REST {
	return &restImpl{
		config: api_interface.NewRESTConfig(host, baseURI, session.NewSessionlessSession(username, password), utils.ContentTypeJSON, client),
	}
}

// NewRESTXMLSessionless returns a new REST interface
// from config with predefined ContentTypeXML and
// sessionless session
func NewRESTXMLSessionless(host, baseURI, username, password string, client utils.RequestClient) REST {
	return &restImpl{
		config: api_interface.NewRESTConfig(host, baseURI, session.NewSessionlessSession(username, password), utils.ContentTypeXML, client),
	}
}

// NewRESTDefault returns a new REST interface
// from config with predefined ContentTypeJSON,
// sessionless session and http.DefaultClient as client
func NewRESTDefault(host, username, password string) REST {
	return &restImpl{
		config: api_interface.NewRESTConfig(host, "/rest", session.NewSessionlessSession(username, password), utils.ContentTypeJSON, http.DefaultClient),
	}
}
