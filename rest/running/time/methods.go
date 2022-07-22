package time

import (
	"net/http"
	"strconv"

	"github.com/projectbadger/brocade-go/rest/api_interface"
)

// RESTTime describes an interface for interacting with the
// *time* module.
// Fetch a new instance using the NewRESTTime function.
type RESTTime interface {
	Name() string
	URIPath() string
	GetTimeZone() (*TimeZone, error)
	GetTimeZoneResponse() (*http.Response, error)
	GetClockServer() (*ClockServer, error)
	GetClockServerResponse() (*http.Response, error)
	GetNTPClockServer(server string) (*NTPClockServer, error)
	GetNTPClockServerResponse(server string) (*http.Response, error)
	GetNTPClockServerKey(index int32) (*NTPClockServerKey, error)
	GetNTPClockServerKeyResponse(index int32) (*http.Response, error)
}

// RESTTime interface implementation
type restTimeImpl struct {
	TimeZone
	config *api_interface.RESTConfig
}

func NewRESTTime(cfg *api_interface.RESTConfig) RESTTime {
	return &restTimeImpl{
		config: cfg,
	}
}

func (time restTimeImpl) Name() string {
	return "brocade-time"
}

func (time restTimeImpl) URIPath() string {
	return "/running/brocade-time"
}

func (time *restTimeImpl) GetTimeZone() (*TimeZone, error) {
	req, err := http.NewRequest("GET", time.config.Host()+time.config.BaseURI()+"/"+time.URIPath()+"/time-zone", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, time.config)
	if errs != nil {
		return nil, errs
	}
	var timezone TimeZone
	ct := time.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &timezone)

	return &timezone, err
}

func (c *restTimeImpl) GetTimeZoneResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+c.URIPath()+"/time-zone", nil)
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

func (time *restTimeImpl) GetClockServer() (*ClockServer, error) {
	req, err := http.NewRequest("GET", time.config.Host()+time.config.BaseURI()+"/"+time.URIPath()+"/clock-server", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, time.config)
	if errs != nil {
		return nil, errs
	}
	var cs ClockServer
	ct := time.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &cs)

	return &cs, err
}

func (c *restTimeImpl) GetClockServerResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+c.URIPath()+"/clock-server", nil)
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

func (time *restTimeImpl) GetNTPClockServer(server string) (*NTPClockServer, error) {
	if server != "" {
		server = "/" + server
	}
	req, err := http.NewRequest("GET", time.config.Host()+time.config.BaseURI()+"/"+time.URIPath()+"/time-zone/"+server, nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, time.config)
	if errs != nil {
		return nil, errs
	}
	var cs NTPClockServer
	ct := time.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &cs)

	return &cs, err
}

func (c *restTimeImpl) GetNTPClockServerResponse(server string) (*http.Response, error) {
	if server != "" {
		server = "/" + server
	}
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+c.URIPath()+"/ntp-clock-server/"+server, nil)
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

func (time *restTimeImpl) GetNTPClockServerKey(index int32) (*NTPClockServerKey, error) {
	ixStr := ""
	if index > 0 {
		ixStr = "/" + strconv.Itoa(int(index))
	}
	req, err := http.NewRequest("GET", time.config.Host()+time.config.BaseURI()+"/"+time.URIPath()+"/time-zone/"+ixStr, nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, time.config)
	if errs != nil {
		return nil, errs
	}
	var csKey NTPClockServerKey
	ct := time.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &csKey)

	return &csKey, err
}

func (c *restTimeImpl) GetNTPClockServerKeyResponse(index int32) (*http.Response, error) {
	ixStr := ""
	if index > 0 {
		ixStr = "/" + strconv.Itoa(int(index))
	}
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+c.URIPath()+"/ntp-clock-server/"+ixStr, nil)
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
