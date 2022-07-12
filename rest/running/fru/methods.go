package fru

import (
	"brocade/rest/api_interface"
	"net/http"
)

type RESTFRU interface {
	Name() string
	// GetBlade() ([]Port, error)
	GetBlade() ([]Blade, error)
	GetFan() ([]Fan, error)
	GetPowerSupply() ([]PowerSupply, error)
}

type restFRUImpl struct {
	config *api_interface.RESTConfig
}

func (r *restFRUImpl) Name() string {
	return "brocade-fru"
}

func (r *restFRUImpl) URIPath() string {
	return "/running/brocade-fru"
}

func NewRESTFRU(config *api_interface.RESTConfig) RESTFRU {
	return &restFRUImpl{
		config: config,
	}
}

func (r *restFRUImpl) GetBlade() ([]Blade, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/blade", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var blade []Blade
	ct := r.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &blade)
	return blade, err
}

func (r *restFRUImpl) GetFan() ([]Fan, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/fan", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var fan []Fan
	ct := r.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &fan)
	return fan, err
}

func (r *restFRUImpl) GetPowerSupply() ([]PowerSupply, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/power-supply", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var ps []PowerSupply
	ct := r.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &ps)
	return ps, err
}
