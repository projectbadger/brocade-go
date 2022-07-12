package chassis

import (
	"brocade/rest/api_interface"
	"fmt"
	"net/http"
)

type RESTChassis interface {
	Name() string
	URIPath() string
	GetChassis() (*Chassis, error)
	GetChassisResponse() (*http.Response, error)
	GetHAStatus() (*HAStatus, error)
	GetHAStatusResponse() (*http.Response, error)
}

type restChassisImpl struct {
	config *api_interface.RESTConfig
}

func NewChassis(cfg *api_interface.RESTConfig) *restChassisImpl {
	return &restChassisImpl{
		config: cfg,
	}
}

func (c restChassisImpl) Name() string {
	return "brocade-chassis"
}

func (c restChassisImpl) URIPath() string {
	return "/running/brocade-chassis"
}

func (c *restChassisImpl) GetChassis() (*Chassis, error) {
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+"/"+c.URIPath()+"/chassis", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, c.config)
	if errs != nil {
		return nil, errs
	}
	var chassis Chassis
	ct := c.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &chassis)

	return &chassis, err
}

func (c *restChassisImpl) GetChassisResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+c.URIPath()+"/chassis", nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, c.config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}

func (c *restChassisImpl) GetHAStatus() (*HAStatus, error) {
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+"/"+c.URIPath()+"/ha-status", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, c.config)
	if errs != nil {
		return nil, errs
	}
	var hastatus HAStatus
	ct := c.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &hastatus)

	fmt.Println("Response struct", hastatus)
	fmt.Println("Response body", string(responseBytes))
	return &hastatus, err
}

func (c *restChassisImpl) GetHAStatusResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+c.URIPath()+"/ha-status", nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, c.config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}
