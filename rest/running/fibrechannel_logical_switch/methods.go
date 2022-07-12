package fibrechannel_logical_switch

import "brocade/rest/api_interface"

type RESTFibrechannelLogicalSwitch interface {
	Name() string
	URIPath() string
}

type restFibrechannelLogicalSwitchImpl struct {
	config *api_interface.RESTConfig
}

func NewFibrechannelLogicalSwitch(cfg *api_interface.RESTConfig) RESTFibrechannelLogicalSwitch {
	return &restFibrechannelLogicalSwitchImpl{
		config: cfg,
	}
}

func (fls restFibrechannelLogicalSwitchImpl) Name() string {
	return "fibrechannel-logical-switch"
}

func (fls restFibrechannelLogicalSwitchImpl) URIPath() string {
	return "/running/brocade-fibrechannel-logical-switch"
}
