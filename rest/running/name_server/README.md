
# name_server

## Index

- [type BrocadeNameServer](#type-brocadenameserver)
- [type FibrechannelNameServer](#type-fibrechannelnameserver)
- [type RESTNameServer](#type-restnameserver)
  - [NewRESTNameServer() RESTNameServer](#func-newrestnameserver-restnameserver)


## type [BrocadeNameServer](<brocadeNameServer.go#L5>)
```go
type BrocadeNameServer struct {
	XMLName			xml.Name			`json:"-" xml:"brocade-name-server"`
	FibrechannelNameServer	[]FibrechannelNameServer	`json:"fibrechannel-name-server" xml:"fibrechannel-name-server>fibrechannel-name-server"`
}
```

## type [FibrechannelNameServer](<brocadeNameServer.go#L10>)
```go
type FibrechannelNameServer struct {
	XMLName				xml.Name	`json:"-" xml:"fibrechannel-name-server"`
	PortId				string		`json:"port-id" xml:"port-id"`
	PortName			string		`json:"port-name,omitempty" xml:"port-name"`
	PortSymbolicName		string		`json:"port-symbolic-name,omitempty" xml:"port-symbolic-name"`
	FabricPortName			string		`json:"fabric-port-name,omitempty" xml:"fabric-port-name"`
	PermanentPortName		string		`json:"permanent-port-name,omitempty" xml:"permanent-port-name"`
	NodeName			string		`json:"node-name,omitempty" xml:"node-name"`
	NodeSymbolicName		string		`json:"node-symbolic-name,omitempty" xml:"node-symbolic-name"`
	ClassOfService			string		`json:"class-of-service,omitempty" xml:"class-of-service"`
	Fc4Type				string		`json:"fc4-type,omitempty" xml:"fc4-type"`
	Fc4Features			string		`json:"fc4-features,omitempty" xml:"fc4-features"`
	PortType			string		`json:"port-type,omitempty" xml:"port-type"`
	StateChangeRegistration		string		`json:"state-change-registration,omitempty" xml:"state-change-registration"`
	NameServerDeviceType		string		`json:"name-server-device-type,omitempty" xml:"name-server-device-type"`
	PortIndex			int		`json:"port-index,omitempty" xml:"port-index"`
	ShareArea			string		`json:"share-area,omitempty" xml:"share-area"`
	FrameRedirection		string		`json:"frame-redirection,omitempty" xml:"frame-redirection"`
	Partial				string		`json:"partial,omitempty" xml:"partial"`
	LSAN				string		`json:"lsan,omitempty" xml:"lsan"`
	LinkSpeed			string		`json:"link-speed,omitempty" xml:"link-speed"`
	ProtocolSpeed			string		`json:"protocol-speed,omitempty" xml:"protocol-speed"`
	PortProperties			string		`json:"port-properties,omitempty" xml:"port-properties"`
	CascadedAg			string		`json:"cascaded-ag,omitempty" xml:"cascaded-ag"`
	ConnectedThroughAg		string		`json:"connected-through-ag,omitempty" xml:"connected-through-ag"`
	RealDeviceBehindAg		string		`json:"real-device-behind-ag,omitempty" xml:"real-device-behind-ag"`
	FCOEDevice			string		`json:"fcoe-device,omitempty" xml:"fcoe-device"`
	SlowDrainDeviceQuarantine	string		`json:"slow-drain-device-quarantine,omitempty" xml:"slow-drain-device-quarantine"`
	ConnectedThroughFcLag		bool		`json:"connected-through-fc-lag,omitempty" xml:"connected-through-fc-lag"`
}
```

## type [RESTNameServer](<methods.go#L9>)
```go
type RESTNameServer interface {
	Name() string
	URIPath() string
	GetNameServer() ([]FibrechannelNameServer, error)
	GetNameServerResponse() (*http.Response, error)
	GetNameServerForPort(portId string) (*FibrechannelNameServer, error)
	GetNameServerForPortResponse(portId string) (*http.Response, error)
}
```

## func [NewRESTNameServer() RESTNameServer](<methods.go#L22>)

```go
func NewRESTNameServer(cfg *api_interface.RESTConfig) RESTNameServer
```

