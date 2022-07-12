package fibrechannel_configuration

import (
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
	"github.com/projectbadger/brocade-go/rest/errors"
)

type RESTFibrechannelConfiguration interface {
	Name() string
	URIPath() string
	GetSwitchConfiguration() (*SwitchConfiguration, errors.BrocadeErr)
	GetPortConfiguration() (*PortConfiguration, errors.BrocadeErr)
	GetFPortLoginSettings() (*FPortLoginSettings, errors.BrocadeErr)
	GetZoneConfiguration() (*ZoneConfiguration, errors.BrocadeErr)
	GetFabric() (*Fabric, errors.BrocadeErr)
}

type restFibrechannelConfigurationImpl struct {
	config *api_interface.RESTConfig
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

func (r *restFibrechannelConfigurationImpl) GetSwitchConfiguration() (*SwitchConfiguration, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/switch-configuration", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var sc SwitchConfiguration
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &sc)

	return &sc, errs
}

func (r *restFibrechannelConfigurationImpl) GetPortConfiguration() (*PortConfiguration, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/port", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var pc PortConfiguration
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &pc)

	return &pc, errs
}

func (r *restFibrechannelConfigurationImpl) GetFPortLoginSettings() (*FPortLoginSettings, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/f-port-login-settings", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var fpls FPortLoginSettings
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &fpls)

	return &fpls, errs
}

func (r *restFibrechannelConfigurationImpl) GetZoneConfiguration() (*ZoneConfiguration, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/zone-configuration", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var zc ZoneConfiguration
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &zc)

	return &zc, errs
}

func (r *restFibrechannelConfigurationImpl) GetFabric() (*Fabric, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/fabric", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var f Fabric
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &f)

	return &f, errs
}
