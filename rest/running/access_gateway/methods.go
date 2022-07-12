package access_gateway

import (
	"fmt"
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
)

type RESTAccessGateway interface {
	Name() string
	GetPortGroupResponse() (*http.Response, error)
	GetPortGroup() (*PortGroup, error)
	GetNPortMapResponse() (*http.Response, error)
	GetNPortMap() ([]NPortMap, error)
	GetFPortListResponse() (*http.Response, error)
	GetFPortList() ([]FPortList, error)
}

type restAccessGatewayImpl struct {
	config *api_interface.RESTConfig
}

func NewAccessGateway(cfg *api_interface.RESTConfig) *restAccessGatewayImpl {
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

func (ag *restAccessGatewayImpl) GetPortGroupResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", ag.config.Host()+ag.config.BaseURI()+ag.URIPath()+"/port-group", nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, ag.config)
	fmt.Println("Response: ", resp, "\nerror:", err)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}

func (ag *restAccessGatewayImpl) GetPortGroup() (*PortGroup, error) {
	req, err := http.NewRequest("GET", ag.config.Host()+ag.config.BaseURI()+ag.URIPath()+"/port-group", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, ag.config)
	if err == nil {
		// Errors returned
		return nil, errs
	}
	var pg PortGroup
	ct := ag.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &pg)

	fmt.Println("Response struct", pg)
	fmt.Println("Response body", string(responseBytes))
	return &pg, err
}

func (ag *restAccessGatewayImpl) GetNPortMapResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", ag.config.Host()+ag.config.BaseURI()+ag.URIPath()+"/n-port-map", nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, ag.config)
	fmt.Println("Response: ", resp, "\nerror:", err)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}

func (ag *restAccessGatewayImpl) GetNPortMap() ([]NPortMap, error) {
	req, err := http.NewRequest("GET", ag.config.Host()+ag.config.BaseURI()+ag.URIPath()+"/n-port-map", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, ag.config)
	if errs != nil {
		return nil, errs
	}
	var m []NPortMap
	ct := ag.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &m)

	fmt.Println("Response struct", m)
	fmt.Println("Response body", string(responseBytes))
	return m, err
}

func (ag *restAccessGatewayImpl) GetFPortListResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", ag.config.Host()+ag.config.BaseURI()+ag.URIPath()+"/f-port-list", nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, ag.config)
	fmt.Println("Response: ", resp, "\nerror:", err)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}

func (ag *restAccessGatewayImpl) GetFPortList() ([]FPortList, error) {
	req, err := http.NewRequest("GET", ag.config.Host()+ag.config.BaseURI()+ag.URIPath()+"/f-port-list", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, ag.config)
	if errs != nil {
		return nil, errs
	}
	var fpl []FPortList
	ct := ag.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &fpl)

	fmt.Println("Response struct", fpl)
	fmt.Println("Response body", string(responseBytes))
	return fpl, err
}
