
# brocade_interface

175

## Index

- [type BrocadeInterface](#type-brocadeinterface)
- [type Fibrechannel](#type-fibrechannel)
- [type Neighbor](#type-neighbor)
- [type RESTInterface](#type-restinterface)
  - [NewRESTInterface() RESTInterface](#func-newrestinterface-restinterface)


## type [BrocadeInterface](<interface.go#L4>)
```go
type BrocadeInterface struct {
	Fibrechannel Fibrechannel `json:"fibrechannel" xml:"fibrechannel"`
}
```

## type [Fibrechannel](<interfaceFibrechannel.go#L7>)
```go
type Fibrechannel struct {
	XMLName				xml.Name	`json:"-,omitempty" xml:"fibrechannel"`
	Name				string		`json:"name,omitempty" xml:"name"`
	WWN				string		`json:"wwn,omitempty" xml:"wwn"`
	PortType			int		`json:"port-type,omitempty" xml:"port-type"`
	Speed				int		`json:"speed,omitempty" xml:"speed"`
	MaxSpeed			int		`json:"max-speed,omitempty" xml:"max-speed"`
	UserFriendlyName		string		`json:"user-friendly-name,omitempty" xml:"user-friendly-name"`
	OperationalStatus		int		`json:"operational-status,omitempty" xml:"operational-status"`
	EnabledState			int		`json:"enabled-state,omitempty" xml:"enabled-state"`
	IsEnabledState			bool		`json:"is-enabled-state,omitempty" xml:"is-enabled-state"`
	AutoNegotiate			int		`json:"auto-negotiate,omitempty" xml:"auto-negotiate"`
	ISLReadyModeEnabled		int		`json:"isl-ready-mode-enabled,omitempty" xml:"isl-ready-mode-enabled"`
	LongDistance			int		`json:"long-distance,omitempty" xml:"long-distance"`
	TrunkPortEnabled		int		`json:"trunk-port-enabled,omitempty" xml:"trunk-port-enabled"`
	VcLinkInit			int		`json:"vc-link-init,omitempty" xml:"vc-link-init"`
	PodLicenseStatus		bool		`json:"pod-license-status,omitempty" xml:"pod-license-status"`
	DefaultIndex			int		`json:"default-index,omitempty" xml:"default-index"`
	FCID				int		`json:"fcid,omitempty" xml:"fcid"`
	FCIDHex				string		`json:"fcid-hex,omitempty" xml:"fcid-hex"`
	PhysicalState			string		`json:"physical-state,omitempty" xml:"physical-state"`
	PersistentDisable		int		`json:"persistent-disable,omitempty" xml:"persistent-disable"`
	GPortLocked			int		`json:"g-port-locked,omitempty" xml:"g-port-locked"`
	EPortDisable			int		`json:"e-port-disable,omitempty" xml:"e-port-disable"`
	ExPortEnabled			int		`json:"ex-port-enabled,omitempty" xml:"ex-port-enabled"`
	NPIVEnabled			int		`json:"npiv-enabled,omitempty" xml:"npiv-enabled"`
	NPIVPPLimit			int		`json:"npiv-pp-limit,omitempty" xml:"npiv-pp-limit"`
	RateLimitEnabled		int		`json:"rate-limit-enabled,omitempty" xml:"rate-limit-enabled"`
	QOSEnabled			int		`json:"qos-enabled,omitempty" xml:"qos-enabled"`
	CSCTLModeEnabled		int		`json:"csctl-mode-enabled,omitempty" xml:"csctl-mode-enabled"`
	PortAutoDisableEnabled		int		`json:"port-autodisable-enabled,omitempty" xml:"port-autodisable-enabled"`
	EPortCredit			int		`json:"e-port-credit,omitempty" xml:"e-port-credit"`
	DPortEnable			int		`json:"d-port-enable,omitempty" xml:"d-port-enable"`
	OctetSpeedCombo			int		`json:"octet-speed-combo,omitempty" xml:"octet-speed-combo"`
	CompressionConfigured		int		`json:"compression-configured,omitempty" xml:"compression-configured"`
	CompressionActive		int		`json:"compression-active,omitempty" xml:"compression-active"`
	EncryptionActive		int		`json:"encryption-active,omitempty" xml:"encryption-active"`
	MirrorPortEnabled		int		`json:"mirror-port-enabled,omitempty" xml:"mirror-port-enabled"`
	NONDFEEnabled			int		`json:"non-dfe-enabled,omitempty" xml:"non-dfe-enabled"`
	FECEnabled			int		`json:"fec-enabled,omitempty" xml:"fec-enabled"`
	FECActive			int		`json:"fec-active,omitempty" xml:"fec-active"`
	ViaTTSFECEnabled		int		`json:"via-tts-fec-enabled,omitempty" xml:"via-tts-fec-enabled"`
	Neighbour			Neighbor	`json:"neighbor,omitempty" xml:"neighbor"`
	TargetDrivenZoningEnable	int		`json:"target-driven-zoning-enable,omitempty" xml:"target-driven-zoning-enable"`
	NPIVFLOGILogoutEnabled		int		`json:"npiv-flogi-logout-enabled,omitempty" xml:"npiv-flogi-logout-enabled"`
	SIMPortEnabled			int		`json:"sim-port-enabled,omitempty" xml:"sim-port-enabled"`
	FPortBuffers			int		`json:"f-port-buffers,omitempty" xml:"f-port-buffers"`
	FaultyDelayEnabled		int		`json:"fault-delay-enabled,omitempty" xml:"fault-delay-enabled"`
	CreditRecoveryEnabled		int		`json:"credit-recovery-enabled,omitempty" xml:"credit-recovery-enabled"`
	CreditRecoveryActive		int		`json:"credit-recovery-active,omitempty" xml:"credit-recovery-active"`
	RSCNSuppressionEnabled		int		`json:"rscn-suppression-enabled,omitempty" xml:"rscn-suppression-enabled"`
	LOSTOVModeEnabled		int		`json:"los-tov-mode-enabled,omitempty" xml:"los-tov-mode-enabled"`
}
```

## type [Neighbor](<interfaceFibrechannel.go#L67>)

Neighbor holds a list of neighbo(u)r WWNs.
```go
type Neighbor struct {
	WWN []string `json:"wwn,omitempty" xml:"wwn"`
}
```

## type [RESTInterface](<methods.go#L10>)
```go
type RESTInterface interface {
	Name() string
	// GetFibrechannel() ([]Port, error)
	GetFibrechannelResponse() (*http.Response, errors.BrocadeErr)
	GetFibrechannel() ([]Fibrechannel, errors.BrocadeErr)
}
```

## func [NewRESTInterface() RESTInterface](<methods.go#L34>)

```go
func NewRESTInterface(config *api_interface.RESTConfig) RESTInterface
```

