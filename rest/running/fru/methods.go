package fru

import (
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
	"github.com/projectbadger/brocade-go/rest/errors"
)

// RESTFRU describes an interface for interacting with the
// *fru* module.
// Fetch a new instance using the NewRESTFRU function.
type RESTFRU interface {
	Name() string
	GetBlade() ([]Blade, errors.BrocadeErr)
	GetFan() ([]Fan, errors.BrocadeErr)
	GetPowerSupply() ([]PowerSupply, errors.BrocadeErr)
}

// RESTFRU implementation
type restFRUImpl struct {
	config *api_interface.RESTConfig
}

func (r restFRUImpl) Name() string {
	return "brocade-fru"
}

func (r restFRUImpl) URIPath() string {
	return "/running/brocade-fru"
}

// Returns a new RESTFRU interface implementation for
// interactiong with the *fru* module
func NewRESTFRU(config *api_interface.RESTConfig) RESTFRU {
	return &restFRUImpl{
		config: config,
	}
}

func (r *restFRUImpl) GetBlade() ([]Blade, errors.BrocadeErr) {
	req, httpErr := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/blade", nil)
	if httpErr != nil {
		return nil, errors.NewFromErr(httpErr)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var blade []Blade
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &blade)
	return blade, errs
}

func (r *restFRUImpl) GetFan() ([]Fan, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/fan", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var fan []Fan
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &fan)
	return fan, errs
}

func (r *restFRUImpl) GetPowerSupply() ([]PowerSupply, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/power-supply", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var ps []PowerSupply
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &ps)
	return ps, errs
}
