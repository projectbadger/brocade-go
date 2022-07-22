package license

import (
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
	"github.com/projectbadger/brocade-go/rest/errors"
)

// RESTLicense describes an interface for interacting with the
// *license* module.
// Fetch a new instance using the NewRESTLicense function.
type RESTLicense interface {
	Name() string
	GetLicense() (*License, errors.BrocadeErr)
}

// RESTLicense implementation
type restLicenseImpl struct {
	config *api_interface.RESTConfig
}

func (r *restLicenseImpl) Name() string {
	return "brocade-fru"
}

// NewRESTLicense returns a new RESTLicense implementation
// for interacting with the *license* module
func NewRESTLicense(config *api_interface.RESTConfig) RESTLicense {
	return &restLicenseImpl{
		config: config,
	}
}

// GetLicense returns the license.
func (r *restLicenseImpl) GetLicense() (*License, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.Name()+"/blade", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var license License
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &license)
	if err != nil {
		// Errors returned
		return nil, errs
	}

	return &license, errs
}
