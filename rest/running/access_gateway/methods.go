package access_gateway

import (
	"fmt"
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
	"github.com/projectbadger/brocade-go/rest/errors"
)

// RESTAccessGateway describes an interface for interacting
// with the *brocade-access-gateway* module.
// Fetch a new instance using the NewRESTBrocadeAccessGateway
// function.
type RESTAccessGateway interface {
	Name() string
	GetPortGroupResponse() (*http.Response, errors.BrocadeErr)
	GetPortGroup() (*PortGroup, errors.BrocadeErr)
	GetNPortMapResponse() (*http.Response, errors.BrocadeErr)
	GetNPortMap() ([]NPortMap, errors.BrocadeErr)
	GetFPortListResponse() (*http.Response, errors.BrocadeErr)
	GetFPortList() ([]FPortList, errors.BrocadeErr)
}

// RESTAccesGateway implementation
type restAccessGatewayImpl struct {
	config *api_interface.RESTConfig
}

// Returns a new RESTAccessGateway implementation for
// interacting with the *brocade-access-gateway* module.
func NewRESTAccessGateway(cfg *api_interface.RESTConfig) RESTAccessGateway {
	return &restAccessGatewayImpl{
		config: cfg,
	}
}

func (ag restAccessGatewayImpl) Name() string {
	return "brocade-access-gateway"
}

func (ag restAccessGatewayImpl) URIPath() string {
	return "/running/brocade-access-gateway"
}

func (ag *restAccessGatewayImpl) GetPortGroupResponse() (*http.Response, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", ag.config.Host()+ag.config.BaseURI()+ag.URIPath()+"/port-group", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	resp, errs := api_interface.GetHTTPResponse(req, ag.config)
	fmt.Println("Response: ", resp, "\nerror:", err)
	if errs != nil {
		return nil, errs
	}
	defer resp.Body.Close()
	return resp, errs
}

func (ag *restAccessGatewayImpl) GetPortGroup() (*PortGroup, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", ag.config.Host()+ag.config.BaseURI()+ag.URIPath()+"/port-group", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, ag.config)
	if err == nil {
		// Errors returned
		return nil, errs
	}
	var pg PortGroup
	ct := ag.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &pg)

	// fmt.Println("Response struct", pg)
	// fmt.Println("Response body", string(responseBytes))
	return &pg, errs
}

func (ag *restAccessGatewayImpl) GetNPortMapResponse() (*http.Response, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", ag.config.Host()+ag.config.BaseURI()+ag.URIPath()+"/n-port-map", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	resp, errs := api_interface.GetHTTPResponse(req, ag.config)
	fmt.Println("Response: ", resp, "\nerror:", err)
	if err != nil {
		return nil, errs
	}
	defer resp.Body.Close()
	return resp, errs
}

func (ag *restAccessGatewayImpl) GetNPortMap() ([]NPortMap, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", ag.config.Host()+ag.config.BaseURI()+ag.URIPath()+"/n-port-map", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, ag.config)
	if errs != nil {
		return nil, errs
	}
	var m []NPortMap
	ct := ag.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &m)

	// fmt.Println("Response struct", m)
	// fmt.Println("Response body", string(responseBytes))
	return m, errs
}

func (ag *restAccessGatewayImpl) GetFPortListResponse() (*http.Response, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", ag.config.Host()+ag.config.BaseURI()+ag.URIPath()+"/f-port-list", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	resp, errs := api_interface.GetHTTPResponse(req, ag.config)
	fmt.Println("Response: ", resp, "\nerror:", err)
	if errs != nil {
		return nil, errs
	}
	defer resp.Body.Close()
	return resp, errs
}

func (ag *restAccessGatewayImpl) GetFPortList() ([]FPortList, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", ag.config.Host()+ag.config.BaseURI()+ag.URIPath()+"/f-port-list", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, ag.config)
	if errs != nil {
		return nil, errs
	}
	var fpl []FPortList
	ct := ag.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &fpl)

	fmt.Println("Response struct", fpl)
	fmt.Println("Response body", string(responseBytes))
	return fpl, errs
}
