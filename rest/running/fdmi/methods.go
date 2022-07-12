package fdmi

import (
	"fmt"
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
)

type RESTFDMI interface {
	Name() string
	URIPath() string
	GetHBA() ([]HBA, error)
	GetHBAResponse() (*http.Response, error)
	GetPort() ([]Port, error)
	GetPortResponse() (*http.Response, error)
}

type restFDMIImpl struct {
	config *api_interface.RESTConfig
	// host        string
	// baseURI     string
	// session     session.Session
	// contentType utils.ContentType
	// client      utils.RequestClient
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

func (r *restFDMIImpl) GetHBA() ([]HBA, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.Name()+"/hba", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var hba []HBA
	ct := r.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &hba)

	fmt.Println("Response struct", hba)
	fmt.Println("Response body", string(responseBytes))
	return hba, err
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

func (r *restFDMIImpl) GetPort() ([]Port, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.Name()+"/hba", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var port []Port
	ct := r.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &port)

	fmt.Println("Response struct", port)
	fmt.Println("Response body", string(responseBytes))
	return port, err
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
