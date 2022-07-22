package logging

import (
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
	"github.com/projectbadger/brocade-go/rest/errors"
)

// RESTLogging describes an interface for interacting with the
// *logging* module.
// Fetch a new instance using the NewRESTLogging function.
type RESTLogging interface {
	Name() string
	URIPath() string
	GetLogging() (*BrocadeLogging, errors.BrocadeErr)
	GetAudit() (*Audit, errors.BrocadeErr)
	GetSyslogServer() ([]SyslogServer, errors.BrocadeErr)
	GetRASLog() ([]RASLog, errors.BrocadeErr)
	GetLogQuietControl() ([]LogQuietControl, errors.BrocadeErr)
	GetRASLogModule() ([]RASLogModule, errors.BrocadeErr)
	GetLogSetting() (*LogSettings, errors.BrocadeErr)
}

// RESTLogging implementation
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

func (r *restLoggingImpl) GetLogging() (*BrocadeLogging, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/logging", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var logging BrocadeLogging
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &logging)

	return &logging, errs
}

func (r *restLoggingImpl) GetAudit() (*Audit, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/audit", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var audit Audit
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &audit)

	return &audit, errs
}

func (r *restLoggingImpl) GetSyslogServer() ([]SyslogServer, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/syslog-server", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var ss []SyslogServer
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &ss)

	return ss, errs
}

func (r *restLoggingImpl) GetRASLog() ([]RASLog, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/raslog", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var ras []RASLog
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &ras)

	return ras, errs
}

func (r *restLoggingImpl) GetRASLogModule() ([]RASLogModule, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/raslog-module", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var ras []RASLogModule
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &ras)

	return ras, errs
}

func (r *restLoggingImpl) GetLogQuietControl() ([]LogQuietControl, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/log-quiet-control", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var lqc []LogQuietControl
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &lqc)

	return lqc, errs
}

func (r *restLoggingImpl) GetLogSetting() (*LogSettings, errors.BrocadeErr) {
	req, err := http.NewRequest("GET", r.config.Host()+r.config.BaseURI()+"/"+r.URIPath()+"/log-setting", nil)
	if err != nil {
		return nil, errors.NewFromErr(err)
	}
	_, responseBytes, errs := api_interface.GetResponse(req, r.config)
	if errs != nil {
		return nil, errs
	}
	var ls LogSettings
	ct := r.config.ContentType()
	errs = ct.UnmarshalResponse(responseBytes, &ls)

	return &ls, errs
}
