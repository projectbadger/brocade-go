package snmp

import (
	"net/http"
	"strconv"

	"github.com/projectbadger/brocade-go/rest/api_interface"
	"github.com/projectbadger/brocade-go/rest/errors"
)

// RESTSNMP describes an interface for interacting with the
// *snmp* module.
// Fetch a new instance using the NewRESTSNMP function.
type RESTSNMP interface {
	Name() string
	URIPath() string
	GetSystem() (*System, errors.BrocadeErr)
	GetSystemResponse() (*http.Response, error)
	GetMIBCapability(name string) (*MIBCapability, errors.BrocadeErr)
	GetMIBCapabilityResponse(name string) (*http.Response, error)
	GetTrapCapability(name string) (*TrapCapability, errors.BrocadeErr)
	GetTrapCapabilityResponse(name string) (*http.Response, error)
	GetV1Trap(index uint16) (*V1Trap, errors.BrocadeErr)
	GetV1TrapResponse(index uint16) (*http.Response, error)
	GetV3Account(index uint16) (*V3Account, errors.BrocadeErr)
	GetV3AccountResponse(index uint16) (*http.Response, error)
	GetV3Trap(index uint16) (*V3Trap, errors.BrocadeErr)
	GetV3TrapResponse(index uint16) (*http.Response, error)
	GetAccessControl(index uint16) (*AccessControl, errors.BrocadeErr)
	GetAccessControlResponse(index uint16) (*http.Response, error)
}

// RESTSNMP interface implementation
type restSNMPImpl struct {
	config *api_interface.RESTConfig
}

func NewRESTSNMP(cfg *api_interface.RESTConfig) RESTSNMP {
	return &restSNMPImpl{
		config: cfg,
	}
}

func (snmp restSNMPImpl) Name() string {
	return "brocade-snmp"
}

func (snmp restSNMPImpl) URIPath() string {
	return "/running/brocade-snmp"
}

func (snmp *restSNMPImpl) GetSystem() (*System, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", snmp.config.Host()+snmp.config.BaseURI()+"/"+snmp.URIPath()+"/system", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, snmp.config)
	if errs != nil {
		return nil, errs
	}
	var system System
	ct := snmp.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &system)

	return &system, errors.NewFromErr(err)
}

func (c *restSNMPImpl) GetSystemResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+c.URIPath()+"/system", nil)
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

func (snmp *restSNMPImpl) GetMIBCapability(name string) (*MIBCapability, errors.BrocadeErr) {
	if name != "" {
		name = "/" + name
	}
	req, err := http.NewRequest("GET", snmp.config.Host()+snmp.config.BaseURI()+"/"+snmp.URIPath()+"/mib-capability/"+name, nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, snmp.config)
	if errs != nil {
		return nil, errs
	}
	var capability MIBCapability
	ct := snmp.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &capability)

	return &capability, errors.NewFromErr(err)
}

func (c *restSNMPImpl) GetMIBCapabilityResponse(name string) (*http.Response, error) {
	if name != "" {
		name = "/" + name
	}
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+c.URIPath()+"/mib-capability/"+name, nil)
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

func (snmp *restSNMPImpl) GetTrapCapability(name string) (*TrapCapability, errors.BrocadeErr) {
	if name != "" {
		name = "/" + name
	}
	req, err := http.NewRequest("GET", snmp.config.Host()+snmp.config.BaseURI()+"/"+snmp.URIPath()+"/trap-capability/"+name, nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, snmp.config)
	if errs != nil {
		return nil, errs
	}
	var capability TrapCapability
	ct := snmp.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &capability)

	return &capability, errors.NewFromErr(err)
}

func (c *restSNMPImpl) GetTrapCapabilityResponse(name string) (*http.Response, error) {
	if name != "" {
		name = "/" + name
	}
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+c.URIPath()+"/trap-capability/"+name, nil)
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

func (snmp *restSNMPImpl) GetV1Account(index uint16) (*V1Account, errors.BrocadeErr) {
	ixStr := ""
	if index > 0 {
		ixStr = "/" + strconv.Itoa(int(index))
	}
	req, err := http.NewRequest("GET", snmp.config.Host()+snmp.config.BaseURI()+"/"+snmp.URIPath()+"/v1-account/"+ixStr, nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, snmp.config)
	if errs != nil {
		return nil, errs
	}
	var v1Account V1Account
	ct := snmp.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &v1Account)

	return &v1Account, errors.NewFromErr(err)
}

