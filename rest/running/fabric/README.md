
# fabric

## Index

- [type BrocadeFabric](#type-brocadefabric)
- [type FabricSwitch](#type-fabricswitch)
- [type RESTFabric](#type-restfabric)
  - [NewRESTFabric() RESTFabric](#func-newrestfabric-restfabric)


## type [BrocadeFabric](<brocadeFabric.go#L7>)

Module to retrieve information on the switches in a
fabric.
```go
type BrocadeFabric struct {
	XMLName	xml.Name	`json:"-" xml:"brocade-fabric"`
	// The list of configured switches in the fabric.
	FabricSwitch	[]FabricSwitch	`json:"fabric-switch" xml:"fabric-switch"`
}
```

## type [FabricSwitch](<brocadeFabric.go#L13>)
```go
type FabricSwitch struct {
	XMLName	xml.Name	`json:"-" xml:"fabric-switch"`
	// he Fibre Channel WWN of the switch.
	Name	string	`json:"name" xml:"name"`
	// The ASCII name assigned to the switch by the
	// administrator.
	SwitchUserFriendlyName	string	`json:"switch-user-friendly-name,omitempty" xml:"switch-user-friendly-name"`
	// The chassis WWN.
	ChassisWWN	string	`json:"chassis-wwn,omitempty" xml:"chassis-wwn"`
	// The ASCII name assigned to the switch chassis by the
	// administrator.
	ChassisUserFriendlyName	string	`json:"chassis-user-friendly-name,omitempty" xml:"chassis-user-friendly-name"`
	// The highest level in the three-level addressing
	// hierarchy used in the Fibre Channel address
	// identifier.
	// A domain typically is associated with a single Fibre
	// Channel switch.
	DomainID	string	`json:"domain-id,omitempty" xml:"domain-id"`
	// Indicates if this switch is the fabric's principal switch.
	Principal	uint8	`json:"principal,omitempty" xml:"principal"`
	// This parameter is deprecated.
	// Use the fcid-hex parameter. The destination Fibre
	// Channel ID (D_ID) of the switch (decimal format).
	FCID	string	`json:"fcid,omitempty" xml:"fcid"`
	// The destination Fibre Channel ID (D_ID) of the
	// switch (hexadecimal string).
	FCIDHex	string	`json:"fcid-hex,omitempty" xml:"fcid-hex"`
	// The IPv4 address for the switch.
	IPAddress	string	`json:"ip-address,omitempty" xml:"ip-address"`
	// The IPv4 address the switch is using for
	// Fibre Channel over IP
	FCIPAddress	string	`json:"fcip-address,omitempty" xml:"fcip-address"`
	// The IPv6 address for the switch.
	IPv6Address	string	`json:"ipv6-address,omitempty" xml:"ipv6-address"`
	// A human-readable string identifying the firmware
	// version running on the switch.
	FirmwareVersion	string	`json:"firmware-version,omitempty" xml:"firmware-version"`
	// The number of paths available to each remote domain.
	PathCount	uint32	`json:"path-count,omitempty" xml:"path-count"`
}
```

## type [RESTFabric](<methods.go#L13>)

RESTFabric describes an interface for interacting with the
*fabric* module.
Fetch a new instance using the NewRESTFabric function.
```go
type RESTFabric interface {
	Name() string
	URIPath() string
	GetFabricSwitch() (*FabricSwitch, errors.BrocadeErr)
	GetFabricSwitchResponse() (*http.Response, error)
}
```

## func [NewRESTFabric() RESTFabric](<methods.go#L35>)

Returns a new RESTFabric implementation for interacting
with the *fabric* module.


```go
func NewRESTFabric(config *api_interface.RESTConfig) RESTFabric
```

