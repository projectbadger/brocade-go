
# fdmi

## Index

- [type BrocadeFDMI](#type-brocadefdmi)
- [type HBA](#type-hba)
- [type Port](#type-port)
- [type RESTFDMI](#type-restfdmi)
  - [NewRESTFDMI() RESTFDMI](#func-newrestfdmi-restfdmi)


## type [BrocadeFDMI](<brocadeFDMI.go#L9>)

Fabric Device Management Interface (FDMI) information
for the specified switch.

Documentation source: https://docs.broadcom.com/doc/FOS-82X-REST-API-RM - page 114 (20220701)
```go
type BrocadeFDMI struct {
	XMLName	xml.Name	`json:"-" xml:"brocade-fdmi"`
	// List of HBA attributes registered with FDMI.
	HBA	[]HBA	`json:"hba" xml:"hba"`
	Port	[]Port	`json:"port" xml:"port"`
}
```

## type [HBA](<brocadeFDMI.go#L17>)

Fibre Channel host bus adapter
```go
type HBA struct {
	XMLName	xml.Name	`json:"-" xml:"hba"`
	// A 64-bit Name_Identifier that is uniquely associated
	// with the HBA among all HBAs in the same Fibre
	// Channel interaction space (see FC-FS). The HBA
	// identifier for an HBA may be the same as the
	// Name_Identifier of an Nx_Port on the HBA if the
	// required persistence is satisfied. Once an HBA has
	// registered a Name_Identifier as its HBA identifier,
	// that Name_Identifier persists (for example, across
	// power cycles) as the HBA identifier for the HBA.
	HBAId	string	`json:"hba-id,omitempty" xml:"hba-id"`
	// The domain directly attached to the HBA.
	DomainId	string	`json:"domain-id,omitempty" xml:"domain-id"`
	// Manufacturer of the host adapter.
	Manufacturer	string	`json:"manufacturer,omitempty" xml:"manufacturer"`
	// Serial number of the host adapter.
	SerialNumber	string	`json:"serial-number,omitempty" xml:"serial-number"`
	// Model of the host adapter.
	Model	string	`json:"model,omitempty" xml:"model"`
	// Description of the model of the host adapter.
	ModelDescription	string	`json:"model-description,omitempty" xml:"model-description"`
	// An 8-byte value that identifies the node that
	// contains the Nx_Ports on the host adapter. If all
	// Nx_Ports on the host bus adapter have the same node
	// name, the node-name resource returned for the host
	// bus adapter matches the node-name resource for its
	// Nx_Ports. If not all Nx_Ports on the host bus
	// adapter have the same node name, the node-name
	// resource is not returned for the host bus adapter.
	NodeName	string	`json:"node-name,omitempty" xml:"node-name"`
	// The name registered with the Name Server as a Name
	// Server Symbolic Node Name object, which is subject
	// to all the description and constraints of that
	// object.
	NodeSymbolicName	string	`json:"node-symbolic-name,omitempty" xml:"node-symbolic-name"`
	// The hardware version level of the host adapter.
	HardwareVersion	string	`json:"hardware-version,omitempty" xml:"hardware-version"`
	// The ersion level of the driver software that
	// controls a host adapter.
	DriverVersion	string	`json:"driver-version,omitempty" xml:"driver-version"`
	//  Option ROM or BIOS version of a host adapter.
	OptionRomVersion	string	`json:"option-rom-version,omitempty" xml:"option-rom-version"`
	// The version of firmware executed by a host adapter.
	FirmwareVersion	string	`json:"firmware-version,omitempty" xml:"firmware-version"`
	// The type and version of the operating system that
	// controls the host bus adapter.
	OsNameAndVersion	string	`json:"os-name-and-version,omitempty" xml:"os-name-and-version"`
	// The maximum size CT payload in 32-bit words,
	// including all CT headers but no FC frame header(s)
	// that may be sent or received by application software
	// resident in the host in which the host bus adapter
	// is installed. If the host bus adapter does not
	// support generic CT capability for application
	// software on the host in which it is installed, this
	// attribute will not be returned.
	MaxCtPayload	uint32	`json:"max-ct-payload,omitempty" xml:"max-ct-payload"`
	// The T10 vendor ID of the manufacturer of the HBA, or
	// an OEM of the HBA.
	VendorId	string	`json:"vendor-id,omitempty" xml:"vendor-id"`
	// A value with vendor-specific use.
	VendorSpecificInfo	string	`json:"vendor-specific-info,omitempty" xml:"vendor-specific-info"`
	// The number of Nx_Ports on the HBA.
	NumberOfPorts	uint32	`json:"number-of-ports,omitempty" xml:"number-of-ports"`
	// An 8-byte binary value equal to the Fabric_Name of
	// the fabric associated with the HBA. If the HBA is
	// associated with more than one fabric, fabric-name
	// will not be provided.
	FabricName	string	`json:"fabric-name,omitempty" xml:"fabric-name"`
	// The identification and version of a boot BIOS
	// provided by the HBA.
	BootBiosVersion	string	`json:"boot-bios-version,omitempty" xml:"boot-bios-version"`
	//  Whether a boot BIOS provided by the HBA is enabled or disabled.
	BootBiosEnabled	bool	`json:"boot-bios-enabled,omitempty" xml:"boot-bios-enabled"`
	// A list of port WWNs associated with the HBA. This
	// list corresponds with the port-name resource under
	// the port list.
	HbaPortList	[]string	`json:"hba-port-list,omitempty" xml:"hba-port-list"`
}
```

## type [Port](<brocadeFDMI.go#L97>)
```go
type Port struct {
	XMLName	xml.Name	`json:"-" xml:"port"`
	// A WWN that identifies the port associated with the
	// host bus adapter specified by the hba-id resource.
	PortName	string	`json:"port-name,omitempty" xml:"port-name"`
	//  64-bit Name_Identifier that is uniquely associated
	// with the HBA among all HBAs in the same Fibre
	// Channel interaction space.
	HbaId	string	`json:"hba-id,omitempty" xml:"hba-id"`
	// The domain directly attached to the HBA port.
	DomainId	string	`json:"domain-id,omitempty" xml:"domain-id"`
	// The Port Symbolic Name object registered with the
	// Name Server, which is subject to all the description
	// and constraints of that object defined in FC-GS.
	PortSymbolicName	string	`json:"port-symbolic-name,omitempty" xml:"port-symbolic-name"`
	// The Port Identifier object registered with the Name
	// Server, which is subject to all the description and
	// constraints of that object defined in FC-GS.
	PortId	string	`json:"port-id,omitempty" xml:"port-id"`
	// The Name Server port type object registered with the
	// Name Server, which is subject to all the description
	// and constraints of that object defined in FC-GS.
	PortType	string	`json:"port-type,omitempty" xml:"port-type"`
	// The Class of Service object registered with the Name
	// Server, which is subject to all the description and
	// constraints of that object defined in FC-GS.
	SupportedClassOfService	string	`json:"supported-class-of-service,omitempty" xml:"supported-class-of-service"`
	// The FC-4 types attribute registered with the HBA
	// Management Server as a port attribute. An Nx_Port
	// registers a supported FC-4 types value that
	// indicates "support" for any FC-4 type that it is
	// able to be configured to support.
	SupportedFc4Type	string	`json:"supported-fc4-type,omitempty" xml:"supported-fc4-type"`
	// The Port active FC-4 types attribute registered with
	// the Name Server, which is subject to all the
	// description and constraints of that object defined
	// in FC-GS. An Nx_Port registers a supported FC-4
	// types value that indicates "support" for any FC-4
	// type that it is able to be configured to support.
	ActiveFc4Types	string	`json:"active-fc4-types,omitempty" xml:"active-fc4-types"`
	// The supported transmission speeds of the Nx_Port.
	SupportedSpeed	string	`json:"supported-speed,omitempty" xml:"supported-speed"`
	// The Current Port Speed attribute returned by the HBA
	// Management Server as a port attribute if either it
	// has been registered or it has been determined by the
	// HBA Management Server.
	CurrentPortSpeed	string	`json:"current-port-speed,omitempty" xml:"current-port-speed"`
	// The Maximum Frame Size attribute registered with the
	// HBA Management Server as a Port attribute.
	MaximumFrameSize	uint32	`json:"maximum-frame-size,omitempty" xml:"maximum-frame-size"`
	// The OS Device Name attribute registered with the HBA
	// Management Server as a port attribute.
	OsDeviceName	string	`json:"os-device-name,omitempty" xml:"os-device-name"`
	// The Host Name attribute registered with the HBA
	// Management Server as a Port attribute.
	HostName	string	`json:"host-name,omitempty" xml:"host-name"`
	// The Node Name attribute registered with the Name
	// Server for the Nx_Port.
	NodeName	string	`json:"node-name,omitempty" xml:"node-name"`
	// The Port Fabric Name attribute registered with the
	// HBA Management Server as an Nx_Port attribute. The
	// Port Fabric Name attribute contains an 8-byte binary
	// value equal to the Fabric_Name of the fabric
	// associated with the Nx_Port
	FabricName	string	`json:"fabric-name,omitempty" xml:"fabric-name"`
	// The Port State (see SM-HBA) for the specified
	// Nx_Port. The Port State is an integer where the
	// value indicates the current state of the Nx_Port.
	PortState	string	`json:"port-state,omitempty" xml:"port-state"`
	// The number of FC_Ports that are visible to the
	// Nx_Port identified in the request.
	NumberOfDiscoveredPorts	uint32	`json:"number-of-discovered-ports,omitempty" xml:"number-of-discovered-ports"`
	// The vendor-specific service category resource.
	VsaServiceCategory	string	`json:"vsa-service-category,omitempty" xml:"vsa-service-category"`
	// The vendor-specific guid resource.
	VsaGuid	string	`json:"vsa-guid,omitempty" xml:"vsa-guid"`
	// The vendor-specific version.
	VsaVersion	string	`json:"vsa-version,omitempty" xml:"vsa-version"`
	// The vendor-specific product name resource.
	VsaProductName	string	`json:"vsa-product-name,omitempty" xml:"vsa-product-name"`
	// The vendor-specific port information resource.
	VsaPortInfo	string	`json:"vsa-port-info,omitempty" xml:"vsa-port-info"`
	// The vendor-specific QOS supported resource.
	VsaQusSupported	string	`json:"vsa-qus-supported,omitempty" xml:"vsa-qus-supported"`
	// The vendor-specific security resource.
	VsaSecurity	string	`json:"vsa-security,omitempty" xml:"vsa-security"`
	// The vendor-specific storage array family.
	VsaStorageArrayFamily	string	`json:"vsa-storage-array-family,omitempty" xml:"vsa-storage-array-family"`
	// The vendor-specific storage array name.
	VsaStorageArrayName	string	`json:"vsa-storage-array-name,omitempty" xml:"vsa-storage-array-name"`
	// The vendor-specific storage system model.
	VsaStorageArraySystemModel	string	`json:"vsa-storage-array-system-model,omitempty" xml:"vsa-storage-array-system-model"`
	//  The vendor-specific storage array operating system.
	VsaStorageArrayOs	string	`json:"vsa-storage-array-os,omitempty" xml:"vsa-storage-array-os"`
	// The vendor-specific storage array node count.
	VsaStorageArrayNodeCount	uint32	`json:"vsa-storage-array-node-count,omitempty" xml:"vsa-storage-array-node-count"`
	// A list of vendor-specific storage array nodes.
	VsaStorageArrayNodes	[]string	`json:"vsa-storage-array-nodes,omitempty" xml:"vsa-storage-array-nodes>nodes"`
	// A list of the vendor-specific connected ports.
	VsaConnectedPorts	[]string	`json:"vsa-connected-ports,omitempty" xml:"vsa-connected-ports>wwns"`
}
```

## type [RESTFDMI](<methods.go#L9>)
```go
type RESTFDMI interface {
	Name() string
	URIPath() string
	GetHBA() ([]HBA, error)
	GetHBAResponse() (*http.Response, error)
	GetPort() ([]Port, error)
	GetPortResponse() (*http.Response, error)
}
```

## func [NewRESTFDMI() RESTFDMI](<methods.go#L35>)

```go
func NewRESTFDMI(config *api_interface.RESTConfig) RESTFDMI
```

