package maps

import (
	"net/http"

	"github.com/projectbadger/brocade-go/rest/api_interface"
)

type RESTMAPS interface {
	Name() string
	URIPath() string
	GetSwitchStatusPolicyReport() (*SwitchStatusPolicyReport, error)
	GetSwitchStatusPolicyReportResponse() (*http.Response, error)
}

type restMapsImpl struct {
	config *api_interface.RESTConfig
}

func NewMaps(cfg *api_interface.RESTConfig) *restMapsImpl {
	return &restMapsImpl{
		config: cfg,
	}
}

func (c restMapsImpl) Name() string {
	return "brocade-maps"
}

func (c restMapsImpl) URIPath() string {
	return "/running/brocade-maps"
}

func (maps *restMapsImpl) GetSwitchStatusPolicyReport() (*SwitchStatusPolicyReport, error) {
	req, err := http.NewRequest("GET", maps.config.Host()+maps.config.BaseURI()+"/"+maps.URIPath()+"/switch-status-policy-report", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, maps.config)
	if errs != nil {
		return nil, errs
	}
	var sspr SwitchStatusPolicyReport
	ct := maps.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &sspr)

	return &sspr, err
}

func (maps *restMapsImpl) GetSwitchStatusPolicyReportResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", maps.config.Host()+maps.config.BaseURI()+maps.URIPath()+"/switch-status-policy-report", nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, maps.config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}

func (maps *restMapsImpl) GetGroup() (*Group, error) {
	req, err := http.NewRequest("GET", maps.config.Host()+maps.config.BaseURI()+"/"+maps.URIPath()+"/group", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, maps.config)
	if errs != nil {
		return nil, errs
	}
	var group Group
	ct := maps.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &group)

	return &group, err
}

func (maps *restMapsImpl) GetGroupResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", maps.config.Host()+maps.config.BaseURI()+maps.URIPath()+"/group", nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, maps.config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}

func (maps *restMapsImpl) GetRule() (*Rule, error) {
	req, err := http.NewRequest("GET", maps.config.Host()+maps.config.BaseURI()+"/"+maps.URIPath()+"/rule", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, maps.config)
	if errs != nil {
		return nil, errs
	}
	var rule Rule
	ct := maps.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &rule)

	return &rule, err
}

func (maps *restMapsImpl) GetRuleResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", maps.config.Host()+maps.config.BaseURI()+maps.URIPath()+"/rule", nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, maps.config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}

func (maps *restMapsImpl) GetDashboardRule() (*DashboardRule, error) {
	req, err := http.NewRequest("GET", maps.config.Host()+maps.config.BaseURI()+"/"+maps.URIPath()+"/dashboard-rule", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, maps.config)
	if errs != nil {
		return nil, errs
	}
	var dashboardRule DashboardRule
	ct := maps.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &dashboardRule)

	return &dashboardRule, err
}

func (maps *restMapsImpl) GetDashboardRuleResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", maps.config.Host()+maps.config.BaseURI()+maps.URIPath()+"/dashboard-rule", nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, maps.config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}

func (maps *restMapsImpl) GetMAPSPolicy() (*MAPSPolicy, error) {
	req, err := http.NewRequest("GET", maps.config.Host()+maps.config.BaseURI()+"/"+maps.URIPath()+"/maps-policy", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, maps.config)
	if errs != nil {
		return nil, errs
	}
	var policy MAPSPolicy
	ct := maps.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &policy)

	return &policy, err
}

func (maps *restMapsImpl) GetMAPSPolicyResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", maps.config.Host()+maps.config.BaseURI()+maps.URIPath()+"/maps-policy", nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, maps.config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}

func (maps *restMapsImpl) GetPausedCFG() (*PausedCfg, error) {
	req, err := http.NewRequest("GET", maps.config.Host()+maps.config.BaseURI()+"/"+maps.URIPath()+"/paused-cfg", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, maps.config)
	if errs != nil {
		return nil, errs
	}
	var config PausedCfg
	ct := maps.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &config)

	return &config, err
}

func (maps *restMapsImpl) GetPausedCFGResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", maps.config.Host()+maps.config.BaseURI()+maps.URIPath()+"/paused-cfg", nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, maps.config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}

func (maps *restMapsImpl) GetSystemResources() (*SystemResources, error) {
	req, err := http.NewRequest("GET", maps.config.Host()+maps.config.BaseURI()+"/"+maps.URIPath()+"/system-resources", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, maps.config)
	if errs != nil {
		return nil, errs
	}
	var resources SystemResources
	ct := maps.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &resources)

	return &resources, err
}

func (maps *restMapsImpl) GetSystemResourcesResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", maps.config.Host()+maps.config.BaseURI()+maps.URIPath()+"/system-resources", nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, maps.config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}

func (maps *restMapsImpl) GetMonitoringSystemMatrix() (*MonitoringSystemMatrix, error) {
	req, err := http.NewRequest("GET", maps.config.Host()+maps.config.BaseURI()+"/"+maps.URIPath()+"/monitoring-system-matrix", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, maps.config)
	if errs != nil {
		return nil, errs
	}
	var msMatrix MonitoringSystemMatrix
	ct := maps.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &msMatrix)

	return &msMatrix, err
}

func (maps *restMapsImpl) GetMonitoringSystemMatrixResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", maps.config.Host()+maps.config.BaseURI()+maps.URIPath()+"/monitoring-system-matrix", nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, maps.config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}

func (maps *restMapsImpl) GetMAPSConfig() (*MAPSConfig, error) {
	req, err := http.NewRequest("GET", maps.config.Host()+maps.config.BaseURI()+"/"+maps.URIPath()+"/maps-config", nil)
	if err != nil {
		return nil, err
	}
	_, responseBytes, errs := api_interface.GetResponse(req, maps.config)
	if errs != nil {
		return nil, errs
	}
	var config MAPSConfig
	ct := maps.config.ContentType()
	err = ct.UnmarshalResponse(responseBytes, &config)

	return &config, err
}

func (maps *restMapsImpl) GetMAPSConfigResponse() (*http.Response, error) {
	req, err := http.NewRequest("GET", maps.config.Host()+maps.config.BaseURI()+maps.URIPath()+"/maps-config", nil)
	if err != nil {
		return nil, err
	}
	resp, err := api_interface.GetHTTPResponse(req, maps.config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}
