package fibrechannel_logical_switch

import (
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
	"github.com/projectbadger/brocade-go/rest/errors"
)

// RESTFibrechannelLogicalSwitch describes an interface for
// interacting with the *fibrechannel-logical-switch* module.
// Fetch a new instance using the
// NewRESTFibrechannelLogicalSwitch function.
type RESTFibrechannelLogicalSwitch interface {
	Name() string
	URIPath() string
	GetFibrechannelLogicalSwitch() ([]FibrechannelLogicalSwitch, errors.BrocadeErr)
	GetFibrechannelLogicalSwitchResponse() (*http.Response, error)
}

// RESTFibrechannelLogicalSwitch implementation
type restFibrechannelLogicalSwitchImpl struct {
	config *api_interface.RESTConfig
}

// Returns a new RESTFibrechannelLogicalSwitch implementation
// for interacting with the *fibrechannel-logical-switch*
// module.
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

func (fls *restFibrechannelLogicalSwitchImpl) GetFibrechannelLogicalSwitch() ([]FibrechannelLogicalSwitch, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", fls.config.Host()+fls.config.BaseURI()+"/"+fls.URIPath()+"/fibrechannel-logical-switch", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, fls.config)
	if errs != nil {
		return nil, errs
	}
	var fibrechannelLogicalSwitch []FibrechannelLogicalSwitch
	ct := fls.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &fibrechannelLogicalSwitch)

	return fibrechannelLogicalSwitch, errors.NewFromErr(err)
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
