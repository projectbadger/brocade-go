package fibrechannel_configuration

import (
	"brocade/rest/api_interface"
	"fmt"
	"net/http"
)

type RESTFibrechannelConfiguration interface {
	Name() string
	URIPath() string
	GetSwitchConfiguration() (*SwitchConfiguration, error)
	GetPortConfiguration() (*PortConfiguration, error)
	GetFPortLoginSettings() (*FPortLoginSettings, error)
	GetZoneConfiguration() (*ZoneConfiguration, error)
	GetFabric() (*Fabric, error)
}

type restFibrechannelConfigurationImpl struct {
	config *api_interface.RESTConfig
	// host        string
	// baseURI     string
	// session     session.Session
	// contentType utils.ContentType
	// client      utils.RequestClient
}

func (r *restFibrechannelConfigurationImpl) Name() string {
	return "brocade-fibrechannel-configuration"
}

func (r *restFibrechannelConfigurationImpl) URIPath() string {
	return "/running/brocade-fibrechannel-configuration"
}

func NewRESTFibrechannelConfiguration(config *api_interface.RESTConfig) RESTFibrechannelConfiguration {
	return &restFibrechannelConfigurationImpl{
		config: config,
	}
}

func (r *restFibrechannelConfigurationImpl) GetSwitchConfiguration() (*SwitchConfiguration, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/switch-configuration", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var sc SwitchConfiguration
	ct := r.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &sc)

	fmt.Println("Response struct", sc)
	fmt.Println("Response body", string(responseBytes))
	return &sc, err
}

func (r *restFibrechannelConfigurationImpl) GetPortConfiguration() (*PortConfiguration, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/port", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var pc PortConfiguration
	ct := r.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &pc)

	fmt.Println("Response struct", pc)
	fmt.Println("Response body", string(responseBytes))
	return &pc, err
}

func (r *restFibrechannelConfigurationImpl) GetFPortLoginSettings() (*FPortLoginSettings, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/f-port-login-settings", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var fpls FPortLoginSettings
	ct := r.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &fpls)

	fmt.Println("Response struct", fpls)
	fmt.Println("Response body", string(responseBytes))
	return &fpls, err
}

func (r *restFibrechannelConfigurationImpl) GetZoneConfiguration() (*ZoneConfiguration, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/zone-configuration", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var zc ZoneConfiguration
	ct := r.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &zc)

	fmt.Println("Response struct", zc)
	fmt.Println("Response body", string(responseBytes))
	return &zc, err
}

func (r *restFibrechannelConfigurationImpl) GetFabric() (*Fabric, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/fabric", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var f Fabric
	ct := r.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &f)

	fmt.Println("Response struct", f)
	fmt.Println("Response body", string(responseBytes))
	return &f, err
}
