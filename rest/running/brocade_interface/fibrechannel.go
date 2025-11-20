package brocade_interface

import (
	"encoding/xml"
)

type Fibrechannel struct {
	XMLName xml.Name `json:"-,omitempty" xml:"fibrechannel"`

	Name string `json:"name,omitempty" xml:"name"`
	WWN  string `json:"wwn,omitempty" xml:"wwn"`

	// Status & admin state
	OperationalStatus       int    `json:"operational-status,omitempty" xml:"operational-status"`
	OperationalStatusString string `json:"operational-status-string,omitempty" xml:"operational-status-string"`
	EnabledState            int    `json:"enabled-state,omitempty" xml:"enabled-state"`
	IsEnabledState          bool   `json:"is-enabled-state,omitempty" xml:"is-enabled-state"`

	// Speed
	Speed            int `json:"speed,omitempty" xml:"speed"`
	ProtocolSpeed    int `json:"protocol-speed,omitempty" xml:"protocol-speed"`
	MaxSpeed         int `json:"max-speed,omitempty" xml:"max-speed"`
	MaxProtocolSpeed int `json:"max-protocol-speed,omitempty" xml:"max-protocol-speed"`

	// Port type
	PortType       int    `json:"port-type,omitempty" xml:"port-type"`
	PortTypeString string `json:"port-type-string,omitempty" xml:"port-type-string"`

	// Basic configuration
	UserFriendlyName string `json:"user-friendly-name,omitempty" xml:"user-friendly-name"`
	AutoNegotiate    int    `json:"auto-negotiate,omitempty" xml:"auto-negotiate"`
	AutoNegotiateV2  bool   `json:"auto-negotiate-v2,omitempty" xml:"auto-negotiate-v2"`

	// Switch/Platform booleans
	ApplicationHeaderEnabled bool `json:"application-header-enabled,omitempty" xml:"application-header-enabled"`
	LockedEPortEnabled       bool `json:"locked-e-port-enabled,omitempty" xml:"locked-e-port-enabled"`
	DWDMLoSyncEnabled        bool `json:"dwdm-losync-enabled,omitempty" xml:"dwdm-losync-enabled"`
	DPortDWDMEnabled         bool `json:"d-port-dwdm-enabled,omitempty" xml:"d-port-dwdm-enabled"`

	// Feature toggles
	PersistentDisable      int  `json:"persistent-disable,omitempty" xml:"persistent-disable"`
	PersistentDisableV2    bool `json:"persistent-disable-v2,omitempty" xml:"persistent-disable-v2"`
	GPortLocked            int  `json:"g-port-locked,omitempty" xml:"g-port-locked"`
	EPortDisable           int  `json:"e-port-disable,omitempty" xml:"e-port-disable"`
	EPortDisableV2         bool `json:"e-port-disable-v2,omitempty" xml:"e-port-disable-v2"`
	EPortDisableV3         int  `json:"e-port-disable-v3,omitempty" xml:"e-port-disable-v3"`
	NPortEnabled           int  `json:"n-port-enabled,omitempty" xml:"n-port-enabled"`
	NPortEnabledV2         bool `json:"n-port-enabled-v2,omitempty" xml:"n-port-enabled-v2"`
	ExPortEnabled          int  `json:"ex-port-enabled,omitempty" xml:"ex-port-enabled"`
	NPIVEnabled            int  `json:"npiv-enabled,omitempty" xml:"npiv-enabled"`
	NPIVPPLimit            int  `json:"npiv-pp-limit,omitempty" xml:"npiv-pp-limit"`
	NPIVFLOGILogoutEnabled bool `json:"npiv-flogi-logout-enabled-v2,omitempty" xml:"npiv-flogi-logout-enabled-v2"`

	// Feature Flags
	RateLimitEnabled         int  `json:"rate-limit-enabled,omitempty" xml:"rate-limit-enabled"`
	QOSEnabled               int  `json:"qos-enabled,omitempty" xml:"qos-enabled"`
	QOSEnabledV2             bool `json:"qos-enabled-v2,omitempty" xml:"qos-enabled-v2"`
	CSCTLModeEnabled         int  `json:"csctl-mode-enabled,omitempty" xml:"csctl-mode-enabled"`
	CSCTLModeEnabledV2       bool `json:"csctl-mode-enabled-v2,omitempty" xml:"csctl-mode-enabled-v2"`
	PortAutoDisableEnabled   int  `json:"port-autodisable-enabled,omitempty" xml:"port-autodisable-enabled"`
	PortAutoDisableEnabledV2 bool `json:"port-autodisable-enabled-v2,omitempty" xml:"port-autodisable-enabled-v2"`

	// Buffer and credit mgmt
	FPortBuffers            int  `json:"f-port-buffers,omitempty" xml:"f-port-buffers"`
	EPortCredit             int  `json:"e-port-credit,omitempty" xml:"e-port-credit"`
	CreditRecoveryEnabled   int  `json:"credit-recovery-enabled,omitempty" xml:"credit-recovery-enabled"`
	CreditRecoveryEnabledV2 bool `json:"credit-recovery-enabled-v2,omitempty" xml:"credit-recovery-enabled-v2"`
	CreditRecoveryActive    int  `json:"credit-recovery-active,omitempty" xml:"credit-recovery-active"`
	CreditRecoveryActiveV2  bool `json:"credit-recovery-active-v2,omitempty" xml:"credit-recovery-active-v2"`

	// FEC
	FECEnabled         int  `json:"fec-enabled,omitempty" xml:"fec-enabled"`
	FECEnabledV2       bool `json:"fec-enabled-v2,omitempty" xml:"fec-enabled-v2"`
	FECActive          int  `json:"fec-active,omitempty" xml:"fec-active"`
	FECActiveV2        bool `json:"fec-active-v2,omitempty" xml:"fec-active-v2"`
	ViaTTSFECEnabled   int  `json:"via-tts-fec-enabled,omitempty" xml:"via-tts-fec-enabled"`
	ViaTTSFECEnabledV2 bool `json:"via-tts-fec-enabled-v2,omitempty" xml:"via-tts-fec-enabled-v2"`

	// Compression / Encryption
	CompressionConfigured   int  `json:"compression-configured,omitempty" xml:"compression-configured"`
	CompressionConfiguredV2 bool `json:"compression-configured-v2,omitempty" xml:"compression-configured-v2"`
	CompressionActive       int  `json:"compression-active,omitempty" xml:"compression-active"`
	CompressionActiveV2     bool `json:"compression-active-v2,omitempty" xml:"compression-active-v2"`

	EncryptionEnabled   int  `json:"encryption-enabled,omitempty" xml:"encryption-enabled"`
	EncryptionEnabledV2 bool `json:"encryption-enabled-v2,omitempty" xml:"encryption-enabled-v2"`
	EncryptionActive    int  `json:"encryption-active,omitempty" xml:"encryption-active"`
	EncryptionActiveV2  bool `json:"encryption-active-v2,omitempty" xml:"encryption-active-v2"`

	// Misc port features
	TargetDrivenZoningEnable   int  `json:"target-driven-zoning-enable,omitempty" xml:"target-driven-zoning-enable"`
	TargetDrivenZoningEnableV2 bool `json:"target-driven-zoning-enable-v2,omitempty" xml:"target-driven-zoning-enable-v2"`

	SIMPortEnabled   int  `json:"sim-port-enabled,omitempty" xml:"sim-port-enabled"`
	SIMPortEnabledV2 bool `json:"sim-port-enabled-v2,omitempty" xml:"sim-port-enabled-v2"`

	MirrorPortEnabled   int  `json:"mirror-port-enabled,omitempty" xml:"mirror-port-enabled"`
	MirrorPortEnabledV2 bool `json:"mirror-port-enabled-v2,omitempty" xml:"mirror-port-enabled-v2"`

	NONDFEEnabled         int    `json:"non-dfe-enabled,omitempty" xml:"non-dfe-enabled"`
	OctetSpeedCombo       int    `json:"octet-speed-combo,omitempty" xml:"octet-speed-combo"`
	OctetSpeedComboString string `json:"octet-speed-combo-string,omitempty" xml:"octet-speed-combo-string"`

	FaultyDelayEnabled   int  `json:"fault-delay-enabled,omitempty" xml:"fault-delay-enabled"`
	FaultyDelayEnabledV2 bool `json:"fault-delay-enabled-v2,omitempty" xml:"fault-delay-enabled-v2"`

	RSCNSuppressionEnabled   int  `json:"rscn-suppression-enabled,omitempty" xml:"rscn-suppression-enabled"`
	RSCNSuppressionEnabledV2 bool `json:"rscn-suppression-enabled-v2,omitempty" xml:"rscn-suppression-enabled-v2"`

	LOSTOVModeEnabled       int    `json:"los-tov-mode-enabled,omitempty" xml:"los-tov-mode-enabled"`
	LOSTOVModeEnabledString string `json:"los-tov-mode-enabled-string,omitempty" xml:"los-tov-mode-enabled-string"`

	// ISL properties
	ISLReadyModeEnabled   int  `json:"isl-ready-mode-enabled,omitempty" xml:"isl-ready-mode-enabled"`
	ISLReadyModeEnabledV2 bool `json:"isl-ready-mode-enabled-v2,omitempty" xml:"isl-ready-mode-enabled-v2"`

	VCLinkInit   int  `json:"vc-link-init,omitempty" xml:"vc-link-init"`
	VCLinkInitV2 bool `json:"vc-link-init-v2,omitempty" xml:"vc-link-init-v2"`

	TrunkPortEnabled   int  `json:"trunk-port-enabled,omitempty" xml:"trunk-port-enabled"`
	TrunkPortEnabledV2 bool `json:"trunk-port-enabled-v2,omitempty" xml:"trunk-port-enabled-v2"`

	// Long distance
	LongDistance       int    `json:"long-distance,omitempty" xml:"long-distance"`
	LongDistanceString string `json:"long-distance-string,omitempty" xml:"long-distance-string"`

	// Licensing
	PodLicenseStatus bool   `json:"pod-license-status,omitempty" xml:"pod-license-status"`
	PodLicenseState  string `json:"pod-license-state,omitempty" xml:"pod-license-state"`

	// Neighbor
	Neighbor             Neighbor `json:"neighbor,omitempty" xml:"neighbor"`
	NeighborSlotPort     string   `json:"neighbor-slot-port,omitempty" xml:"neighbor-slot-port"`
	NeighborNodeWWN      string   `json:"neighbor-node-wwn,omitempty" xml:"neighbor-node-wwn"`
	NeighborPortIndex    int      `json:"neighbor-port-index,omitempty" xml:"neighbor-port-index"`
	NeighborSwitchUFName string   `json:"neighbor-switch-user-friendly-name,omitempty" xml:"neighbor-switch-user-friendly-name"`
	NeighborSwitchIPv4   string   `json:"neighbor-switch-ipv4-address,omitempty" xml:"neighbor-switch-ipv4-address"`
	NeighborSwitchIPv6   string   `json:"neighbor-switch-ipv6-address,omitempty" xml:"neighbor-switch-ipv6-address"`
	NeighborFabricName   string   `json:"neighbor-fabric-name,omitempty" xml:"neighbor-fabric-name"`

	// FC addressing
	FCID         int    `json:"fcid,omitempty" xml:"fcid"`
	FCIDHex      string `json:"fcid-hex,omitempty" xml:"fcid-hex"`
	DefaultIndex int    `json:"default-index,omitempty" xml:"default-index"`
	Index        int    `json:"index,omitempty" xml:"index"`

	// Diagnostics / health
	DiagnosticsStatus    int    `json:"diagnostics-status,omitempty" xml:"diagnostics-status"`
	BladePortNumber      int    `json:"blade-port-number,omitempty" xml:"blade-port-number"`
	PhysicalState        string `json:"physical-state,omitempty" xml:"physical-state"`
	SegmentationReason   string `json:"segmentation-reason,omitempty" xml:"segmentation-reason"`
	PortHealth           string `json:"port-health,omitempty" xml:"port-health"`
	MeasuredLinkDistance string `json:"measured-link-distance,omitempty" xml:"measured-link-distance"`
	ChipInstance         int    `json:"chip-instance,omitempty" xml:"chip-instance"`
	ChipBuffersAvailable int    `json:"chip-buffers-available,omitempty" xml:"chip-buffers-available"`

	// Address binding
	UserBoundEnabled   bool   `json:"user-bound-enabled,omitempty" xml:"user-bound-enabled"`
	AddressBindingMode string `json:"address-binding-mode,omitempty" xml:"address-binding-mode"`

	// ACL settings
	MSACLApplicationServerAccess          string `json:"ms-acl-application-server-access,omitempty" xml:"ms-acl-application-server-access"`
	MSACLEnhancedFabricConfigServerAccess string `json:"ms-acl-enhanced-fabric-configuration-server-access,omitempty" xml:"ms-acl-enhanced-fabric-configuration-server-access"`
	MSACLFabricConfigServerAccess         string `json:"ms-acl-fabric-configuration-server-access,omitempty" xml:"ms-acl-fabric-configuration-server-access"`
	MSACLFabricDeviceMgmtAccess           string `json:"ms-acl-fabric-device-management-interface-access,omitempty" xml:"ms-acl-fabric-device-management-interface-access"`
	MSACLFabricZoneServerAccess           string `json:"ms-acl-fabric-zone-server-access,omitempty" xml:"ms-acl-fabric-zone-server-access"`
	MSACLUnzonedNameServerAccess          string `json:"ms-acl-unzoned-name-server-access,omitempty" xml:"ms-acl-unzoned-name-server-access"`

	// Misc
	FlexportProtocol         string `json:"flexport-protocol,omitempty" xml:"flexport-protocol"`
	EthernetLAGName          string `json:"ethernet-lag-name,omitempty" xml:"ethernet-lag-name"`
	EthernetLAGMemberTimeout int    `json:"ethernet-lag-member-timeout,omitempty" xml:"ethernet-lag-member-timeout"`
	EthernetLLDPEnabled      bool   `json:"ethernet-lldp-enabled,omitempty" xml:"ethernet-lldp-enabled"`
	EthernetLLDPProfileName  string `json:"ethernet-lldp-profile-name,omitempty" xml:"ethernet-lldp-profile-name"`

	LAGName             string `json:"lag-name,omitempty" xml:"lag-name"`
	LAGMemberLinkStatus string `json:"lag-member-link-status,omitempty" xml:"lag-member-link-status"`

	// Buffers and metrics
	ReservedBuffers        int `json:"reserved-buffers,omitempty" xml:"reserved-buffers"`
	AvgTransmitBufferUsage int `json:"average-transmit-buffer-usage,omitempty" xml:"average-transmit-buffer-usage"`
	AvgTransmitFrameSize   int `json:"average-transmit-frame-size,omitempty" xml:"average-transmit-frame-size"`
	AvgReceiveBufferUsage  int `json:"average-receive-buffer-usage,omitempty" xml:"average-receive-buffer-usage"`
	AvgReceiveFrameSize    int `json:"average-receive-frame-size,omitempty" xml:"average-receive-frame-size"`
	CurrentBufferUsage     int `json:"current-buffer-usage,omitempty" xml:"current-buffer-usage"`
	RecommendedBuffers     int `json:"recommended-buffers,omitempty" xml:"recommended-buffers"`

	// Misc
	PortGenerationNumber   int    `json:"port-generation-number,omitempty" xml:"port-generation-number"`
	PortSCN                string `json:"port-scn,omitempty" xml:"port-scn"`
	CableDistance          int    `json:"cable-distance,omitempty" xml:"cable-distance"`
	DisableReason          string `json:"disable-reason,omitempty" xml:"disable-reason"`
	AuthenticationProtocol string `json:"authentication-protocol,omitempty" xml:"authentication-protocol"`

	// Bool toggles
	PortPeerBeaconEnabled   bool `json:"port-peer-beacon-enabled,omitempty" xml:"port-peer-beacon-enabled"`
	CleanAddressEnabled     bool `json:"clean-address-enabled,omitempty" xml:"clean-address-enabled"`
	CongestionSignalEnabled bool `json:"congestion-signal-enabled,omitempty" xml:"congestion-signal-enabled"`

	// Fabric routing
	FCRouterPortCost       int    `json:"fc-router-port-cost,omitempty" xml:"fc-router-port-cost"`
	EdgeFabricID           int    `json:"edge-fabric-id,omitempty" xml:"edge-fabric-id"`
	DestinationFabricPSWWN string `json:"destination-fabric-principal-switch-wwn,omitempty" xml:"destination-fabric-principal-switch-wwn"`
	PreferredFrontDomainID int    `json:"preferred-front-domain-id,omitempty" xml:"preferred-front-domain-id"`

	LinkLatencyDeterminationEnabled bool `json:"link-latency-determination-enabled,omitempty" xml:"link-latency-determination-enabled"`
}

// FibrechannelPort represents fibrechannel interface port data.
/*
Description on page 180
*/

// Neighbor holds a list of neighbo(u)r WWNs.
type Neighbor struct {
	WWN []string `json:"wwn,omitempty" xml:"wwn"`
}
