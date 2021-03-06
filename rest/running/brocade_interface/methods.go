package brocade_interface

import (
	"net/http"
	"strconv"

	"github.com/projectbadger/brocade-go/rest/api_interface"
	"github.com/projectbadger/brocade-go/rest/errors"
)

// RESTInterface describes an interface for interacting with
// the *brocade-interface* module.
// Fetch a new instance using the NewRESTInterface function.
type RESTInterface interface {
	Name() string
	// GetFibrechannel() ([]Port, error)
	GetFibrechannelResponse() (*http.Response, errors.BrocadeErr)
	GetFibrechannel() ([]Fibrechannel, errors.BrocadeErr)
}

// RESTInterface implementation
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

// Returns a new RESTInterface implementation for interacting
// with the *brocade-interface* module.
func NewRESTInterface(config *api_interface.RESTConfig) RESTInterface {
	return &restInterfaceImpl{
		config: config,
	}
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
	err = ct.UnmarshalResponse(responseBytes, &fc)
	return fc, errors.NewFromErr(err)
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

func (r *restInterfaceImpl) GetLogicalEPort(portIndex int) ([]LogicalEPort, errors.BrocadeErr) {
	portIndexStr := ""
	if portIndex > 0 {
		portIndexStr = "/" + strconv.Itoa(portIndex)
	}
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/logical-e-port"+portIndexStr, nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var lep []LogicalEPort
	ct := r.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &lep)
	return lep, errors.NewFromErr(err)
}

func (r *restInterfaceImpl) GetLogicalEPortResponse(portIndex int) (*http.Response, error) {
	portIndexStr := ""
	if portIndex > 0 {
		portIndexStr = "/" + strconv.Itoa(portIndex)
	}
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+r.URIPath()+"/logical-e-port"+portIndexStr, nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, r.config)
	// fmt.Println("Response: ", resp, "\nerror:", err)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}
