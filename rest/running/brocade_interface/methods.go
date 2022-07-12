package brocade_interface

import (
	"fmt"
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
)

type RESTInterface interface {
	Name() string
	// GetFibrechannel() ([]Port, error)
	GetFibrechannelResponse() (*http.Response, error)
	GetFibrechannel() ([]Fibrechannel, error)
}

type restInterfaceImpl struct {
	config *api_interface.RESTConfig
	// host        string
	// baseURI     string
	// session     session.Session
	// contentType utils.ContentType
	// client      utils.RequestClient
}

func (r *restInterfaceImpl) Name() string {
	return "brocade-interface"
}

func (r *restInterfaceImpl) URIPath() string {
	return "/running/brocade-interface"
}

func NewRESTInterface(config *api_interface.RESTConfig) RESTInterface {
	return &restInterfaceImpl{
		config: config,
	}
}

func (r *restInterfaceImpl) GetFibrechannelResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+r.URIPath()+"/fibrechannel", nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, r.config)
	fmt.Println("Response: ", resp, "\nerror:", err)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}

func (r *restInterfaceImpl) GetFibrechannel() ([]Fibrechannel, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/fibrechannel", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var fc []Fibrechannel
	ct := r.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &fc)
	return fc, err
}
