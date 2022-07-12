package rest

import (
	"brocade/rest/api_interface"
	"brocade/rest/running"
	"brocade/session"
	"brocade/utils"
	"net/http"
)

const (
	EndpointRESTBaseURI = "/rest"
)

type REST interface {
	// SetSession(sess session.Session)
	Running() running.RESTRunning
}

type restImpl struct {
	config *api_interface.RESTConfig
	// host        string
	// baseURI     string
	// session     session.Session
	// contentType utils.ContentType
	// client      utils.RequestClient
}

func (i *restImpl) Name() string {
	return "rest"
}

// func (i *restImpl) SetSession(sess session.Session) {
// 	if sess != nil {
// 		i.session = sess
// 	}
// }

// func (i *restImpl) SetClient(client utils.RequestClient) {
// 	if i != nil {
// 		i.client = client
// 	}
// }

func (i *restImpl) Running() running.RESTRunning {
	return running.NewRunning(i.config)
}

func NewREST(config *api_interface.RESTConfig) REST {
	return &restImpl{
		config: config,
	}
}

func NewRESTJSON(host, baseURI string, sess session.Session, client utils.RequestClient) REST {
	return &restImpl{
		config: api_interface.NewRESTConfig(host, baseURI, sess, utils.ContentTypeJSON, client),
	}
}

func NewRESTXML(host, baseURI string, sess session.Session, client utils.RequestClient) REST {
	return &restImpl{
		config: api_interface.NewRESTConfig(host, baseURI, sess, utils.ContentTypeXML, client),
	}
}

func NewRESTJSONSessionless(host, baseURI, username, password string, client utils.RequestClient) REST {
	return &restImpl{
		config: api_interface.NewRESTConfig(host, baseURI, session.NewSessionlessSession(username, password), utils.ContentTypeJSON, client),
	}
}

func NewRESTXMLSessionless(host, baseURI, username, password string, client utils.RequestClient) REST {
	return &restImpl{
		config: api_interface.NewRESTConfig(host, baseURI, session.NewSessionlessSession(username, password), utils.ContentTypeXML, client),
	}
}

func NewRESTDefault(host, username, password string) REST {
	return &restImpl{
		config: api_interface.NewRESTConfig(host, "/rest", session.NewSessionlessSession(username, password), utils.ContentTypeJSON, http.DefaultClient),
	}
}