func (c *restSNMPImpl) GetV1AccountResponse(index uint16) (*http.Response, error) {
	ixStr := ""
	if index > 0 {
		ixStr = "/" + strconv.Itoa(int(index))
	}
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+c.URIPath()+"/v1-account/"+ixStr, nil)
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

func (snmp *restSNMPImpl) GetV1Trap(index uint16) (*V1Trap, errors.BrocadeErr) {
	ixStr := ""
	if index > 0 {
		ixStr = "/" + strconv.Itoa(int(index))
	}
	req, err := http.NewRequest("GET", snmp.config.Host()+snmp.config.BaseURI()+"/"+snmp.URIPath()+"/v1-trap/"+ixStr, nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, snmp.config)
	if errs != nil {
		return nil, errs
	}
	var v1Trap V1Trap
	ct := snmp.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &v1Trap)

	return &v1Trap, errors.NewFromErr(err)
}

func (c *restSNMPImpl) GetV1TrapResponse(index uint16) (*http.Response, error) {
	ixStr := ""
	if index > 0 {
		ixStr = "/" + strconv.Itoa(int(index))
	}
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+c.URIPath()+"/v1-trap/"+ixStr, nil)
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

func (snmp *restSNMPImpl) GetV3Account(index uint16) (*V3Account, errors.BrocadeErr) {
	ixStr := ""
	if index > 0 {
		ixStr = "/" + strconv.Itoa(int(index))
	}
	req, err := http.NewRequest("GET", snmp.config.Host()+snmp.config.BaseURI()+"/"+snmp.URIPath()+"/v3-account/"+ixStr, nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, snmp.config)
	if errs != nil {
		return nil, errs
	}
	var v3Account V3Account
	ct := snmp.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &v3Account)

	return &v3Account, errors.NewFromErr(err)
}

func (c *restSNMPImpl) GetV3AccountResponse(index uint16) (*http.Response, error) {
	ixStr := ""
	if index > 0 {
		ixStr = "/" + strconv.Itoa(int(index))
	}
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+c.URIPath()+"/v3-account/"+ixStr, nil)
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

func (snmp *restSNMPImpl) GetV3Trap(index uint16) (*V3Trap, errors.BrocadeErr) {
	ixStr := ""
	if index > 0 {
		ixStr = "/" + strconv.Itoa(int(index))
	}
	req, err := http.NewRequest("GET", snmp.config.Host()+snmp.config.BaseURI()+"/"+snmp.URIPath()+"/v3-trap/"+ixStr, nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, snmp.config)
	if errs != nil {
		return nil, errs
	}
	var v3Trap V3Trap
	ct := snmp.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &v3Trap)

	return &v3Trap, errors.NewFromErr(err)
}

func (c *restSNMPImpl) GetV3TrapResponse(index uint16) (*http.Response, error) {
	ixStr := ""
	if index > 0 {
		ixStr = "/" + strconv.Itoa(int(index))
	}
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+c.URIPath()+"/v3-trap/"+ixStr, nil)
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

func (snmp *restSNMPImpl) GetAccessControl(index uint16) (*AccessControl, errors.BrocadeErr) {
	ixStr := ""
	if index > 0 {
		ixStr = "/" + strconv.Itoa(int(index))
	}
	req, err := http.NewRequest("GET", snmp.config.Host()+snmp.config.BaseURI()+"/"+snmp.URIPath()+"/access-control/"+ixStr, nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, snmp.config)
	if errs != nil {
		return nil, errs
	}
	var ac AccessControl
	ct := snmp.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &ac)

	return &ac, errors.NewFromErr(err)
}

func (c *restSNMPImpl) GetAccessControlResponse(index uint16) (*http.Response, error) {
	ixStr := ""
	if index > 0 {
		ixStr = "/" + strconv.Itoa(int(index))
	}
	req, err := http.NewRequest("GET", c.config.Host()+c.config.BaseURI()+c.URIPath()+"/access-control/"+ixStr, nil)
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
