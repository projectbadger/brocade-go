package media

import (
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
)

type RESTMediaRDP interface {
	Name() string
	URIPath() string
	GetMediaRDP(name string) ([]MediaRDP, error)
	GetMediaRDPResponse(name string) (*http.Response, error)
}

type restMediaImpl struct {
	config *api_interface.RESTConfig
}

func NewRESTMedia(cfg *api_interface.RESTConfig) *restMediaImpl {
	return &restMediaImpl{
		config: cfg,
	}
}

func (media restMediaImpl) Name() string {
	return "brocade-media"
}

func (media restMediaImpl) URIPath() string {
	return "/running/brocade-media"
}

func (media *restMediaImpl) GetMediaRDP(name string) ([]MediaRDP, error) {
	if name != "" {
		name = "/" + name
	}
	req, err := http.NewRequest("GET", media.config.Host()+media.config.BaseURI()+"/"+media.URIPath()+"/media-rdp/"+name, nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, media.config)
	if errs != nil {
		return nil, errs
	}
	var mediaRDP []MediaRDP
	ct := media.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &mediaRDP)

	return mediaRDP, err
}

func (c *restMediaImpl) GetMediaRDPResponse(name string) (*http.Response, error) {
	if name != "" {
		name = "/" + name
	}
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+c.URIPath()+"/media-rdp/"+name, nil)
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

func (media *restMediaImpl) GetMediaRDPForSFP(name string) ([]MediaRDP, error) {
	if name != "" {
		name = "/" + name
	}
	req, err := http.NewRequest("GET", media.config.Host()+media.config.BaseURI()+"/"+media.URIPath()+"/media-rdp/name/"+name, nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, media.config)
	if errs != nil {
		return nil, errs
	}
	var mediaRDP []MediaRDP
	ct := media.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &mediaRDP)

	return mediaRDP, err
}

func (c *restMediaImpl) GetMediaRDPForSFPResponse(name string) (*http.Response, error) {
	if name != "" {
		name = "/" + name
	}
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+c.URIPath()+"/media-rdp/name/"+name, nil)
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
