package brocade_interface

import (
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
	"github.com/projectbadger/brocade-go/rest/errors"
)

type RESTInterface interface {
	Name() string
	// GetFibrechannel() ([]Port, error)
	GetFibrechannelResponse() (*http.Response, errors.BrocadeErr)
	GetFibrechannel() ([]Fibrechannel, errors.BrocadeErr)
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

func (r *restInterfaceImpl) GetFibrechannelResponse() (*http.Response, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+r.URIPath()+"/fibrechannel", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	resp, errs := api_interface.GetHTTPResponse(req, r.config)
	// fmt.Println("Response: ", resp, "\nerror:", err)
	if errs != nil {
		return nil, errs
	}
	defer resp.Body.Close()
	return resp, errs
}

func (r *restInterfaceImpl) GetFibrechannel() ([]Fibrechannel, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/fibrechannel", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var fc []Fibrechannel
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &fc)
	return fc, errs
}
