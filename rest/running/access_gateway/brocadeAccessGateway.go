package access_gateway

import (
	"encoding/xml"
)

// Configuration and runtime information of the Access Gateway.
//
// Documentation source: https://docs.broadcom.com/doc/FOS-82X-REST-API-RM - page 79 (20220701)
type BrocadeAccessGateway struct {
	XMLName xml.Name `json:"-" xml:"brocade-access-gateway"`
	// The port group configuration.
	// The port group defines a set of N_Ports to be
	// included in the Port Grouping policy.
	// The factory default port group is "0",
	// which includes all N_Ports.
	// The default port group cannot be removed or renamed.
	// This parameter is available only when
	// the Port Grouping policy is enabled
	// (port-group-policy-enabled = 1)
	PortGroup PortGroup  `json:"port-group,omitempty" xml:"port-group"`
	NPortMap  []NPortMap `json:"n-port-map,omitempty" xml:"n-port-map>n-port"`
	// The N_Port logon information and the attached switch
	// details. This parameter is available only when the
	// port is online (online-status = 1)
	NPortInfo NPortInfo `json:"n-port-info,omitempty" xml:"n-port-info"`
	// List of all F_Ports present on the Access Gateway.
	FPortList FPortList `json:"f-port-list,omitempty" xml:"f-port-list"`
	// The Access Gateway policy configuration.
	// To enable or disable any of the AG policies, the
	// switch must be in a disabled state.
	Policy Policy `json:"policy,omitempty" xml:"policy"`
	// The N_Port-related configuration parameters.
	// Note that the port group policy must be enabled
	// (port-group-policy-enabled = 1).
	NPortSettings NPortSettings `json:"n-port-settings,omitempty" xml:"n-port-settings"`
	// A list of devices logged on to the Access Gateway and
	// their F_Port, N_Port, and FCID mapping.
	DeviceList []DeviceList `json:"device-list,omitempty" xml:"device-list"`
}

func (b *BrocadeAccessGateway) URI() string {
	return "/running/brocade-access-gateway"
}

type PortGroup struct {
	XMLName xml.Name `json:"-" xml:"port-group"`
	// The port group ID.
	// The maximum number of port groups
	// that can be created is equal to the total number of
	// ports available on the given platform.
	// For instance, on a 64-port platform, there can be
	// 64 port groups with the port-group ID ranging
	// from 0 to 63
	PortGroupID uint8 `json:"port-group-id,omitempty" xml:"port-group-id"`
	// The port group name.
	// 1 to 64 alphanumeric characters.
	// The name must be alphanumeric.
	// The default port group name uses the format
	// pg<port-group-id>
	PortGroupName string `json:"port-group-name,omitempty" xml:"port-group-name"`
	// List of N_Ports that are in the port group.
	// he port group must contain at least one N_Port.
	PortGroupNPorts []string `json:"port-group-n-ports,omitempty" xml:"port-group-n-ports>n-port"`
	// List of F_Ports that are in the port group.
	// This parameter is available only when load-balancing
	// mode is enabled (load-balancing-mode-enabled = 1)
	// for the port group.
	PortGroupFPorts []string `json:"port-group-f-ports,omitempty" xml:"port-group-f-ports>f-port"`
	// The mode configured for the port group.
	PortGroupMode PortGroupMode `json:"port-group-mode,omitempty" xml:"port-group-mode"`
}

func (b *PortGroup) URI() string {
	return "/rest/running/brocade-access-gateway/port-group"
}

// The mode configured for the port group.
type PortGroupMode struct {
	XMLName xml.Name `json:"-" xml:"port-group-mode"`
	// Enables or disables load-balancing mode for the
	// specified port group
	LoadBalancingModeEnabled uint8 `json:"load-balancing-mode-enabled,omitempty" xml:"load-balancing-mode-enabled"`
	// Enables or disables multiple fabric name
	// monitor mode for the specified port group.
	MultipleFabricNameMonitoringModeEnabled uint8 `json:"multiple-fabric-name-monitoring-mode-enabled,omitempty" xml:"multiple-fabric-name-monitoring-mode-enabled"`
}

// Defines the N_Port to F_Port mappings
type NPortMap struct {
	XMLName xml.Name `json:"-" xml:"n-port-map"`
	// The slot and port number of the specified port
	// in the format slot/port.
	// Enables mapping F_Ports to an N_Port.
	NPort string `json:"n-port,omitempty" xml:"n-port"`
	// Enables or disables failover for an N_Port.
	// 0 = Failover disabled. 1 = Failover enabled.
	FailoverEnabled uint8 `json:"failover-enabled,omitempty" xml:"failover-enabled"`
	// Enables or disables failback for an N_Port.
	// 0 = Failback disabled. 1 = Failback enabled.
	FallbackEnabled uint8 `json:"fallback-enabled,omitempty" xml:"fallback-enabled"`
	// Whether the N_Port is online or offline.
	// 0 = Offline. 1 = Online.
	OnlineStatus uint8 `json:"online-status,omitempty" xml:"online-status"`
	// The reliable status of the N_Port.
	// This parameter is available only when the
	// reliability counter is enabled
	// (reliability-counter ! = 0).
	ReliableStatus uint8 `json:"reliable-status,omitempty" xml:"reliable-status"`
}

