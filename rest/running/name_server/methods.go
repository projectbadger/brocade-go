package name_server

import (
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
)

type RESTNameServer interface {
	Name() string
	URIPath() string
	GetNameServer() ([]FibrechannelNameServer, error)
	GetNameServerResponse() (*http.Response, error)
	GetNameServerForPort(portId string) (*FibrechannelNameServer, error)
	GetNameServerForPortResponse(portId string) (*http.Response, error)
}

type restNameServerImpl struct {
	config *api_interface.RESTConfig
}

func NewRESTNameServer(cfg *api_interface.RESTConfig) RESTNameServer {
	return &restNameServerImpl{
		config: cfg,
	}
}

func (ns restNameServerImpl) Name() string {
	return "brocade-name-server"
}

func (ns restNameServerImpl) URIPath() string {
	return "/running/brocade-name-server"
}

func (ns *restNameServerImpl) GetNameServer() ([]FibrechannelNameServer, error) {
	req, err := http.NewRequest("GET", ns.config.Host()+ns.config.BaseURI()+"/"+ns.URIPath()+"/fibrechannel-name-server", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, ns.config)
	if errs != nil {
		return nil, errs
	}
	var nameServers []FibrechannelNameServer
	ct := ns.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &nameServers)

	return nameServers, err
}

func (ns *restNameServerImpl) GetNameServerResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", ns.config.Host()+ns.config.BaseURI()+ns.URIPath()+"/fibrechannel-name-server", nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, ns.config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}

func (ns *restNameServerImpl) GetNameServerForPort(portId string) (*FibrechannelNameServer, error) {
	if portId != "" {
		portId = "/" + portId
	}
	req, err := http.NewRequest("GET", ns.config.Host()+ns.config.BaseURI()+"/"+ns.URIPath()+"/fibrechannel-name-server/"+portId, nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, ns.config)
	if errs != nil {
		return nil, errs
	}
	var nameServer FibrechannelNameServer
	ct := ns.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &nameServer)

	return &nameServer, err
}

func (ns *restNameServerImpl) GetNameServerForPortResponse(portId string) (*http.Response, error) {
	if portId != "" {
		portId = "/" + portId
	}
	req, err := http.NewRequest("GET", ns.config.Host()+ns.config.BaseURI()+ns.URIPath()+"/fibrechannel-name-server/"+portId, nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, ns.config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}
