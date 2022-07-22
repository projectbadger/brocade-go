package fabric

import (
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
	"github.com/projectbadger/brocade-go/rest/errors"
)

// RESTFabric describes an interface for interacting with the
// *fabric* module.
// Fetch a new instance using the NewRESTFabric function.
type RESTFabric interface {
	Name() string
	URIPath() string
	GetFabricSwitch() (*FabricSwitch, errors.BrocadeErr)
	GetFabricSwitchResponse() (*http.Response, error)
}

// RESTFabric implementation
type restFabricImpl struct {
	config *api_interface.RESTConfig
}

func (r *restFabricImpl) Name() string {
	return "brocade-fabric"
}

func (r *restFabricImpl) URIPath() string {
	return "/running/brocade-fabric"
}

// Returns a new RESTFabric implementation for interacting
// with the *fabric* module.
func NewRESTFabric(config *api_interface.RESTConfig) RESTFabric {
	return &restFabricImpl{
		config: config,
	}
}

func (r *restFabricImpl) GetFabricSwitch() (*FabricSwitch, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.Name()+"/ha-status", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var fs FabricSwitch
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &fs)

	// fmt.Println("Response struct", fs)
	// fmt.Println("Response body", string(responseBytes))
	return &fs, errs
}

func (b *restFabricImpl) GetFabricSwitchResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", b.config.Host()+b.config.BaseURI()+b.URIPath()+"/fabric-switch", nil)
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
