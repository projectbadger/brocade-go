package license

import (
	"encoding/xml"
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
	"github.com/projectbadger/brocade-go/rest/errors"
)

type RESTLicense interface {
	Name() string
	GetLicense() (*License, errors.BrocadeErr)
}

type restLicenseImpl struct {
	config *api_interface.RESTConfig
}

func (r *restLicenseImpl) Name() string {
	return "brocade-fru"
}

func NewRESTLicense(config *api_interface.RESTConfig) RESTLicense {
	return &restLicenseImpl{
		config: config,
	}
}

func (r *restLicenseImpl) GetLicense() (*License, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.Name()+"/blade", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var response struct {
		XMLName  xml.Name `json:"-" xml:"Response"`
		Response struct {
			License License `json:"fibrechannel" xml:"fibrechannel"`
		}
	}
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &response)
	if err != nil {
		// Errors returned
		return nil, errs
	}

	return &response.Response.License, errs
}
