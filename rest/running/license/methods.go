package license

import (
	"brocade/rest/api_interface"
	"encoding/xml"
	"fmt"
	"net/http"
)

type RESTLicense interface {
	Name() string
	// GetBlade() ([]Port, error)
	GetLicense() (*License, error)
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

func (r *restLicenseImpl) GetLicense() (*License, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.Name()+"/blade", nil)
	if err != nil {
		return nil, err
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
	err = r.config.Unmarshal(responseBytes, &response)
	if err != nil {
		// Errors returned
		return nil, err
	}

	fmt.Println("Response struct", response)
	fmt.Println("Response body", string(responseBytes))
	return &response.Response.License, err
}