// The N_Port logon information and the attached switch
// details. This parameter is available only when the
// port is online (online-status = 1)
type NPortInfo struct {
	XMLName xml.Name `json:"-" xml:"n-port-info"`
	// The WWN of the fabric attached to the N_Port.
	AttachedFabricName string `json:"attached-fabric-name,omitempty" xml:"attached-fabric-name"`
	// The WWN of the port attached to the N_Port.
	AttachedPortWWN string `json:"attached-port-wwn,omitempty" xml:"attached-port-wwn"`
	// The FCID of the N_Port.
	NPortFCID string `json:"n-port-fcid,omitempty" xml:"n-port-fcid"`
	// The ASCII name assigned to the switch by the
	// administrator
	AttachedSwitchUserFriendlyName string `json:"attached-switch-user-friendly-name,omitempty" xml:"attached-switch-user-friendly-name"`
	// The fabric switch port number of the port attached
	// to the N_Port.
	AttachedSwitchFPort string `json:"attached-switch-f-port,omitempty" xml:"attached-switch-f-port"`
	// The out-of-band IP address of the attached switch.
	AttachedSwitchIPAddress string `json:"attached-switch-ip-address,omitempty" xml:"attached-switch-ip-address"`
	// The F_Port to N_Port mappings.
	// List of F_Ports that are mapped to the N_Port.
	//There must be at least one F_Port mapped to an
	//N_Port. Once you enable Access Gateway mode in the
	// switch, F_Ports are mapped to the default N_Port.
	// You must clear the existing mapping before you map
	// an F_Port to another N_Port. Note that configured
	// and static port mapping cannot overlap.
	//
	// The slot and port number of the specified port in
	// the format slot/port.
	ConfiguredFPortList []string `json:"configured-f-port-list" xml:"configured-f-port-list>f-port"`
	// The static F_Port to N_Port mappings.
	// List of F_Ports that are statically mapped to the
	// N_Port. There must be at least one F_Port statically
	// mapped to an N_Port.
	// Note that configured and static port mapping
	// cannot overlap.
	//
	// The slot and port number of the specified port
	// in the format slot/port.
	StaticFPortList []string `json:"static-f-port-list" xml:"static-f-port-list>f-port"`
}

// List of all F_Ports present on the Access Gateway.
type FPortList struct {
	XMLName xml.Name `json:"-" xml:"f-port-list"`
	// The F_Port for which information is being fetched.
	//
	// The slot and port number of the specified port in
	// the format slot/port
	FPort string `json:"f-port,omitempty" xml:"f-port"`
	// Whether the F_Port is online or offline.
	//
	// 0 = Offline. 1 = Online.
	OnlineStatus uint8 `json:"online-status,omitempty" xml:"online-status"`
	// The F_Port login information. This parameter is
	// available only when the port is online
	// (online-status = 1).
	FPortInfo FPortInfo `json:"f-port-info,omitempty" xml:"f-port-info"`
}

// The F_Port login information. This parameter is
// available only when the port is online
// (online-status = 1).
type FPortInfo struct {
	XMLName xml.Name `json:"-" xml:"f-port-info"`
	// The N_Port to which this F_Port is mapped.
	//
	// The slot and port number of the specified port in
	// the format slot/port
	NPort string `json:"n-port,omitempty" xml:"n-port"`
	// The login exceeded state of the F_Port.
	//
	// 0 = Login limit not exceeded. 1 = Login limit exceeded.
	LoginExceeded uint8 `json:"login-exceeded,omitempty" xml:"login-exceeded"`
}

// The Access Gateway policy configuration.
// To enable or disable any of the AG policies, the switch
// must be in a disabled state.
type Policy struct {
	XMLName xml.Name `json:"-" xml:"policy"`
	// Enables or disables the port group policy.
	// Note that the auto policy must be disabled
	// (auto-policy-enabled = 0). To modify the policy
	// configuration, the switch must be in a disabled
	// state.
	PortGroupPolicyEnabled uint8 `json:"port-group-policy-enabled,omitempty" xml:"port-group-policy-enabled"`
	// Enables or disables the auto policy.
	// Note that the port group policy must be disabled
	// (port-group-policy-enabled = 0).
	// To modify the policy configuration, the switch must
	// be in a disabled state.
	AutoPolicyEnabled uint8 `json:"auto-policy-enabled,omitempty" xml:"auto-policy-enabled"`
}

// The N_Port-related configuration parameters.
// Note that the port group policy must be enabled
// (port-group-policy-enabled = 1).
type NPortSettings struct {
	XMLName xml.Name `json:"-" xml:"n-port-settings"`
	// Reliability counter configuration for N_Ports.
	//
	// 0 = Disabled; 10-100 = Enabled.
	ReliabilityCounter uint64 `json:"reliability-counter,omitempty" xml:"reliability-counter"`
}

// A list of devices logged on to the Access Gateway
// and their F_Port, N_Port, and FCID mapping.
type DeviceList struct {
	XMLName xml.Name `json:"-" xml:"device-list"`
	// The port world wide name of the connected device.
	WWN string `json:"wwn,omitempty" xml:"wwn"`
	// The fibre channel ID (FCID) of the connected device.
	FCID string `json:"fcid,omitempty" xml:"fcid"`
	// The F_Port to which the device is connected.
	//
	// The slot and port number of the specified port in
	// the format slot/port.
	FPort string `json:"f-port,omitempty" xml:"f-port"`
	// The N_Port through which the device logged on.
	//
	// The slot and port number of the specified port in
	// the format slot/port.
	NPort string `json:"n-port,omitempty" xml:"n-port"`
}
