package fibrechannel_switch

import "encoding/xml"

// Switch state parameters.
//
// Documentation source: https://docs.broadcom.com/doc/FOS-82X-REST-API-RM - page 150 (20220701)
type BrocadeFibrechannelSwitch struct {
	XMLName            xml.Name             `json:"-" xml:"brocade-fibrechannel-switch"`
	FibrechannelSwitch []FibrechannelSwitch `json:"fibrechannel-switch" xml:"fibrechannel-switch"`
}

// Switch identification parameters.
type FibrechannelSwitch struct {
	XMLName xml.Name `json:"-" xml:"fibrechannel-switch"`
	// The switch world wide name (WWN).
	Name string `json:"name,omitempty" xml:"name"`
	// Domain ID of the switch.
	DomainId string `json:"domain-id,omitempty" xml:"domain-id"`
	// The ASCII name assigned to the switch by the
	// administrator.
	UserFriendlyName string `json:"user-friendly-name,omitempty" xml:"user-friendly-name"`
	// This parameter is deprecated.
	// Use the fcid-hex parameter. The destination ID
	// (D_ID) of the switch (decimal format).
	Fcid string `json:"fcid,omitempty" xml:"fcid"`
	// The destination ID (D_ID) of the switch
	// (hexidecimal string format).
	FcidHex string `json:"fcid-hex,omitempty" xml:"fcid-hex"`
	// The Virtual Fabric identifier (VFID) of the
	// logical switch.
	VfId int16 `json:"vf-id,omitempty" xml:"vf-id"`
	// Indicates the principal switch in the fabric.
	Principal uint8 `json:"principal,omitempty" xml:"principal"`
	// This parameter is deprecated.
	// Use the is-enabled-state leaf to configure and the
	// operational-status leaf to obtain the current state
	// of the switch.
	//
	// The current state of the switch.
	EnabledState uint32 `json:"enabled-state,omitempty" xml:"enabled-state"`
	// The current state of the switch.
	IsEnabledState bool `json:"is-enabled-state,omitempty" xml:"is-enabled-state"`
	// The current state of the switch.
	OperationalStatus uint32 `json:"operational-status,omitempty" xml:"operational-status"`
	// The period of time elapsed since the last reboot of the specified switch (UNIX).
	UpTime int `json:"up-time,omitempty" xml:"up-time"`
	// The model name of the switch.
	Model string `json:"model,omitempty" xml:"model"`
	// The firmware version running on the switch.
	FirmwareVersion string `json:"firmware-version,omitempty" xml:"firmware-version"`
	// A list of IP addresses assigned to the Ethernet port
	// on the switch.
	IpAddress []string `json:"ip-address,omitempty" xml:"ip-address>ip-address"`
	// A list of static gateway IPv4 and IPv6 addresses for
	// the switch IP address.
	IpStaticGatewayList []string `json:"ip-static-gateway-list,omitempty" xml:"ip-static-gateway-list>ip-static-gateway"`
	// The IPv4 subnet mask of the switch IP network.
	SubnetMask string `json:"subnet-mask,omitempty" xml:"subnet-mask"`
	// The DNS domain name of the switch.
	// Note that leaving this field empty clears the DNS
	// domain name.
	DomainName string `json:"domain-name,omitempty" xml:"domain-name"`
	// A list of DNS server addresses which can handle the
	// mapping of the domain names to the IP address of an
	// internet resource which is needed by various
	// networking protocols. Brocade FC switches support a
	// maximum of two DNS servers.
	DnsServers []string `json:"dns-servers,omitempty" xml:"dns-servers>dns-server"`
	// The human-readable name for the fabric.
	FabricUserFriendlyName string `json:"fabric-user-friendly-name,omitempty" xml:"fabric-user-friendly-name"`
	// Indicates the Access Gateway (AG) mode capability
	// and enablement state. A switch is capable of AG mode
	// support when this value is not zero.
	AgMode uint8 `json:"ag-mode,omitempty" xml:"ag-mode"`
	// The text that displays during the log on process.
	Banner string `json:"banner,omitempty" xml:"banner"`
}
