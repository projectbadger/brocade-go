package fdmi

import (
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
	"github.com/projectbadger/brocade-go/rest/errors"
)

type RESTFDMI interface {
	Name() string
	URIPath() string
	GetHBA() ([]HBA, errors.BrocadeErr)
	GetHBAResponse() (*http.Response, error)
	GetPort() ([]Port, errors.BrocadeErr)
	GetPortResponse() (*http.Response, error)
}

type restFDMIImpl struct {
	config *api_interface.RESTConfig
}

func (r *restFDMIImpl) Name() string {
	return "brocade-fdmi"
}

func (r *restFDMIImpl) URIPath() string {
	return "/running/brocade-fdmi"
}

func NewRESTFDMI(config *api_interface.RESTConfig) RESTFDMI {
	return &restFDMIImpl{
		config: config,
	}
}

func (r *restFDMIImpl) GetHBA() ([]HBA, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.Name()+"/hba", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var hba []HBA
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &hba)

	return hba, errs
}

func (b *restFDMIImpl) GetHBAResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", b.config.Host()+b.config.BaseURI()+b.URIPath()+"/hba", nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, b.config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}

func (r *restFDMIImpl) GetPort() ([]Port, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.Name()+"/hba", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var port []Port
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &port)

	return port, errs
}

func (b *restFDMIImpl) GetPortResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", b.config.Host()+b.config.BaseURI()+b.URIPath()+"/port", nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, b.config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}
