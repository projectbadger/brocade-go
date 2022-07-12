package fibrechannel_logical_switch

import "encoding/xml"

// Information of all logical switches in a operational
// chassis and logical switch configuration.
//
// Documentation source: https://docs.broadcom.com/doc/FOS-82X-REST-API-RM - page 140 (20220701)
type BrocadeFibrechannelLogicalSwitch struct {
	XMLName                   xml.Name                    `json:"-" xml:"brocade-fibrechannel-logical-switch"`
	FibrechannelLogicalSwitch []FibrechannelLogicalSwitch `json:"fibrechannel-logical-switch" xml:"brocade-fibrechannel-logical-switch>fibrechannel-logical-switch"`
}

// Provides logical switch state parameters of all
// configured logical switches.
type FibrechannelLogicalSwitch struct {
	XMLName xml.Name `json:"-" xml:"fibrechannel-logical-switch"`
	// The virtual fabric identification (VFID) of the
	// logical switch.
	FabricId uint32 `json:"fabric-id" xml:"fabric-id"`
	// The switch WWN, which is a unique numeric identifier
	// for the switch assigned by the chassis.
	SwitchWWN string `json:"switch-wwn" xml:"switch-wwn"`
	// The logical switch that is the base switch in the
	// chassis. You can configure the base switch when
	// creating a logical switch. This parameter is
	// available only when XISL and ficon are disabled.
	BaseSwitchEnabled uint8 `json:"base-switch-enabled" xml:"base-switch-enabled"`
	// The logical switch is the default switch in the
	// chassis.
	DefaultSwitchStatus uint8 `json:"default-switch-status" xml:"default-switch-status"`
	// Enables logical ISLs (LISLs) on a logical switch.
	// You can only disable LISLs when you create a logical
	// switch. This parameter is available only when the
	// base switch is disabled
	// (base-switch-enabled=0).
	LogicalISLEnabled uint8 `json:"logical-isl-enabled" xml:"logical-isl-enabled"`
	// Indicates the FICON mode. This parameter is
	// available only when the base switch is disabled
	// (base-switch-enabled=0) and default switch status is
	// disabled (default-switch-status=0). You cannot
	// disable FICON mode in a logical switch.
	FiconModeEnabled uint8 `json:"ficon-mode-enabled" xml:"ficon-mode-enabled"`
	// A list of logical switch port interface (fibre
	// channel and flex) names.
	PortMemberList []string `json:"port-member-list" xml:"port-member-list>port-member"`
	// A list of logical switch port interface Gigabit
	// Ethernet (GE) names.
	GePortMemberList []string `json:"ge-port-member-list" xml:"ge-port-member-list>port-member"`
}
