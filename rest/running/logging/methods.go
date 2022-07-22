package logging

import (
	"fmt"
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
)

type RESTLogging interface {
	Name() string
	URIPath() string
	GetLogging() (*BrocadeLogging, error)
	GetAudit() (*Audit, error)
	GetSyslogServer() ([]SyslogServer, error)
	GetRASLog() ([]RASLog, error)
	GetRASLogModule() ([]RASLogModule, error)
	GetLogQuietControl() ([]LogQuietControl, error)
	GetLogSetting() (*LogSettings, error)
}

type restLoggingImpl struct {
	config *api_interface.RESTConfig
}

func (r *restLoggingImpl) Name() string {
	return "brocade-logging"
}

func NewRESTLogging(config *api_interface.RESTConfig) RESTLogging {
	return &restLoggingImpl{
		config: config,
	}
}

func (r restLoggingImpl) URIPath() string {
	return "/running/brocade-logging"
}

func (r *restLoggingImpl) GetLogging() (*BrocadeLogging, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/logging", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var logging BrocadeLogging
	ct := r.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &logging)

	fmt.Println("Response struct", logging)
	fmt.Println("Response body", string(responseBytes))
	return &logging, err
}

func (r *restLoggingImpl) GetAudit() (*Audit, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/audit", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var audit Audit
	ct := r.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &audit)

	fmt.Println("Response struct", audit)
	fmt.Println("Response body", string(responseBytes))
	return &audit, err
}

func (r *restLoggingImpl) GetSyslogServer() ([]SyslogServer, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/syslog-server", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var ss []SyslogServer
	ct := r.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &ss)

	fmt.Println("Response struct", ss)
	fmt.Println("Response body", string(responseBytes))
	return ss, err
}

func (r *restLoggingImpl) GetRASLog() ([]RASLog, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/raslog", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var ras []RASLog
	ct := r.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &ras)

	fmt.Println("Response struct", ras)
	fmt.Println("Response body", string(responseBytes))
	return ras, err
}

func (r *restLoggingImpl) GetRASLogModule() ([]RASLogModule, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/raslog-module", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var ras []RASLogModule
	ct := r.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &ras)

	fmt.Println("Response struct", ras)
	fmt.Println("Response body", string(responseBytes))
	return ras, err
}

func (r *restLoggingImpl) GetLogQuietControl() ([]LogQuietControl, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/log-quiet-control", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var lqc []LogQuietControl
	ct := r.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &lqc)

	fmt.Println("Response struct", lqc)
	fmt.Println("Response body", string(responseBytes))
	return lqc, err
}

func (r *restLoggingImpl) GetLogSetting() (*LogSettings, error) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/log-setting", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var ls LogSettings
	ct := r.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &ls)

	fmt.Println("Response struct", ls)
	fmt.Println("Response body", string(responseBytes))
	return &ls, err
}
