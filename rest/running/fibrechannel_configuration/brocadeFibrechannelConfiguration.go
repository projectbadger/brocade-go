package fibrechannel_configuration

import "encoding/xml"

// This module enables you to configure a switch or
// director.
//
// Documentation source: https://docs.broadcom.com/doc/FOS-82X-REST-API-RM - page 129 (20220701)
type BrocadeFibrechannelConfiguration struct {
	XMLName xml.Name `json:"-" xml:"brocade-fibrechannel-configuration"`
	// Provides switch configuration parameters.
	SwitchConfiguration SwitchConfiguration `json:"switch-configuration" xml:"switch-configuration"`
	// Provides parameters for F_Port login settings.
	// Note that the switch must be in a disabled state to
	// configure F_Port login settings
	// (fibrechannel-switch:enabled-state = 3).
	FPortLoginSettings FPortLoginSettings `json:"f-port-login-settings" xml:"f-port-login-settings"`
	// Provides the port configuration parameters.
	PortConfiguration PortConfiguration `json:"port-configuration" xml:"port-configuration"`
	// Provides zoning configuration parameters.
	ZoneConfiguration ZoneConfiguration `json:"zone-configuration" xml:"zone-configuration"`
	// Provides attributes applicable for a switch.
	Fabric Fabric `json:"fabric" xml:"fabric"`
}

type SwitchConfiguration struct {
	XMLName xml.Name `json:"-" xml:"switch-configuration"`
	// Enables or disables trunking on all ports in the
	// logical switch.
	TrunkEnabled bool `json:"trunk-enabled,omitempty" xml:"trunk-enabled"`
	// Indicates whether WWN-based persistent PID is
	// enabled on a switch.
	WWNPortIdNode bool `json:"wwn-port-id-node,omitempty" xml:"wwn-port-id-node"`
	// The maximum time, in milleseconds, that a frame can
	// wait after it is received on the ingress port and
	// before it is delivered to the egress port.
	EdgeHoldTime uint16 `json:"edge-hold-time,omitempty" xml:"edge-hold-time"`
	// The area mode. Area mode has a different meaning
	// (see Value below) in a default switch and
	// non-default, non-FICON switch. Note that the switch
	// must be in disabled state for this setting
	// (fibrechannel-switch:enabled-state = 3).
	AreaMode uint8 `json:"area-mode,omitempty" xml:"area-mode"`
}

type FPortLoginSettings struct {
	XMLName xml.Name `json:"-" xml:"f-port-login-settings"`
	// he maximum number of allowed logins for the switch.
	MaxLogins uint16 `json:"max-logins,omitempty" xml:"max-logins"`
	// The maximum number of fabric logins (FLOGIs)
	// allowed, per second, on a switch.
	MaxFlogiRatePerSwitch uint16 `json:"max-flogi-rate-per-switch,omitempty" xml:"max-flogi-rate-per-switch"`
	// The interval setting, in milliseconds, for the rate
	// at which F_Ports are enabled.
	StageInterval uint16 `json:"stage-interval,omitempty" xml:"stage-interval"`
	// The number of logins allowed before staging. This
	// parameter, if nonzero, enables staging of FDISC
	// logins by rejecting the FDISC requests with logical
	// busy, when the requests are more than the number of
	// configured logins per second.
	FreeFdisc uint16 `json:"free-fdisc,omitempty" xml:"free-fdisc"`
	// The precedence for login when two devices with same
	// the port WWN (PWWN) compete for login. All modes are
	// for NPIV and non-NPIV F_Ports.
	//
	//  0 = The first login takes precedence over the second login.
	//  1 = The second login overrides the first login.
	//  2 = For FDISC, the second FDISC login takes precedence. For FLOGI, the first FLOGI takes precedence.
	EnforceLogin uint16 `json:"enforce-login,omitempty" xml:"enforce-login"`
	// The maximum number of logins allowed, per second, on
	// a given port. If the number is exceeded, the port is
	// fenced.
	MaxFlogiRatePerPort uint16 `json:"max-flogi-rate-per-port,omitempty" xml:"max-flogi-rate-per-port"`
}

type PortConfiguration struct {
	XMLName xml.Name `json:"-" xml:"port-configuration"`
	// The current port name mode
	// (off, default, fdmi or dynamic).
	PortnameMode string `json:"portname-mode,omitempty" xml:"portname-mode"`
	// The format of a dynamic port name is composed of
	// fields that are mapped using the following
	// characters:
	//  S = Switch name
	//  T = Port type
	//  I = Port index (note that I and C are mutually exclusive)
	//  C = Slot number/port number (only applicable on the chassis)
	//  A = Alias name
	//  F = FDMI host name
	//  R = Remote switch name
	// Multiple unique fields may be specified, separated
	// by hyphens (-), periods (.), or underscores (_).
	// Fields may not be repeated, and field separators
	// must be the same. The switch must be an FC switch in
	// Native mode.
	DynamicPortnameFormat string `json:"dynamic-portname-format,omitempty" xml:"dynamic-portname-format"`
	// The port has D_Port capability, but it is not
	// explicitly configured. Enabling a dynamic D_Port
	// switch-wide configuration forces the ports on that
	// switch or chassis to respond to D_Port requests from
	// the other end of the connection. Basically, the port
	// responds to a remote port request to change its mode
	// to D_Port mode and run diagnostic tests
	// automatically.
	DynamicDPortEnabled bool `json:"dynamic-d-port-enabled,omitempty" xml:"dynamic-d-port-enabled"`
	// The port has D_Port capability, but it is not
	// explicitly configured. Enabling an on-demand D_Port
	// switch-wide configuration forces the ports on that
	// switch or chassis to respond to an internal request
	// within the switch as a result of certain event
	// (slot power off or on, persistent disable or enable,
	// and so on). Basically, the switch responds to an
	// internal request to change a port mode to D_Port
	// mode, and run diagnostic tests automatically. The
	// D_Ports change to normal port mode after successful
	// completion of the tests.
	OnDemandDPortEnabled bool `json:"on-demand-d-port-enabled,omitempty" xml:"on-demand-d-port-enabled"`
}

type ZoneConfiguration struct {
	XMLName xml.Name `json:"-" xml:"zone-configuration"`
	// Indicates whether node name checking for zoning is
	// enabled.
	NodeNameZoningEnabled bool `json:"node-name-zoning-enabled,omitempty" xml:"node-name-zoning-enabled"`
}

type Fabric struct {
	XMLName xml.Name `json:"-" xml:"fabric"`
	// Indicates whether Insistent Domain ID (IDID) is
	// enabled.
	InsistentDomainIDEnabled bool `json:"insistent-domain-id-enabled,omitempty" xml:"insistent-domain-id-enabled"`
}
