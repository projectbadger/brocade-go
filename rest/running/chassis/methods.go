package chassis

import (
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
	"github.com/projectbadger/brocade-go/rest/errors"
)

type RESTChassis interface {
	Name() string
	URIPath() string
	GetChassis() (*Chassis, errors.BrocadeErr)
	GetChassisResponse() (*http.Response, error)
	GetHAStatus() (*HAStatus, errors.BrocadeErr)
	GetHAStatusResponse() (*http.Response, error)
}

type restChassisImpl struct {
	config *api_interface.RESTConfig
}

func NewChassis(cfg *api_interface.RESTConfig) RESTChassis {
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

func (c *restChassisImpl) GetChassis() (*Chassis, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+"/"+c.URIPath()+"/chassis", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, c.config)
	if errs != nil {
		return nil, errs
	}
	var chassis Chassis
	ct := c.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &chassis)

	return &chassis, errs
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

func (c *restChassisImpl) GetHAStatus() (*HAStatus, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+"/"+c.URIPath()+"/ha-status", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, c.config)
	if errs != nil {
		return nil, errs
	}
	var hastatus HAStatus
	ct := c.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &hastatus)

	return &hastatus, errs
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
