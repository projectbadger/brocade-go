package fibrechannel_trunk

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/projectbadger/brocade-go/rest/api_interface"
)

type RESTFibrechannelTrunk interface {
	Name() string
	URIPath() string
	GetTrunk(groupSourcePort string) (*BrocadeFibrechannelTrunk, error)
	GetGetTrunkResponse(groupSourcePort string) (*http.Response, error)
	GetPerformance() (*Performance, error)
	GetPerformanceResponse() (*http.Response, error)
}

type restFibrechannelTrunkImpl struct {
	config *api_interface.RESTConfig
}

func NewRESTFibrechannelTrunk(cfg *api_interface.RESTConfig) *restFibrechannelTrunkImpl {
	return &restFibrechannelTrunkImpl{
		config: cfg,
	}
}

func (c restFibrechannelTrunkImpl) Name() string {
	return "brocade-fibrechannel-trunk"
}

func (c restFibrechannelTrunkImpl) URIPath() string {
	return "/running/brocade-fibrechannel-trunk"
}

func (c *restFibrechannelTrunkImpl) GetTrunk(groupSourcePort string) (*BrocadeFibrechannelTrunk, error) {
	if groupSourcePort != "" {
		groupSourcePort = "/" + groupSourcePort
	}
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+"/"+c.URIPath()+"/trunk"+groupSourcePort, nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, c.config)
	if errs != nil {
		return nil, errs
	}
	var trunk BrocadeFibrechannelTrunk
	ct := c.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &trunk)

	return &trunk, err
}

func (c *restFibrechannelTrunkImpl) GetTrunkResponse(groupSourcePort string) (*http.Response, error) {
	if groupSourcePort != "" {
		groupSourcePort = "/" + groupSourcePort
	}
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+c.URIPath()+"/trunk"+groupSourcePort, nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, c.config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}

func (c *restFibrechannelTrunkImpl) GetPerformance(group string) (*Performance, error) {
	if group != "" {
		group = "/" + group
	}
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+"/"+c.URIPath()+"/performance"+group, nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, c.config)
	if errs != nil {
		return nil, errs
	}
	var performance Performance
	ct := c.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &performance)

	fmt.Println("Response struct", performance)
	fmt.Println("Response body", string(responseBytes))
	return &performance, err
}

func (c *restFibrechannelTrunkImpl) GetPerformanceResponse(group string) (*http.Response, error) {
	if group != "" {
		group = "/" + group
	}
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+c.URIPath()+"/performance"+group, nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, c.config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}

func (c *restFibrechannelTrunkImpl) GetTrunkArea(trunkIndex int) (*TrunkArea, error) {
	trunkIndexStr := ""
	if trunkIndex > 0 {
		trunkIndexStr = "/" + strconv.Itoa(trunkIndex)
	}
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+"/"+c.URIPath()+"/trunk-area"+trunkIndexStr, nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, c.config)
	if errs != nil {
		return nil, errs
	}
	var trunkArea TrunkArea
	ct := c.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &trunkArea)

	fmt.Println("Response struct", trunkArea)
	fmt.Println("Response body", string(responseBytes))
	return &trunkArea, err
}

func (c *restFibrechannelTrunkImpl) GetTrunkAreaResponse(trunkIndex int) (*http.Response, error) {
	trunkIndexStr := ""
	if trunkIndex > 0 {
		trunkIndexStr = "/" + strconv.Itoa(trunkIndex)
	}
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+c.URIPath()+"/trunk-area"+trunkIndexStr, nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, c.config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}
