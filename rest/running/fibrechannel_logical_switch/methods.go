package fibrechannel_logical_switch

import (
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
)

type RESTFibrechannelLogicalSwitch interface {
	Name() string
	URIPath() string
}

type restFibrechannelLogicalSwitchImpl struct {
	config *api_interface.RESTConfig
}

func NewFibrechannelLogicalSwitch(cfg *api_interface.RESTConfig) RESTFibrechannelLogicalSwitch {
	return &restFibrechannelLogicalSwitchImpl{
		config: cfg,
	}
}

func (fls restFibrechannelLogicalSwitchImpl) Name() string {
	return "brocade-fibrechannel-logical-switch"
}

func (fls restFibrechannelLogicalSwitchImpl) URIPath() string {
	return "/running/brocade-fibrechannel-logical-switch"
}

func (fls *restFibrechannelLogicalSwitchImpl) GetFibrechannelLogicalSwitch() ([]FibrechannelLogicalSwitch, error) {
	req, err := http.NewRequest("GET", fls.config.Host()+fls.config.BaseURI()+"/"+fls.URIPath()+"/fibrechannel-logical-switch", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, fls.config)
	if errs != nil {
		return nil, errs
	}
	var fibrechannelLogicalSwitch []FibrechannelLogicalSwitch
	ct := fls.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &fibrechannelLogicalSwitch)

	return fibrechannelLogicalSwitch, err
}

func (fls *restFibrechannelLogicalSwitchImpl) GetFibrechannelLogicalSwitchResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", fls.config.Host()+fls.config.BaseURI()+fls.URIPath()+"/fibrechannel-logical-switch", nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, fls.config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}
