package fabric

import (
	"fmt"
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
)

type RESTFabric interface {
	Name() string
	URIPath() string
	GetFabricSwitch() (*FabricSwitch, error)
	GetFabricSwitchResponse() (*http.Response, error)
}

type restFabricImpl struct {
	// *api_interface.RESTImpl
	// host        string
	// baseURI     string
	// session     session.Session
	// contentType utils.ContentType
	// client      utils.RequestClient
	config *api_interface.RESTConfig
}

func (r *restFabricImpl) Name() string {
	return "brocade-fabric"
}

func (r *restFabricImpl) URIPath() string {
	return "/running/brocade-fabric"
}

func NewRESTFabric(config *api_interface.RESTConfig) RESTFabric {
	return &restFabricImpl{
		config: config,
	}
	// return &restFabricImpl{
	// 	session:     sess,
	// 	host:        host,
	// 	baseURI:     baseURI,
	// 	client:      client,
	// 	contentType: contentType,
	// }
}

func (r *restFabricImpl) GetFabricSwitch() (*FabricSwitch, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.Name()+"/ha-status", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var fs FabricSwitch
	ct := r.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &fs)

	fmt.Println("Response struct", fs)
	fmt.Println("Response body", string(responseBytes))
	return &fs, err
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
