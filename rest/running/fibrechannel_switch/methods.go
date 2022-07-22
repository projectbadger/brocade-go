package fibrechannel_switch

import (
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
)

type RESTChassis interface {
	Name() string
	URIPath() string
	GetFibrechannelSwitch() (*FibrechannelSwitch, error)
	GetFibrechannelSwitchResponse() (*http.Response, error)
	GetVirtualFabricFibrechannelSwitch() (*FibrechannelSwitch, error)
	GetVirtualFabricFibrechannelSwitchResponse() (*http.Response, error)
}

type restFibrechannelSwitchImpl struct {
	config *api_interface.RESTConfig
}

func NewChassis(cfg *api_interface.RESTConfig) *restFibrechannelSwitchImpl {
	return &restFibrechannelSwitchImpl{
		config: cfg,
	}
}

func (c restFibrechannelSwitchImpl) Name() string {
	return "brocade-fibrechannel-switch"
}

func (c restFibrechannelSwitchImpl) URIPath() string {
	return "/running/brocade-fibrechannel-switch"
}

func (c *restFibrechannelSwitchImpl) GetFibrechannelSwitch() (*FibrechannelSwitch, error) {
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+"/"+c.URIPath()+"/fibrechannel-switch", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, c.config)
	if errs != nil {
		return nil, errs
	}
	var fibrechannelSwitch FibrechannelSwitch
	ct := c.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &fibrechannelSwitch)

	return &fibrechannelSwitch, err
}

func (c *restFibrechannelSwitchImpl) GetFibrechannelSwitchResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+c.URIPath()+"/fibrechannel-switch", nil)
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

func (c *restFibrechannelSwitchImpl) GetVirtualFabricFibrechannelSwitch(vfId string) (*FibrechannelSwitch, error) {
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+"/running/fibrechannel-switch?vfid="+vfId, nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, c.config)
	if errs != nil {
		return nil, errs
	}
	var fibrechannelSwitch FibrechannelSwitch
	ct := c.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &fibrechannelSwitch)

	return &fibrechannelSwitch, err
}

func (c *restFibrechannelSwitchImpl) GetVirtualFabricFibrechannelSwitchResponse(vfId string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+"/running/fibrechannel-switch?vfid="+vfId, nil)
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